package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FintTenantPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFintTenantPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FintTenantPurchaseLogic {
	return &FintTenantPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单位应用
func (l *FintTenantPurchaseLogic) FintTenantPurchase(in *market.TenantPurchaseReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.IsDeleted(0)).Where(asmarketapp.Or(
		asmarketapp.HasAppGroupDistribsWith(asmarketappgroupdistribution.IsDeleted(0), asmarketappgroupdistribution.TenantID(in.TenantCode)),
		asmarketapp.HasAppPurchasesWith(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.TenantID(in.TenantCode))))
	if !tools.IsNilOrEmpty(in.AppName) {
		query = query.Where(asmarketapp.AppNameContains(in.AppName))
	}
	total, err := query.Count(l.ctx)
	appEntitys, err := query.WithAppPurchases(func(amapq *schema.AsMarketAppPurchaseQuery) { amapq.Where(asmarketapppurchase.TenantID(in.TenantCode)) }).
		Order(schema.Asc(asmarketapp.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Offset)).All(l.ctx)
	type AppModel struct {
		*schema.AsMarketApp
		IfPurchase bool   `json:"ifPurchase"`
		TenantName string `json:"tenantName"`
	}
	appModels := make([]*AppModel, 0)
	linq.From(appEntitys).SelectT(func(a *schema.AsMarketApp) *AppModel {
		ifPurchase := false
		if a.Edges.AppPurchases != nil && len(a.Edges.AppPurchases) > 0 {
			ifPurchase = true
			a.CreateTime = a.Edges.AppPurchases[0].CreateTime
			a.UpdateTime = a.Edges.AppPurchases[0].UpdateTime
		}
		tenantName := l.svcCtx.MarketStore.GetTenantNameByTenantCode(l.ctx, a.TenantID, "未知")
		a.Edges = schema.AsMarketAppEdges{}
		return &AppModel{
			IfPurchase:  ifPurchase,
			AsMarketApp: a,
			TenantName:  tenantName,
		}
	}).ToSlice(&appModels)
	return market.PageResult(in.Page, total, appModels, err)
}
