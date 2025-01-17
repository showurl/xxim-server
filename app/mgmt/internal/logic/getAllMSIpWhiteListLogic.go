package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/mgmt/mgmtmodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"

	"github.com/cherish-chat/xxim-server/app/mgmt/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllMSIpWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllMSIpWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllMSIpWhiteListLogic {
	return &GetAllMSIpWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllMSIpWhiteListLogic) GetAllMSIpWhiteList(in *pb.GetAllMSIpWhiteListReq) (*pb.GetAllMSIpWhiteListResp, error) {
	var models []*mgmtmodel.MSIPWhitelist
	wheres := xorm.NewGormWhere()
	if in.Filter != nil {
		for k, v := range in.Filter {
			if v == "" {
				continue
			}
			switch k {
			case "id":
				wheres = append(wheres, xorm.Where("id = ?", v))
			case "isEnable":
				if v == "true" || v == "1" {
					wheres = append(wheres, xorm.Where("isEnable = ?", true))
				} else {
					wheres = append(wheres, xorm.Where("isEnable = ?", false))
				}
			case "remark":
				wheres = append(wheres, xorm.Where("remark like ?", "%"+v+"%"))
			case "ip":
				wheres = append(wheres, xorm.Where("startIp <= ? AND endIp >= ?", v, v))
			case "time_gte":
				val := utils.AnyToInt64(v)
				wheres = append(wheres, xorm.Where("createTime >= ?", val))
			case "time_lte":
				val := utils.AnyToInt64(v)
				wheres = append(wheres, xorm.Where("createTime <= ?", val))
			}
		}
	}
	count, err := xorm.ListWithPagingOrder(l.svcCtx.Mysql(), &models, &mgmtmodel.MSIPWhitelist{}, in.Page.Page, in.Page.Size, "createTime DESC", wheres...)
	if err != nil {
		l.Errorf("ListWithPagingOrder err: %v", err)
		return &pb.GetAllMSIpWhiteListResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	var resp []*pb.MSIpWhiteList
	for _, model := range models {
		resp = append(resp, model.ToPB())
	}
	return &pb.GetAllMSIpWhiteListResp{
		IpWhiteLists: resp,
		Total:        count,
	}, nil
}
