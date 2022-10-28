package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupDistributeAppTenantListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupDistributeAppTenantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupDistributeAppTenantListLogic {
	return &GetGroupDistributeAppTenantListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团租户应用购买列表
func (l *GetGroupDistributeAppTenantListLogic) GetGroupDistributeAppTenantList(in *market.GroupPurchaseAppReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppGroupDistribution.Query().Where(asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.GroupID(in.GroupId), asmarketappgroupdistribution.AppID(in.AppId),
		asmarketappgroupdistribution.HasAppxWith(asmarketapp.IsDeleted(0)))
	tenantCodes, err := query.Select(asmarketappgroupdistribution.FieldTenantID).Strings(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	queryTenant := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(tenantCodes...), astenant.HasUnitWith(asunit.IsDeleted(0)))
	if !tools.IsNilOrEmpty(in.UnitName) {
		queryTenant = queryTenant.Where(astenant.HasUnitWith(asunit.UnitNameContains(in.UnitName)))
	}
	total, err := queryTenant.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	tenantEntitys, err := queryTenant.WithUnit().Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	type UnitModel struct {
		*schema.AsUnit
		TenantCode string `json:"tenantCode"`
		TenantType int64  `json:"tenantType"`
	}
	unitModels := make([]*UnitModel, 0)
	linq.From(tenantEntitys).SelectT(func(t *schema.AsTenant) *UnitModel {
		model := new(UnitModel)
		model.AsUnit = t.Edges.Unit
		model.TenantCode = t.TenantCode
		model.TenantType = t.TenantType
		return model
	}).ToSlice(&unitModels)
	return market.PageResult(in.Page, total, unitModels, err)
}
