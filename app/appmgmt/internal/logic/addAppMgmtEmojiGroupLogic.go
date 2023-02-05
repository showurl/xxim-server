package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/appmgmt/appmgmtmodel"
	"time"

	"github.com/cherish-chat/xxim-server/app/appmgmt/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAppMgmtEmojiGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAppMgmtEmojiGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAppMgmtEmojiGroupLogic {
	return &AddAppMgmtEmojiGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAppMgmtEmojiGroupLogic) AddAppMgmtEmojiGroup(in *pb.AddAppMgmtEmojiGroupReq) (*pb.AddAppMgmtEmojiGroupResp, error) {
	model := &appmgmtmodel.EmojiGroup{
		Name:       in.AppMgmtEmojiGroup.Name,
		CoverId:    "",
		IsEnable:   in.AppMgmtEmojiGroup.IsEnable,
		CreateTime: time.Now().UnixMilli(),
	}
	err := model.Insert(l.svcCtx.Mysql())
	if err != nil {
		l.Errorf("insert err: %v", err)
		return &pb.AddAppMgmtEmojiGroupResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddAppMgmtEmojiGroupResp{}, nil
}
