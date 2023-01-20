// Code generated by goctl. DO NOT EDIT!
// Source: group.proto

package groupservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApplyToBeGroupMemberReq                        = pb.ApplyToBeGroupMemberReq
	ApplyToBeGroupMemberResp                       = pb.ApplyToBeGroupMemberResp
	CreateGroupNoticeReq                           = pb.CreateGroupNoticeReq
	CreateGroupNoticeResp                          = pb.CreateGroupNoticeResp
	CreateGroupReq                                 = pb.CreateGroupReq
	CreateGroupResp                                = pb.CreateGroupResp
	DeleteGroupNoticeReq                           = pb.DeleteGroupNoticeReq
	DeleteGroupNoticeResp                          = pb.DeleteGroupNoticeResp
	EditGroupInfoReq                               = pb.EditGroupInfoReq
	EditGroupInfoResp                              = pb.EditGroupInfoResp
	EditGroupNoticeReq                             = pb.EditGroupNoticeReq
	EditGroupNoticeResp                            = pb.EditGroupNoticeResp
	GetGroupHomeReq                                = pb.GetGroupHomeReq
	GetGroupHomeResp                               = pb.GetGroupHomeResp
	GetGroupHomeResp_MemberStatistics              = pb.GetGroupHomeResp_MemberStatistics
	GetGroupMemberInfoReq                          = pb.GetGroupMemberInfoReq
	GetGroupMemberInfoResp                         = pb.GetGroupMemberInfoResp
	GetGroupMemberListReq                          = pb.GetGroupMemberListReq
	GetGroupMemberListReq_GetGroupMemberListFilter = pb.GetGroupMemberListReq_GetGroupMemberListFilter
	GetGroupMemberListReq_GetGroupMemberListOpt    = pb.GetGroupMemberListReq_GetGroupMemberListOpt
	GetGroupMemberListResp                         = pb.GetGroupMemberListResp
	GetGroupMemberListResp_GroupMember             = pb.GetGroupMemberListResp_GroupMember
	GetGroupNoticeListReq                          = pb.GetGroupNoticeListReq
	GetGroupNoticeListResp                         = pb.GetGroupNoticeListResp
	GetMyGroupListReq                              = pb.GetMyGroupListReq
	GetMyGroupListReq_Filter                       = pb.GetMyGroupListReq_Filter
	GetMyGroupListResp                             = pb.GetMyGroupListResp
	GroupBaseInfo                                  = pb.GroupBaseInfo
	GroupMemberInfo                                = pb.GroupMemberInfo
	GroupNotice                                    = pb.GroupNotice
	HandleGroupApplyReq                            = pb.HandleGroupApplyReq
	HandleGroupApplyResp                           = pb.HandleGroupApplyResp
	InviteFriendToGroupReq                         = pb.InviteFriendToGroupReq
	InviteFriendToGroupResp                        = pb.InviteFriendToGroupResp
	KickGroupMemberReq                             = pb.KickGroupMemberReq
	KickGroupMemberResp                            = pb.KickGroupMemberResp
	MapGroupByIdsReq                               = pb.MapGroupByIdsReq
	MapGroupByIdsResp                              = pb.MapGroupByIdsResp
	SetGroupMemberInfoReq                          = pb.SetGroupMemberInfoReq
	SetGroupMemberInfoResp                         = pb.SetGroupMemberInfoResp
	SyncGroupMemberCountReq                        = pb.SyncGroupMemberCountReq
	SyncGroupMemberCountResp                       = pb.SyncGroupMemberCountResp
	TransferGroupOwnerReq                          = pb.TransferGroupOwnerReq
	TransferGroupOwnerResp                         = pb.TransferGroupOwnerResp

	GroupService interface {
		// CreateGroup 创建群聊
		CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
		// GetGroupHome 获取群聊首页
		GetGroupHome(ctx context.Context, in *GetGroupHomeReq, opts ...grpc.CallOption) (*GetGroupHomeResp, error)
		// InviteFriendToGroup 邀请好友加入群聊
		InviteFriendToGroup(ctx context.Context, in *InviteFriendToGroupReq, opts ...grpc.CallOption) (*InviteFriendToGroupResp, error)
		// CreateGroupNotice 创建群公告
		CreateGroupNotice(ctx context.Context, in *CreateGroupNoticeReq, opts ...grpc.CallOption) (*CreateGroupNoticeResp, error)
		// DeleteGroupNotice 删除群公告
		DeleteGroupNotice(ctx context.Context, in *DeleteGroupNoticeReq, opts ...grpc.CallOption) (*DeleteGroupNoticeResp, error)
		// GetGroupNoticeList 获取群公告列表
		GetGroupNoticeList(ctx context.Context, in *GetGroupNoticeListReq, opts ...grpc.CallOption) (*GetGroupNoticeListResp, error)
		// SetGroupMemberInfo 设置群成员信息
		SetGroupMemberInfo(ctx context.Context, in *SetGroupMemberInfoReq, opts ...grpc.CallOption) (*SetGroupMemberInfoResp, error)
		// GetGroupMemberInfo 获取群成员信息
		GetGroupMemberInfo(ctx context.Context, in *GetGroupMemberInfoReq, opts ...grpc.CallOption) (*GetGroupMemberInfoResp, error)
		// EditGroupInfo 编辑群信息
		EditGroupInfo(ctx context.Context, in *EditGroupInfoReq, opts ...grpc.CallOption) (*EditGroupInfoResp, error)
		// TransferGroupOwner 转让群主
		TransferGroupOwner(ctx context.Context, in *TransferGroupOwnerReq, opts ...grpc.CallOption) (*TransferGroupOwnerResp, error)
		// KickGroupMember 踢出群成员
		KickGroupMember(ctx context.Context, in *KickGroupMemberReq, opts ...grpc.CallOption) (*KickGroupMemberResp, error)
		// GetGroupMemberList 获取群成员列表
		GetGroupMemberList(ctx context.Context, in *GetGroupMemberListReq, opts ...grpc.CallOption) (*GetGroupMemberListResp, error)
		// GetMyGroupList 获取我的群聊列表
		GetMyGroupList(ctx context.Context, in *GetMyGroupListReq, opts ...grpc.CallOption) (*GetMyGroupListResp, error)
		// MapGroupByIds 获取群聊信息
		MapGroupByIds(ctx context.Context, in *MapGroupByIdsReq, opts ...grpc.CallOption) (*MapGroupByIdsResp, error)
		// SyncGroupMemberCount 同步群成员数量
		SyncGroupMemberCount(ctx context.Context, in *SyncGroupMemberCountReq, opts ...grpc.CallOption) (*SyncGroupMemberCountResp, error)
		// ApplyToBeGroupMember 申请加入群聊
		ApplyToBeGroupMember(ctx context.Context, in *ApplyToBeGroupMemberReq, opts ...grpc.CallOption) (*ApplyToBeGroupMemberResp, error)
		// HandleGroupApply 处理群聊申请
		HandleGroupApply(ctx context.Context, in *HandleGroupApplyReq, opts ...grpc.CallOption) (*HandleGroupApplyResp, error)
	}

	defaultGroupService struct {
		cli zrpc.Client
	}
)

func NewGroupService(cli zrpc.Client) GroupService {
	return &defaultGroupService{
		cli: cli,
	}
}

// CreateGroup 创建群聊
func (m *defaultGroupService) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

// GetGroupHome 获取群聊首页
func (m *defaultGroupService) GetGroupHome(ctx context.Context, in *GetGroupHomeReq, opts ...grpc.CallOption) (*GetGroupHomeResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupHome(ctx, in, opts...)
}

// InviteFriendToGroup 邀请好友加入群聊
func (m *defaultGroupService) InviteFriendToGroup(ctx context.Context, in *InviteFriendToGroupReq, opts ...grpc.CallOption) (*InviteFriendToGroupResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.InviteFriendToGroup(ctx, in, opts...)
}

// CreateGroupNotice 创建群公告
func (m *defaultGroupService) CreateGroupNotice(ctx context.Context, in *CreateGroupNoticeReq, opts ...grpc.CallOption) (*CreateGroupNoticeResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.CreateGroupNotice(ctx, in, opts...)
}

// DeleteGroupNotice 删除群公告
func (m *defaultGroupService) DeleteGroupNotice(ctx context.Context, in *DeleteGroupNoticeReq, opts ...grpc.CallOption) (*DeleteGroupNoticeResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.DeleteGroupNotice(ctx, in, opts...)
}

// GetGroupNoticeList 获取群公告列表
func (m *defaultGroupService) GetGroupNoticeList(ctx context.Context, in *GetGroupNoticeListReq, opts ...grpc.CallOption) (*GetGroupNoticeListResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupNoticeList(ctx, in, opts...)
}

// SetGroupMemberInfo 设置群成员信息
func (m *defaultGroupService) SetGroupMemberInfo(ctx context.Context, in *SetGroupMemberInfoReq, opts ...grpc.CallOption) (*SetGroupMemberInfoResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.SetGroupMemberInfo(ctx, in, opts...)
}

// GetGroupMemberInfo 获取群成员信息
func (m *defaultGroupService) GetGroupMemberInfo(ctx context.Context, in *GetGroupMemberInfoReq, opts ...grpc.CallOption) (*GetGroupMemberInfoResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupMemberInfo(ctx, in, opts...)
}

// EditGroupInfo 编辑群信息
func (m *defaultGroupService) EditGroupInfo(ctx context.Context, in *EditGroupInfoReq, opts ...grpc.CallOption) (*EditGroupInfoResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.EditGroupInfo(ctx, in, opts...)
}

// TransferGroupOwner 转让群主
func (m *defaultGroupService) TransferGroupOwner(ctx context.Context, in *TransferGroupOwnerReq, opts ...grpc.CallOption) (*TransferGroupOwnerResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.TransferGroupOwner(ctx, in, opts...)
}

// KickGroupMember 踢出群成员
func (m *defaultGroupService) KickGroupMember(ctx context.Context, in *KickGroupMemberReq, opts ...grpc.CallOption) (*KickGroupMemberResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.KickGroupMember(ctx, in, opts...)
}

// GetGroupMemberList 获取群成员列表
func (m *defaultGroupService) GetGroupMemberList(ctx context.Context, in *GetGroupMemberListReq, opts ...grpc.CallOption) (*GetGroupMemberListResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupMemberList(ctx, in, opts...)
}

// GetMyGroupList 获取我的群聊列表
func (m *defaultGroupService) GetMyGroupList(ctx context.Context, in *GetMyGroupListReq, opts ...grpc.CallOption) (*GetMyGroupListResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.GetMyGroupList(ctx, in, opts...)
}

// MapGroupByIds 获取群聊信息
func (m *defaultGroupService) MapGroupByIds(ctx context.Context, in *MapGroupByIdsReq, opts ...grpc.CallOption) (*MapGroupByIdsResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.MapGroupByIds(ctx, in, opts...)
}

// SyncGroupMemberCount 同步群成员数量
func (m *defaultGroupService) SyncGroupMemberCount(ctx context.Context, in *SyncGroupMemberCountReq, opts ...grpc.CallOption) (*SyncGroupMemberCountResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.SyncGroupMemberCount(ctx, in, opts...)
}

// ApplyToBeGroupMember 申请加入群聊
func (m *defaultGroupService) ApplyToBeGroupMember(ctx context.Context, in *ApplyToBeGroupMemberReq, opts ...grpc.CallOption) (*ApplyToBeGroupMemberResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.ApplyToBeGroupMember(ctx, in, opts...)
}

// HandleGroupApply 处理群聊申请
func (m *defaultGroupService) HandleGroupApply(ctx context.Context, in *HandleGroupApplyReq, opts ...grpc.CallOption) (*HandleGroupApplyResp, error) {
	client := pb.NewGroupServiceClient(m.cli.Conn())
	return client.HandleGroupApply(ctx, in, opts...)
}
