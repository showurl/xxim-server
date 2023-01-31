package serverhandler

import (
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"strings"
)

const logoBase64 = `data:image/png;base64,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`
const svcTmpl = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.TotalWidth}}" height="20" role="img" aria-label="xxim: {{.Count}} online"><title>xxim: {{.Count}} online</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="{{.TotalWidth}}" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="50" height="20" fill="#E23F5C"/><rect x="50" width="{{.RightWidth}}" height="20" fill="#FBC200"/><rect width="{{.TotalWidth}}" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><image x="5" y="3.2" width="14" height="14" xlink:href="{{.LogoBase64}}"/><text aria-hidden="true" x="335" y="145" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="290">XXIM</text><text x="335" y="135" transform="scale(.1)" fill="#fff" textLength="290">XXIM</text><text aria-hidden="true" x="{{.textPosition}}" y="145" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{.textLength}}">{{.Count}} online</text><text x="{{.textPosition}}" y="135" transform="scale(.1)" fill="#fff" textLength="{{.textLength}}">{{.Count}} online</text></g></svg>`

func (r *ServerHandler) onlineShield(ctx *gin.Context) {
	logger := logx.WithContext(ctx)
	userConnResp, err := r.svcCtx.ImService().GetUserConn(ctx.Request.Context(), &pb.GetUserConnReq{
		UserIds:   nil,
		Platforms: nil,
		Devices:   nil,
	})
	if err != nil {
		logger.Errorf("get user conn error: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//var userIdMap = make(map[string]bool)
	var deviceIdMap = make(map[string]bool)
	for _, conn := range userConnResp.ConnParams {
		//userIdMap[conn.UserId] = true
		deviceIdMap[conn.DeviceId] = true
	}
	deviceCount := strconv.Itoa(len(deviceIdMap))
	var totalWidth = 106
	var rightWidth = 56
	var textLength = 459
	var textPosition = 780
	totalWidth += (len(deviceCount) - 1) * 5
	rightWidth += (len(deviceCount) - 1) * 5
	textLength += (len(deviceCount) - 1) * 51
	textPosition += (len(deviceCount) - 1) * 25
	resp := strings.ReplaceAll(svcTmpl, "{{.Count}}", deviceCount)
	resp = strings.ReplaceAll(resp, "{{.TotalWidth}}", strconv.Itoa(totalWidth))
	resp = strings.ReplaceAll(resp, "{{.RightWidth}}", strconv.Itoa(rightWidth))
	resp = strings.ReplaceAll(resp, "{{.textLength}}", strconv.Itoa(textLength))
	resp = strings.ReplaceAll(resp, "{{.textPosition}}", strconv.Itoa(textPosition))
	resp = strings.ReplaceAll(resp, "{{.LogoBase64}}", logoBase64)
	// 返回 svg
	ctx.Data(http.StatusOK, "image/svg+xml", []byte(resp))
}
