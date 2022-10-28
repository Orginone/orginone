package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMarketAppCheckListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMarketAppCheckListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMarketAppCheckListLogic {
	return &FindMarketAppCheckListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用审核列表分页
func (l *FindMarketAppCheckListLogic) FindMarketAppCheckList(in *market.FindMarketAppCheckListReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.AppNameContains(in.AppName), asmarketapp.IsDeleted(0))
	switch in.Type {
	case -1:
		query = query.Where(asmarketapp.StatusIn(0, 3, 6))
		break
	case 9:
		query = query.Where(asmarketapp.StatusIn(4, 5))
		break
	case 10:
		break
	case 11:
		query = query.Where(asmarketapp.StatusNotIn(0, 3, 6))
		break
	default:
		query = query.Where(asmarketapp.Status(in.Type))
		break
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	appEntitys, err := query.Order(schema.Desc(asmarketapp.FieldApplyTime)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)

	tenantIds := make([]string, 0)
	for _, app := range appEntitys {
		tenantIds = append(tenantIds, app.TenantID)
	}
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.TenantCodeIn(tenantIds...)).WithUnit().All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	Records := make([]*model.AppModel, 0)
	for _, app := range appEntitys {
		record := &model.AppModel{
			AsMarketApp: app,
		}
		for _, v := range tenantEntitys {
			if app.TenantID == v.TenantCode {
				unit, err := v.Edges.UnitOrErr()
				if err != nil {
					break
				}
				record.TenantName = v.TenantName
				record.UnitAddress = unit.Province + unit.City + unit.StreetAddress
				break
			}
		}
		Records = append(Records, record)
	}
	return market.PageResult(in.Page, int64(total), Records, err)
}
