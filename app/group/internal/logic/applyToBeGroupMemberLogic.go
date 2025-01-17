package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/group/groupmodel"
	"github.com/cherish-chat/xxim-server/app/notice/noticemodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"
	"github.com/cherish-chat/xxim-server/common/xtrace"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"
	"time"

	"github.com/cherish-chat/xxim-server/app/group/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyToBeGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyToBeGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyToBeGroupMemberLogic {
	return &ApplyToBeGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ApplyToBeGroupMember 申请加入群聊
func (l *ApplyToBeGroupMemberLogic) ApplyToBeGroupMember(in *pb.ApplyToBeGroupMemberReq) (*pb.ApplyToBeGroupMemberResp, error) {
	// 获取群信息
	group := &groupmodel.Group{}
	{
		var mapGroupByIdResp *pb.MapGroupByIdsResp
		var err error
		xtrace.StartFuncSpan(l.ctx, "MapGroupByIds", func(ctx context.Context) {
			mapGroupByIdResp, err = NewMapGroupByIdsLogic(ctx, l.svcCtx).MapGroupByIds(&pb.MapGroupByIdsReq{Ids: []string{in.GroupId}})
		})
		if err != nil {
			l.Errorf("MapGroupByIds error: %v", err)
			return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		if len(mapGroupByIdResp.GroupMap) == 0 {
			l.Errorf("MapGroupByIds error: %v", err)
			return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		group = groupmodel.GroupFromBytes(mapGroupByIdResp.GroupMap[in.GroupId])
	}
	// 判断群人数是否达到上限
	if int32(group.MemberCount) >= group.MaxMember {
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewAlertErrorResp("操作失败", "群人数已达上限")}, nil
	}
	// 查询自己加了多少群
	var groupCount int64
	{
		var err error
		xtrace.StartFuncSpan(l.ctx, "GetMyGroupList", func(ctx context.Context) {
			list, e := NewGetMyGroupListLogic(ctx, l.svcCtx).GetMyGroupList(&pb.GetMyGroupListReq{
				CommonReq: in.GetCommonReq(),
				Opt:       pb.GetMyGroupListReq_ONLY_ID,
			})
			if e != nil {
				l.Errorf("GetMyGroupList error: %v", e)
				err = e
			}
			groupCount = int64(len(list.GroupMap))
		})
		if err != nil {
			return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
	}
	if int64(l.svcCtx.Config.GroupConfig.MaxGroupCount) <= groupCount {
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewAlertErrorResp("操作失败", "您已加入了太多的群聊")}, nil
	}
	apply := &groupmodel.GroupApply{
		Id:         utils.GenId(),
		GroupId:    in.GroupId,
		UserId:     in.CommonReq.UserId,
		Result:     pb.GroupApplyHandleResult_UNHANDLED,
		Reason:     in.Reason,
		ApplyTime:  time.Now().UnixMilli(),
		HandleTime: 0,
	}
	// 判断是否申请过了
	var count int64
	err := l.svcCtx.Mysql().Model(&groupmodel.GroupApply{}).Where(
		"groupId = ? and userId = ? and result = ?",
		in.GroupId, in.CommonReq.UserId, pb.GroupApplyHandleResult_UNHANDLED,
	).Count(&count).Error
	if err != nil {
		l.Errorf("ApplyToBeGroupMember Count error: %v", err)
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	if count > 0 {
		// 直接返回成功
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewSuccessResp()}, nil
	}
	// 获取群里所有的管理员
	groupManagers, err := getAllGroupManager(l.ctx, l.svcCtx, in.GroupId, true)
	if err != nil {
		l.Errorf("ApplyToBeGroupMember getAllGroupManager error: %v", err)
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	err = xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
		err := xorm.InsertOne(tx, apply)
		if err != nil {
			l.Errorf("ApplyToBeGroupMember InsertOne error: %v", err)
			return err
		}
		return nil
	}, func(tx *gorm.DB) error {
		for _, manager := range groupManagers {
			notice := &noticemodel.Notice{
				ConvId: pb.HiddenConvIdGroupMember(),
				UserId: manager.UserId,
				Options: noticemodel.NoticeOption{
					StorageForClient: false,
					UpdateConvNotice: false,
				},
				ContentType: pb.NoticeContentType_ApplyToBeGroupMember,
				Content: utils.AnyToBytes(pb.NoticeContent_ApplyToBeGroupMember{
					ApplyId:      apply.Id,
					GroupId:      apply.GroupId,
					UserId:       apply.UserId,
					Result:       apply.Result,
					Reason:       apply.Reason,
					ApplyTime:    apply.ApplyTime,
					HandleTime:   apply.HandleTime,
					HandleUserId: apply.HandleUserId,
				}),
				UniqueId: apply.Id,
				Title:    "",
				Ext:      nil,
			}
			err := notice.Insert(l.ctx, tx, l.svcCtx.Redis())
			if err != nil {
				l.Errorf("insert notice failed, err: %v", err)
				return err
			}
			return nil
		}
		return nil
	})
	if err != nil {
		l.Errorf("ApplyToBeGroupMember Transaction error: %v", err)
		return &pb.ApplyToBeGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	// 通知给群里所有的管理员
	go xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "SendNotice", func(ctx context.Context) {
		utils.RetryProxy(ctx, 12, time.Second, func() error {
			for _, manager := range groupManagers {
				_, err := l.svcCtx.NoticeService().GetUserNoticeData(ctx, &pb.GetUserNoticeDataReq{
					CommonReq: in.CommonReq,
					UserId:    manager.UserId,
					ConvId:    pb.HiddenConvIdGroupMember(),
				})
				if err != nil {
					l.Errorf("ApplyToBeGroupMember SendNoticeData error: %v", err)
					return err
				}
			}
			return nil
		})
	}, propagation.MapCarrier{
		"groupId":  in.GroupId,
		"userId":   in.CommonReq.UserId,
		"noticeId": apply.Id,
	})
	return &pb.ApplyToBeGroupMemberResp{}, nil
}
