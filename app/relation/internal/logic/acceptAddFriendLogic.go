package logic

import (
	"context"
	msgservice "github.com/cherish-chat/xxim-server/app/msg/msgService"
	"github.com/cherish-chat/xxim-server/app/msg/msgmodel"
	"github.com/cherish-chat/xxim-server/app/relation/relationmodel"
	"github.com/cherish-chat/xxim-server/app/user/usermodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xtrace"
	"go.mongodb.org/mongo-driver/bson"
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
				Requester: in.Requester,
			})
		})
		if err != nil {
			l.Errorf("GetFriendCount failed, err: %v", err)
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		if int64(getFriendCountResp.Count) >= utils.AnyToInt64(l.svcCtx.SystemConfigMgr.Get("friend_max_count")) {
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewToastErrorResp(l.svcCtx.T(in.Requester.Language, "好友数量已达上限"))}, nil
		}
	}
	{
		// 添加好友
		friend1 := &relationmodel.Friend{FriendId: in.Requester.Id, UserId: in.ApplyUserId}
		friend2 := &relationmodel.Friend{FriendId: in.ApplyUserId, UserId: in.Requester.Id}
		_, err := l.svcCtx.Mongo().Collection(friend1).Upsert(l.ctx, bson.M{
			"userId":   friend1.UserId,
			"friendId": friend1.FriendId,
		}, friend1)
		if err != nil {
			l.Errorf("InsertOne failed, err: %v", err)
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		_, err = l.svcCtx.Mongo().Collection(friend2).Upsert(l.ctx, bson.M{
			"userId":   friend2.UserId,
			"friendId": friend2.FriendId,
		}, friend2)
		if err != nil {
			l.Errorf("InsertOne failed, err: %v", err)
			return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		// 接受者发送消息：我们已经是好友了，快来聊天吧
		go l.sendMsg(in)
	}
	{
		// 设置申请状态
		if in.RequestId != nil {
			apply := &relationmodel.RequestAddFriend{Id: *in.RequestId}
			_, err := l.svcCtx.Mongo().Collection(apply).UpdateAll(l.ctx, bson.M{
				"$or": []bson.M{{
					"fromUserId": in.ApplyUserId,
					"toUserId":   in.Requester.Id,
				}, {
					"fromUserId": in.Requester.Id,
					"toUserId":   in.ApplyUserId,
				}},
				"status": pb.RequestAddFriendStatus_Unhandled,
			}, bson.M{
				"$set": bson.M{
					"status":     pb.RequestAddFriendStatus_Agreed,
					"updateTime": time.Now().UnixMilli(),
				},
			})
			if err != nil {
				l.Errorf("UpdateOne failed, err: %v", err)
				return &pb.AcceptAddFriendResp{CommonResp: pb.NewRetryErrorResp()}, err
			}
		}
	}
	{
		// 删除缓存
		err := relationmodel.FlushFriendList(l.ctx, l.svcCtx.Redis(), in.ApplyUserId, in.Requester.Id)
		if err != nil {
			l.Errorf("FlushFriendList failed, err: %v", err)
		}
		// 预热缓存
		go xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "CacheWarm", func(ctx context.Context) {
			_, _ = relationmodel.GetMyFriendList(ctx, l.svcCtx.Redis(), l.svcCtx.Mongo().Collection(&relationmodel.Friend{}), in.ApplyUserId)
			_, _ = relationmodel.GetMyFriendList(ctx, l.svcCtx.Redis(), l.svcCtx.Mongo().Collection(&relationmodel.Friend{}), in.Requester.Id)
		}, nil)
	}
	return &pb.AcceptAddFriendResp{}, nil
}

func (l *AcceptAddFriendLogic) sendMsg(in *pb.AcceptAddFriendReq) {
	xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "SendMsg", func(ctx context.Context) {
		// 获取接受者info
		userByIds, err := l.svcCtx.UserService().MapUserByIds(ctx, &pb.MapUserByIdsReq{Ids: []string{in.Requester.Id}})
		if err != nil {
			l.Errorf("MapUserByIds failed, err: %v", err)
		} else {
			selfInfo, ok := userByIds.Users[in.Requester.Id]
			if ok {
				self := usermodel.UserFromBytes(selfInfo)
				_, err = msgservice.SendMsgSync(l.svcCtx.MsgService(), ctx, []*pb.MsgData{
					msgmodel.CreateTextMsgToUser(
						&pb.UserBaseInfo{
							Id:       self.Id,
							Nickname: self.Nickname,
							Avatar:   self.Avatar,
							Xb:       self.Xb,
							Birthday: self.Birthday,
						},
						in.ApplyUserId,
						l.svcCtx.T(in.Requester.Language, "我们已经是好友了，快来聊天吧"),
						msgmodel.MsgOptions{
							OfflinePush:      true,
							StorageForServer: true,
							StorageForClient: true,
							UnreadCount:      false,
							NeedDecrypt:      false,
							UpdateConv:       true,
						},
						&msgmodel.MsgOfflinePush{
							Title:   self.Nickname,
							Content: "我们已经是好友了，快来聊天吧",
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