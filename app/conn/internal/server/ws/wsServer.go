package ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/cherish-chat/xxim-server/app/conn/internal/svc"
	"github.com/cherish-chat/xxim-server/app/conn/internal/types"
	"github.com/cherish-chat/xxim-server/common/xhttp"
	"github.com/cherish-chat/xxim-server/common/xtrace"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"nhooyr.io/websocket"
)

// Server enables broadcasting to a set of subscribers.
type Server struct {
	serveMux         http.ServeMux
	svcCtx           *svc.ServiceContext
	addSubscriber    func(c *types.UserConn)
	deleteSubscriber func(c *types.UserConn)
	beforeConnect    func(ctx context.Context, param types.ConnParam) (int, error)
	onReceive        func(ctx context.Context, c *types.UserConn, typ int, msg []byte)
}

func (s *Server) SetOnReceive(f func(ctx context.Context, c *types.UserConn, typ int, msg []byte)) {
	s.onReceive = f
}

func (s *Server) SetBeforeConnect(f func(ctx context.Context, param types.ConnParam) (int, error)) {
	s.beforeConnect = f
}

func (s *Server) SetAddSubscriber(f func(c *types.UserConn)) {
	s.addSubscriber = f
}

func (s *Server) SetDeleteSubscriber(f func(c *types.UserConn)) {
	s.deleteSubscriber = f
}

// NewServer constructs a Server with the defaults.
func NewServer(
	svcCtx *svc.ServiceContext,
) types.IServer {
	s := &Server{
		svcCtx:           svcCtx,
		addSubscriber:    func(c *types.UserConn) {},
		deleteSubscriber: func(c *types.UserConn) {},
		beforeConnect:    func(ctx context.Context, param types.ConnParam) (int, error) { return 0, nil },
	}
	s.serveMux.HandleFunc("/", s.subscribeHandler)
	s.serveMux.HandleFunc("/ws", s.subscribeHandler)
	return s
}

func (s *Server) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.svcCtx.Config.Websocket.Host, s.svcCtx.Config.Websocket.Port))
	if err != nil {
		return err
	}
	logx.Infof("listening on http://%v", l.Addr())

	hs := &http.Server{
		Handler:      s,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	errc := make(chan error, 1)
	go func() {
		errc <- hs.Serve(l)
	}()

	select {
	case err := <-errc:
		logx.Errorf("failed to serve: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return hs.Shutdown(ctx)
}

type userConn struct {
	ws *websocket.Conn
}

func (c *userConn) Close(code int, desc string) error {
	return c.ws.Close(websocket.StatusCode(code), desc)
}

func (c *userConn) Write(ctx context.Context, typ int, msg []byte) error {
	return c.ws.Write(ctx, websocket.MessageType(typ), msg)
}

func (c *userConn) Read(ctx context.Context) (typ int, msg []byte, err error) {
	messageType, data, err := c.ws.Read(ctx)
	return int(messageType), data, err
}

func (s *Server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	logger := logx.WithContext(r.Context())
	headers := make(map[string]string)
	for k, v := range r.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	param := types.ConnParam{
		UserId:      r.URL.Query().Get("userId"),
		Token:       r.URL.Query().Get("token"),
		DeviceId:    r.URL.Query().Get("deviceId"),
		Platform:    r.URL.Query().Get("platform"),
		NetworkUsed: r.URL.Query().Get("networkUsed"),
		Ips:         xhttp.GetRequestIP(r),
		Headers:     headers,
	}
	code, err := s.beforeConnect(r.Context(), param)
	if err != nil {
		logger.Errorf("beforeConnect error: %v, code:", err, code)
		return
	}
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Errorf("failed to accept websocket connection: %v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	ctx, cancelFunc := context.WithCancel(r.Context())
	//ctx := c.CloseRead(r.Context())
	userConn := &types.UserConn{
		Conn: &userConn{
			ws: c,
		},
		ConnParam:   param,
		Ctx:         ctx,
		ConnectedAt: time.Now(),
	}
	go s.loopRead(ctx, cancelFunc, userConn)
	err = s.subscribe(ctx, userConn)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		logger.Errorf("failed to subscribe: %v", err)
		return
	}
}

func (s *Server) subscribe(ctx context.Context, c *types.UserConn) error {
	s.addSubscriber(c)
	defer s.deleteSubscriber(c)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)
}

func (s *Server) loopRead(ctx context.Context, cancelFunc context.CancelFunc, conn *types.UserConn) {
	defer cancelFunc()
	for {
		logx.WithContext(ctx).Infof("start read")
		typ, msg, err := conn.Conn.Read(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// 正常关闭
			} else if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
				websocket.CloseStatus(err) == websocket.StatusGoingAway {
				// 正常关闭
			} else {
				logx.Errorf("failed to read message: %v", err)
			}
			return
		}
		logx.WithContext(ctx).Infof("read message.length: %d", len(msg))
		go xtrace.RunWithTrace("", "ReadFromConn", func(ctx context.Context) {
			s.onReceive(ctx, conn, typ, msg)
		}, propagation.MapCarrier{
			"length":      strconv.Itoa(len(msg)),
			"userId":      conn.ConnParam.UserId,
			"platform":    conn.ConnParam.Platform,
			"deviceId":    conn.ConnParam.DeviceId,
			"ips":         conn.ConnParam.Ips,
			"networkUsed": conn.ConnParam.NetworkUsed,
		})
	}
}
