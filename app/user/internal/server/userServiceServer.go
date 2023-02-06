// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/user/internal/logic"
	"github.com/cherish-chat/xxim-server/app/user/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) ConfirmRegister(ctx context.Context, in *pb.ConfirmRegisterReq) (*pb.ConfirmRegisterResp, error) {
	l := logic.NewConfirmRegisterLogic(ctx, s.svcCtx)
	return l.ConfirmRegister(in)
}

func (s *UserServiceServer) MapUserByIds(ctx context.Context, in *pb.MapUserByIdsReq) (*pb.MapUserByIdsResp, error) {
	l := logic.NewMapUserByIdsLogic(ctx, s.svcCtx)
	return l.MapUserByIds(in)
}

func (s *UserServiceServer) BatchGetUserBaseInfo(ctx context.Context, in *pb.BatchGetUserBaseInfoReq) (*pb.BatchGetUserBaseInfoResp, error) {
	l := logic.NewBatchGetUserBaseInfoLogic(ctx, s.svcCtx)
	return l.BatchGetUserBaseInfo(in)
}

func (s *UserServiceServer) SearchUsersByKeyword(ctx context.Context, in *pb.SearchUsersByKeywordReq) (*pb.SearchUsersByKeywordResp, error) {
	l := logic.NewSearchUsersByKeywordLogic(ctx, s.svcCtx)
	return l.SearchUsersByKeyword(in)
}

func (s *UserServiceServer) GetUserHome(ctx context.Context, in *pb.GetUserHomeReq) (*pb.GetUserHomeResp, error) {
	l := logic.NewGetUserHomeLogic(ctx, s.svcCtx)
	return l.GetUserHome(in)
}

func (s *UserServiceServer) GetUserSettings(ctx context.Context, in *pb.GetUserSettingsReq) (*pb.GetUserSettingsResp, error) {
	l := logic.NewGetUserSettingsLogic(ctx, s.svcCtx)
	return l.GetUserSettings(in)
}

func (s *UserServiceServer) SetUserSettings(ctx context.Context, in *pb.SetUserSettingsReq) (*pb.SetUserSettingsResp, error) {
	l := logic.NewSetUserSettingsLogic(ctx, s.svcCtx)
	return l.SetUserSettings(in)
}

// AfterConnect conn hook
func (s *UserServiceServer) AfterConnect(ctx context.Context, in *pb.AfterConnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterConnectLogic(ctx, s.svcCtx)
	return l.AfterConnect(in)
}

// AfterDisconnect conn hook
func (s *UserServiceServer) AfterDisconnect(ctx context.Context, in *pb.AfterDisconnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterDisconnectLogic(ctx, s.svcCtx)
	return l.AfterDisconnect(in)
}

func (s *UserServiceServer) BatchGetUserAllDevices(ctx context.Context, in *pb.BatchGetUserAllDevicesReq) (*pb.BatchGetUserAllDevicesResp, error) {
	l := logic.NewBatchGetUserAllDevicesLogic(ctx, s.svcCtx)
	return l.BatchGetUserAllDevices(in)
}

func (s *UserServiceServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UserServiceServer) GetAllUserInvitationCode(ctx context.Context, in *pb.GetAllUserInvitationCodeReq) (*pb.GetAllUserInvitationCodeResp, error) {
	l := logic.NewGetAllUserInvitationCodeLogic(ctx, s.svcCtx)
	return l.GetAllUserInvitationCode(in)
}

func (s *UserServiceServer) GetUserInvitationCodeDetail(ctx context.Context, in *pb.GetUserInvitationCodeDetailReq) (*pb.GetUserInvitationCodeDetailResp, error) {
	l := logic.NewGetUserInvitationCodeDetailLogic(ctx, s.svcCtx)
	return l.GetUserInvitationCodeDetail(in)
}

func (s *UserServiceServer) AddUserInvitationCode(ctx context.Context, in *pb.AddUserInvitationCodeReq) (*pb.AddUserInvitationCodeResp, error) {
	l := logic.NewAddUserInvitationCodeLogic(ctx, s.svcCtx)
	return l.AddUserInvitationCode(in)
}

func (s *UserServiceServer) UpdateUserInvitationCode(ctx context.Context, in *pb.UpdateUserInvitationCodeReq) (*pb.UpdateUserInvitationCodeResp, error) {
	l := logic.NewUpdateUserInvitationCodeLogic(ctx, s.svcCtx)
	return l.UpdateUserInvitationCode(in)
}

func (s *UserServiceServer) DeleteUserInvitationCode(ctx context.Context, in *pb.DeleteUserInvitationCodeReq) (*pb.DeleteUserInvitationCodeResp, error) {
	l := logic.NewDeleteUserInvitationCodeLogic(ctx, s.svcCtx)
	return l.DeleteUserInvitationCode(in)
}

func (s *UserServiceServer) GetAllUserIpWhiteList(ctx context.Context, in *pb.GetAllUserIpWhiteListReq) (*pb.GetAllUserIpWhiteListResp, error) {
	l := logic.NewGetAllUserIpWhiteListLogic(ctx, s.svcCtx)
	return l.GetAllUserIpWhiteList(in)
}

func (s *UserServiceServer) GetUserIpWhiteListDetail(ctx context.Context, in *pb.GetUserIpWhiteListDetailReq) (*pb.GetUserIpWhiteListDetailResp, error) {
	l := logic.NewGetUserIpWhiteListDetailLogic(ctx, s.svcCtx)
	return l.GetUserIpWhiteListDetail(in)
}

func (s *UserServiceServer) AddUserIpWhiteList(ctx context.Context, in *pb.AddUserIpWhiteListReq) (*pb.AddUserIpWhiteListResp, error) {
	l := logic.NewAddUserIpWhiteListLogic(ctx, s.svcCtx)
	return l.AddUserIpWhiteList(in)
}

func (s *UserServiceServer) UpdateUserIpWhiteList(ctx context.Context, in *pb.UpdateUserIpWhiteListReq) (*pb.UpdateUserIpWhiteListResp, error) {
	l := logic.NewUpdateUserIpWhiteListLogic(ctx, s.svcCtx)
	return l.UpdateUserIpWhiteList(in)
}

func (s *UserServiceServer) DeleteUserIpWhiteList(ctx context.Context, in *pb.DeleteUserIpWhiteListReq) (*pb.DeleteUserIpWhiteListResp, error) {
	l := logic.NewDeleteUserIpWhiteListLogic(ctx, s.svcCtx)
	return l.DeleteUserIpWhiteList(in)
}

func (s *UserServiceServer) GetAllUserIpBlackList(ctx context.Context, in *pb.GetAllUserIpBlackListReq) (*pb.GetAllUserIpBlackListResp, error) {
	l := logic.NewGetAllUserIpBlackListLogic(ctx, s.svcCtx)
	return l.GetAllUserIpBlackList(in)
}

func (s *UserServiceServer) GetUserIpBlackListDetail(ctx context.Context, in *pb.GetUserIpBlackListDetailReq) (*pb.GetUserIpBlackListDetailResp, error) {
	l := logic.NewGetUserIpBlackListDetailLogic(ctx, s.svcCtx)
	return l.GetUserIpBlackListDetail(in)
}

func (s *UserServiceServer) AddUserIpBlackList(ctx context.Context, in *pb.AddUserIpBlackListReq) (*pb.AddUserIpBlackListResp, error) {
	l := logic.NewAddUserIpBlackListLogic(ctx, s.svcCtx)
	return l.AddUserIpBlackList(in)
}

func (s *UserServiceServer) UpdateUserIpBlackList(ctx context.Context, in *pb.UpdateUserIpBlackListReq) (*pb.UpdateUserIpBlackListResp, error) {
	l := logic.NewUpdateUserIpBlackListLogic(ctx, s.svcCtx)
	return l.UpdateUserIpBlackList(in)
}

func (s *UserServiceServer) DeleteUserIpBlackList(ctx context.Context, in *pb.DeleteUserIpBlackListReq) (*pb.DeleteUserIpBlackListResp, error) {
	l := logic.NewDeleteUserIpBlackListLogic(ctx, s.svcCtx)
	return l.DeleteUserIpBlackList(in)
}

func (s *UserServiceServer) GetAllUserDefaultConv(ctx context.Context, in *pb.GetAllUserDefaultConvReq) (*pb.GetAllUserDefaultConvResp, error) {
	l := logic.NewGetAllUserDefaultConvLogic(ctx, s.svcCtx)
	return l.GetAllUserDefaultConv(in)
}

func (s *UserServiceServer) GetUserDefaultConvDetail(ctx context.Context, in *pb.GetUserDefaultConvDetailReq) (*pb.GetUserDefaultConvDetailResp, error) {
	l := logic.NewGetUserDefaultConvDetailLogic(ctx, s.svcCtx)
	return l.GetUserDefaultConvDetail(in)
}

func (s *UserServiceServer) AddUserDefaultConv(ctx context.Context, in *pb.AddUserDefaultConvReq) (*pb.AddUserDefaultConvResp, error) {
	l := logic.NewAddUserDefaultConvLogic(ctx, s.svcCtx)
	return l.AddUserDefaultConv(in)
}

func (s *UserServiceServer) UpdateUserDefaultConv(ctx context.Context, in *pb.UpdateUserDefaultConvReq) (*pb.UpdateUserDefaultConvResp, error) {
	l := logic.NewUpdateUserDefaultConvLogic(ctx, s.svcCtx)
	return l.UpdateUserDefaultConv(in)
}

func (s *UserServiceServer) DeleteUserDefaultConv(ctx context.Context, in *pb.DeleteUserDefaultConvReq) (*pb.DeleteUserDefaultConvResp, error) {
	l := logic.NewDeleteUserDefaultConvLogic(ctx, s.svcCtx)
	return l.DeleteUserDefaultConv(in)
}

func (s *UserServiceServer) GetAllUserModel(ctx context.Context, in *pb.GetAllUserModelReq) (*pb.GetAllUserModelResp, error) {
	l := logic.NewGetAllUserModelLogic(ctx, s.svcCtx)
	return l.GetAllUserModel(in)
}

func (s *UserServiceServer) GetUserModelDetail(ctx context.Context, in *pb.GetUserModelDetailReq) (*pb.GetUserModelDetailResp, error) {
	l := logic.NewGetUserModelDetailLogic(ctx, s.svcCtx)
	return l.GetUserModelDetail(in)
}

func (s *UserServiceServer) AddUserModel(ctx context.Context, in *pb.AddUserModelReq) (*pb.AddUserModelResp, error) {
	l := logic.NewAddUserModelLogic(ctx, s.svcCtx)
	return l.AddUserModel(in)
}

func (s *UserServiceServer) UpdateUserModel(ctx context.Context, in *pb.UpdateUserModelReq) (*pb.UpdateUserModelResp, error) {
	l := logic.NewUpdateUserModelLogic(ctx, s.svcCtx)
	return l.UpdateUserModel(in)
}

func (s *UserServiceServer) DeleteUserModel(ctx context.Context, in *pb.DeleteUserModelReq) (*pb.DeleteUserModelResp, error) {
	l := logic.NewDeleteUserModelLogic(ctx, s.svcCtx)
	return l.DeleteUserModel(in)
}

func (s *UserServiceServer) SwitchUserModel(ctx context.Context, in *pb.SwitchUserModelReq) (*pb.SwitchUserModelResp, error) {
	l := logic.NewSwitchUserModelLogic(ctx, s.svcCtx)
	return l.SwitchUserModel(in)
}

func (s *UserServiceServer) GetAllLoginRecord(ctx context.Context, in *pb.GetAllLoginRecordReq) (*pb.GetAllLoginRecordResp, error) {
	l := logic.NewGetAllLoginRecordLogic(ctx, s.svcCtx)
	return l.GetAllLoginRecord(in)
}
