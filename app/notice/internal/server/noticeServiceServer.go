// Code generated by goctl. DO NOT EDIT!
// Source: notice.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/notice/internal/logic"
	"github.com/cherish-chat/xxim-server/app/notice/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type NoticeServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedNoticeServiceServer
}

func NewNoticeServiceServer(svcCtx *svc.ServiceContext) *NoticeServiceServer {
	return &NoticeServiceServer{
		svcCtx: svcCtx,
	}
}

// AfterConnect conn hook
func (s *NoticeServiceServer) AfterConnect(ctx context.Context, in *pb.AfterConnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterConnectLogic(ctx, s.svcCtx)
	return l.AfterConnect(in)
}

// AfterDisconnect conn hook
func (s *NoticeServiceServer) AfterDisconnect(ctx context.Context, in *pb.AfterDisconnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterDisconnectLogic(ctx, s.svcCtx)
	return l.AfterDisconnect(in)
}

// SendNoticeData 发送通知数据
func (s *NoticeServiceServer) SendNoticeData(ctx context.Context, in *pb.SendNoticeDataReq) (*pb.SendNoticeDataResp, error) {
	l := logic.NewSendNoticeDataLogic(ctx, s.svcCtx)
	return l.SendNoticeData(in)
}

// PushNoticeData 推送通知数据
func (s *NoticeServiceServer) PushNoticeData(ctx context.Context, in *pb.PushNoticeDataReq) (*pb.PushNoticeDataResp, error) {
	l := logic.NewPushNoticeDataLogic(ctx, s.svcCtx)
	return l.PushNoticeData(in)
}

// GetUserNoticeData 获取用户通知数据
func (s *NoticeServiceServer) GetUserNoticeData(ctx context.Context, in *pb.GetUserNoticeDataReq) (*pb.GetUserNoticeDataResp, error) {
	l := logic.NewGetUserNoticeDataLogic(ctx, s.svcCtx)
	return l.GetUserNoticeData(in)
}

// AckNoticeData 确认通知数据
func (s *NoticeServiceServer) AckNoticeData(ctx context.Context, in *pb.AckNoticeDataReq) (*pb.AckNoticeDataResp, error) {
	l := logic.NewAckNoticeDataLogic(ctx, s.svcCtx)
	return l.AckNoticeData(in)
}

// GetUserNoticeConvIds 获取用户所有的通知号
func (s *NoticeServiceServer) GetUserNoticeConvIds(ctx context.Context, in *pb.GetUserNoticeConvIdsReq) (*pb.GetUserNoticeConvIdsResp, error) {
	l := logic.NewGetUserNoticeConvIdsLogic(ctx, s.svcCtx)
	return l.GetUserNoticeConvIds(in)
}

// GetNoticeConvAllSubscribers 获取通知号所有的订阅者
func (s *NoticeServiceServer) GetNoticeConvAllSubscribers(ctx context.Context, in *pb.GetNoticeConvAllSubscribersReq) (*pb.GetNoticeConvAllSubscribersResp, error) {
	l := logic.NewGetNoticeConvAllSubscribersLogic(ctx, s.svcCtx)
	return l.GetNoticeConvAllSubscribers(in)
}

// SetUserSubscriptions 设置用户订阅
func (s *NoticeServiceServer) SetUserSubscriptions(ctx context.Context, in *pb.SetUserSubscriptionsReq) (*pb.CommonResp, error) {
	l := logic.NewSetUserSubscriptionsLogic(ctx, s.svcCtx)
	return l.SetUserSubscriptions(in)
}
