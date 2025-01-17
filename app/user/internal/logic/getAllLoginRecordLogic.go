package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/user/usermodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"

	"github.com/cherish-chat/xxim-server/app/user/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLoginRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLoginRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLoginRecordLogic {
	return &GetAllLoginRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLoginRecordLogic) GetAllLoginRecord(in *pb.GetAllLoginRecordReq) (*pb.GetAllLoginRecordResp, error) {
	if in.Page == nil {
		in.Page = &pb.Page{Page: 1, Size: 0}
	}
	var models []*usermodel.LoginRecord
	wheres := xorm.NewGormWhere()
	if in.Filter != nil {
		for k, v := range in.Filter {
			if v == "" {
				continue
			}
			switch k {
			case "id":
				wheres = append(wheres, xorm.Where("id = ?", v))
			case "userId":
				wheres = append(wheres, xorm.Where("userId = ?", v))
			case "ip":
				wheres = append(wheres, xorm.Where("ip LIKE ?", v+"%"))
			case "time_gte":
				val := utils.AnyToInt64(v)
				wheres = append(wheres, xorm.Where("time >= ?", val))
			case "time_lte":
				val := utils.AnyToInt64(v)
				wheres = append(wheres, xorm.Where("time <= ?", val))
			}
		}
	}
	count, err := xorm.ListWithPagingOrder(l.svcCtx.Mysql(), &models, &usermodel.LoginRecord{}, in.Page.Page, in.Page.Size, "time DESC", wheres...)
	if err != nil {
		l.Errorf("GetList err: %v", err)
		return &pb.GetAllLoginRecordResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	var resp []*pb.LoginRecord
	for _, model := range models {
		role := model.ToProto()
		resp = append(resp, role)
	}
	return &pb.GetAllLoginRecordResp{
		LoginRecordList: resp,
		Total:           count,
	}, nil
}
