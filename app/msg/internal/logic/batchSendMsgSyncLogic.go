package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/msg/msgmodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xmgo"
	"github.com/cherish-chat/xxim-server/common/xredis/rediskey"
	"github.com/cherish-chat/xxim-server/common/xtrace"

	"github.com/cherish-chat/xxim-server/app/msg/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchSendMsgSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchSendMsgSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchSendMsgSyncLogic {
	return &BatchSendMsgSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchSendMsgSyncLogic) BatchSendMsgSync(in *pb.BatchSendMsgReq) (*pb.CommonResp, error) {
	msg := msgmodel.NewMsgFromPb(in.MsgData)
	msg.Receiver.UserId = ""
	msg.Receiver.GroupId = ""
	model := &msgmodel.BatchMsg{
		Id:          utils.GenId(),
		Msg:         msg.Check(),
		UserIdList:  utils.AnyMakeSlice(in.UserIdList),
		GroupIdList: utils.AnyMakeSlice(in.GroupIdList),
	}
	var err error
	xtrace.StartFuncSpan(l.ctx, "InsertOne", func(ctx context.Context) {
		_, err = l.svcCtx.Mongo().Collection(model).InsertOne(l.ctx, model)
	})
	if err != nil {
		l.Errorf("BatchSendMsgSync error: %v", err)
		return pb.NewRetryErrorResp(), err
	}
	var uidSeqMap = make(map[string]int64)
	var groupSeqMap = make(map[string]int64)
	xtrace.StartFuncSpan(l.ctx, "MHSet", func(ctx context.Context) {
		var kvs []xmgo.MHSetKv
		for _, userId := range in.UserIdList {
			convId := msgmodel.SingleConvId(in.MsgData.Sender, userId)
			// 给会话生成一个新的seq
			var seq int
			seq, err = IncrConvMaxSeq(l.svcCtx.Redis(), l.ctx, convId)
			if err != nil {
				return
			}
			msgId := msgmodel.ServerMsgId(convId, int64(seq))
			kvs = append(kvs, xmgo.MHSetKv{
				Key: rediskey.ConvMsgIdMapping(convId),
				HK:  msgId,
				V:   model.Id,
			})
			uidSeqMap[userId] = int64(seq)
		}
		for _, groupId := range in.GroupIdList {
			convId := groupId
			// 给会话生成一个新的seq
			var seq int
			seq, err = IncrConvMaxSeq(l.svcCtx.Redis(), l.ctx, convId)
			if err != nil {
				l.Errorf("redis Hincrby error: %v", err)
				return
			}
			msgId := msgmodel.ServerMsgId(convId, int64(seq))
			kvs = append(kvs, xmgo.MHSetKv{
				Key: rediskey.ConvMsgIdMapping(convId),
				HK:  msgId,
				V:   model.Id,
			})
			groupSeqMap[groupId] = int64(seq)
		}
		err = xmgo.MHSet(l.svcCtx.Mongo().Collection(&xmgo.MHSetKv{}), l.ctx, kvs...)
	})
	if err != nil {
		l.Errorf("redis MHSet error: %v", err)
		return pb.NewRetryErrorResp(), err
	}
	// 推送给相关的在线用户
	xtrace.StartFuncSpan(l.ctx, "PushMsgList", func(ctx context.Context) {
		msgDateList := make([]*pb.MsgData, 0, len(in.UserIdList)+len(in.GroupIdList))
		for _, userId := range in.UserIdList {
			m := msg.SinglePb(uidSeqMap[userId], userId)
			msgDateList = append(msgDateList, m)
		}
		for _, groupId := range in.GroupIdList {
			m := msg.GroupPb(groupSeqMap[groupId], groupId)
			msgDateList = append(msgDateList, m)
		}
		_, err = NewPushMsgListLogic(ctx, l.svcCtx).PushMsgList(&pb.PushMsgListReq{MsgDataList: msgDateList})
		if err != nil {
			l.Errorf("PushMsgList error: %v", err)
		}
	})
	return &pb.CommonResp{}, nil
}