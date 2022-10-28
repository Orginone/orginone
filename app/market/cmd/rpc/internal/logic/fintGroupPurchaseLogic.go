package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FintGroupPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFintGroupPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FintGroupPurchaseLogic {
	return &FintGroupPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取集团应用
func (l *FintGroupPurchaseLogic) FintGroupPurchase(in *market.GroupPurchaseReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppPurchase.Query().Where(asmarketapppurchase.IsDeleted(0),
		asmarketapppurchase.GroupID(in.GroupId), asmarketapppurchase.UseStatus(1)).QueryAppx().Where(asmarketapp.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.AppName) {
		query = query.Where(asmarketapp.AppNameContains(in.AppName))
	}
	total, err := query.Count(l.ctx)
	appEntitys, err := query.Order(schema.Asc(asmarketapp.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Offset)).All(l.ctx)
	type AppModel struct {
		*schema.AsMarketApp
		Children   []*schema.AsMarketApp `json:"children"`
		TenantName string                `json:"tenantName"`
	}
	appModels := make([]*AppModel, 0)
	linq.From(appEntitys).SelectT(func(a *schema.AsMarketApp) *AppModel {
		tenantName := l.svcCtx.MarketStore.GetTenantNameByTenantCode(l.ctx, a.TenantID, "未知")
		a.Edges = schema.AsMarketAppEdges{}
		return &AppModel{
			AsMarketApp: a,
			TenantName:  tenantName,
		}
	}).ToSlice(&appModels)
	return market.PageResult(in.Page, total, appModels, err)
}
