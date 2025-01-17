package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/common/xtrace"

	"github.com/cherish-chat/xxim-server/app/msg/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgListSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgListSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgListSyncLogic {
	return &SendMsgListSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgListSyncLogic) SendMsgListSync(in *pb.SendMsgListReq) (*pb.SendMsgListResp, error) {
	var resp *pb.MsgDataList
	var err error
	xtrace.StartFuncSpan(l.ctx, "InsertMsgDataList", func(ctx context.Context) {
		resp, err = NewInsertMsgDataListLogic(ctx, l.svcCtx).InsertMsgDataList(&pb.MsgDataList{MsgDataList: in.MsgDataList})
	})
	if err != nil {
		l.Errorf("InsertMsgDataList error: %v", err)
		return &pb.SendMsgListResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	// 推送给相关的在线用户
	go xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "PushMsgList", func(ctx context.Context) {
		_, err := NewPushMsgListLogic(ctx, l.svcCtx).PushMsgList(&pb.PushMsgListReq{MsgDataList: resp.MsgDataList})
		if err != nil {
			l.Errorf("PushMsgList error: %v", err)
		}
	}, nil)
	return &pb.SendMsgListResp{}, nil
}
