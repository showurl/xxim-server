package svc

import (
	"github.com/cherish-chat/xxim-server/app/group/groupservice"
	"github.com/cherish-chat/xxim-server/app/im/imservice"
	"github.com/cherish-chat/xxim-server/app/mgmt/internal/config"
	msgservice "github.com/cherish-chat/xxim-server/app/msg/msgService"
	"github.com/cherish-chat/xxim-server/app/notice/noticeservice"
	"github.com/cherish-chat/xxim-server/app/relation/relationservice"
	"github.com/cherish-chat/xxim-server/app/user/userservice"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	redis           *redis.Redis
	imService       imservice.ImService
	msgService      msgservice.MsgService
	noticeService   noticeservice.NoticeService
	relationService relationservice.RelationService
	userService     userservice.UserService
	groupService    groupservice.GroupService
}

func NewServiceContext(c config.Config, rc *redis.Redis) *ServiceContext {
	return &ServiceContext{
		Config: c,
		redis:  rc,
	}
}

func (s *ServiceContext) Redis() *redis.Redis {
	return s.redis
}

func (s *ServiceContext) ImService() imservice.ImService {
	if s.imService == nil {
		s.imService = imservice.NewImService(zrpc.MustNewClient(s.Config.ImRpc))
	}
	return s.imService
}

func (s *ServiceContext) MsgService() msgservice.MsgService {
	if s.msgService == nil {
		s.msgService = msgservice.NewMsgService(zrpc.MustNewClient(s.Config.MsgRpc))
	}
	return s.msgService
}

func (s *ServiceContext) NoticeService() noticeservice.NoticeService {
	if s.noticeService == nil {
		s.noticeService = noticeservice.NewNoticeService(zrpc.MustNewClient(s.Config.NoticeRpc))
	}
	return s.noticeService
}

func (s *ServiceContext) RelationService() relationservice.RelationService {
	if s.relationService == nil {
		s.relationService = relationservice.NewRelationService(zrpc.MustNewClient(s.Config.RelationRpc))
	}
	return s.relationService
}

func (s *ServiceContext) UserService() userservice.UserService {
	if s.userService == nil {
		s.userService = userservice.NewUserService(zrpc.MustNewClient(s.Config.UserRpc))
	}
	return s.userService
}

func (s *ServiceContext) GroupService() groupservice.GroupService {
	if s.groupService == nil {
		s.groupService = groupservice.NewGroupService(zrpc.MustNewClient(s.Config.GroupRpc))
	}
	return s.groupService
}
