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

func (l *SendMsgListSyncLogic) SendMsgListSync(in *pb.SendMsgListReq) (*pb.CommonResp, error) {
	var resp *pb.CommonResp
	var err error
	xtrace.StartFuncSpan(l.ctx, "InsertMsgDataListLogic.InsertMsgDataList", func(ctx context.Context) {
		resp, err = NewInsertMsgDataListLogic(ctx, l.svcCtx).InsertMsgDataList(&pb.MsgDataList{MsgDataList: in.MsgDataList})
	})
	if err != nil {
		l.Errorf("NewInsertMsgDataListLogic.InsertMsgDataList error: %v", err)
		return resp, err
	}
	// 推送给相关的在线用户
	xtrace.StartFuncSpan(l.ctx, "PushMsgListLogic.PushMsgList", func(ctx context.Context) {

	})
	return &pb.CommonResp{}, nil
}
