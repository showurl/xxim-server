package logic

import (
	"context"
	msgservice "github.com/cherish-chat/xxim-server/app/msg/msgService"
	"github.com/cherish-chat/xxim-server/app/msg/msgmodel"
	"github.com/cherish-chat/xxim-server/app/notice/noticemodel"
	"github.com/cherish-chat/xxim-server/app/relation/relationmodel"
	"github.com/cherish-chat/xxim-server/app/user/usermodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"
	"github.com/cherish-chat/xxim-server/common/xtrace"
	"gorm.io/gorm"
	"time"

	"github.com/cherish-chat/xxim-server/app/relation/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptAddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAcceptAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptAddFriendLogic {
	return &AcceptAddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AcceptAddFriendLogic) AcceptAddFriend(in *pb.AcceptAddFriendReq) (*pb.AcceptAddFriendResp, error) {
	// 我的好友总数是否已达上限
	{
		var getFriendCountResp *pb.GetFriendCountResp
		var err error
		xtrace.StartFuncSpan(l.ctx, "GetFriendCount", func(ctx context.Context) {
			getFriendCountResp, err = NewGetFriendCountLogic(ctx, l.svcCtx).GetFriendCount(&pb.GetFriendCountReq{
				CommonReq: in.CommonReq,
			})
		})
		if err != nil {
			l.Errorf("GetFriendCount failed, err: %v", err)
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		if int64(getFriendCountResp.Count) >= l.svcCtx.ConfigMgr.FriendMaxCount(l.ctx, in.CommonReq.UserId) {
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewToastErrorResp(l.svcCtx.T(in.CommonReq.Language, "好友数量已达上限"))}, nil
		}
	}
	now := time.Now().UnixMilli()
	friend1 := &relationmodel.Friend{FriendId: in.CommonReq.UserId, UserId: in.ApplyUserId, CreateTime: now}
	friend2 := &relationmodel.Friend{FriendId: in.ApplyUserId, UserId: in.CommonReq.UserId, CreateTime: now}
	{
		// 添加好友
		err := xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
			err := xorm.Upsert(tx, friend1, []string{"friendId", "userId"}, []string{"friendId", "userId"})
			if err != nil {
				l.Errorf("Save friend1 failed, err: %v", err)
				return err
			}
			err = xorm.Upsert(tx, friend2, []string{"friendId", "userId"}, []string{"friendId", "userId"})
			if err != nil {
				l.Errorf("Save friend2 failed, err: %v", err)
				return err
			}
			return nil
		}, func(tx *gorm.DB) error {
			for _, userId := range []string{in.CommonReq.UserId, in.ApplyUserId} {
				notice := &noticemodel.Notice{
					ConvId: pb.HiddenConvIdCommand(),
					UserId: userId,
					Options: noticemodel.NoticeOption{
						StorageForClient: false,
						UpdateConvNotice: false,
					},
					ContentType: pb.NoticeContentType_SyncFriendList,
					Content: utils.AnyToBytes(pb.NoticeContent_SyncFriendList{
						Comment: "acceptAddFriend",
					}),
					UniqueId: "syncFriendList",
					Title:    "",
					Ext:      nil,
				}
				err := notice.Insert(l.ctx, tx, l.svcCtx.Redis())
				if err != nil {
					l.Errorf("insert notice failed, err: %v", err)
					return err
				}
			}
			return nil
		})
		if err != nil {
			l.Errorf("InsertOne failed, err: %v", err)
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
	}
	{
		// 设置申请状态
		if in.RequestId != nil {
			err := l.svcCtx.Mysql().Model(&relationmodel.RequestAddFriend{}).
				Where("status = ? AND ((fromUserId = ? AND toUserId = ?) OR (fromUserId = ? AND toUserId = ?))",
					pb.RequestAddFriendStatus_Unhandled,
					in.CommonReq.UserId, in.ApplyUserId,
					in.ApplyUserId, in.CommonReq.UserId).
				Updates(map[string]interface{}{
					"status":     pb.RequestAddFriendStatus_Agreed,
					"updateTime": time.Now().UnixMilli(),
				}).Error
			if err != nil {
				l.Errorf("UpdateOne failed, err: %v", err)
				return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
			}
		}
	}
	{
		// 删除缓存
		err := relationmodel.FlushFriendList(l.ctx, l.svcCtx.Redis(), in.ApplyUserId, in.CommonReq.UserId)
		if err != nil {
			l.Errorf("FlushFriendList failed, err: %v", err)
		}
		// 预热缓存
		xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "CacheWarm", func(ctx context.Context) {
			_, _ = relationmodel.GetMyFriendList(ctx, l.svcCtx.Redis(), l.svcCtx.Mysql(), in.ApplyUserId)
			_, _ = relationmodel.GetMyFriendList(ctx, l.svcCtx.Redis(), l.svcCtx.Mysql(), in.CommonReq.UserId)
		}, nil)
		// 刷新订阅
		utils.RetryProxy(context.Background(), 12, 1*time.Second, func() error {
			_, err := l.svcCtx.MsgService().FlushUsersSubConv(l.ctx, &pb.FlushUsersSubConvReq{UserIds: []string{
				friend1.UserId, friend1.FriendId,
			}})
			if err != nil {
				l.Errorf("FlushUsersSubConv failed, err: %v", err)
				return err
			}
			for _, userId := range []string{friend1.UserId, friend1.FriendId} {
				_, err = l.svcCtx.NoticeService().GetUserNoticeData(l.ctx, &pb.GetUserNoticeDataReq{
					UserId: userId,
					ConvId: pb.HiddenConvIdCommand(),
				})
				if err != nil {
					l.Errorf("SendNoticeData failed, err: %v", err)
					return err
				}
			}
			return err
		})
		// 接受者发送消息：我们已经是好友了，快来聊天吧
		l.sendMsg(in)
	}
	return &pb.AcceptAddFriendResp{}, nil
}

func (l *AcceptAddFriendLogic) sendMsg(in *pb.AcceptAddFriendReq) {
	go xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "SendMsg", func(ctx context.Context) {
		// 获取接受者info
		userByIds, err := l.svcCtx.UserService().MapUserByIds(ctx, &pb.MapUserByIdsReq{Ids: []string{in.CommonReq.UserId}})
		if err != nil {
			l.Errorf("MapUserByIds failed, err: %v", err)
		} else {
			selfInfo, ok := userByIds.Users[in.CommonReq.UserId]
			if ok {
				self := usermodel.UserFromBytes(selfInfo)
				text := "我们已经是好友了，快来聊天吧"
				if in.SendTextMsg != nil && *in.SendTextMsg != "" {
					text = *in.SendTextMsg
				}
				_, err = msgservice.SendMsgSync(l.svcCtx.MsgService(), ctx, []*pb.MsgData{
					msgmodel.CreateTextMsgToUser(
						&pb.UserBaseInfo{
							Id:       self.Id,
							Nickname: self.Nickname,
							Avatar:   self.Avatar,
							Xb:       self.Xb,
							Birthday: self.Birthday,
							Role:     int32(self.Role),
						},
						in.ApplyUserId,
						l.svcCtx.T(in.CommonReq.Language, text),
						msgmodel.MsgOptions{
							OfflinePush:       true,
							StorageForServer:  true,
							StorageForClient:  true,
							UpdateUnreadCount: true,
							NeedDecrypt:       false,
							UpdateConvMsg:     true,
						},
						&msgmodel.MsgOfflinePush{
							Title:   self.Nickname,
							Content: text,
							Payload: "",
						},
						nil,
					).ToMsgData(),
				})
				if err != nil {
					l.Errorf("SendMsgSync failed, err: %v", err)
					err = nil
				}
			}
		}
	}, nil)
}
