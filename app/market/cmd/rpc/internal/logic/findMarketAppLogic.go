package logic

import (
	"context"
	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMarketAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMarketAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMarketAppLogic {
	return &FindMarketAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用市场应用分页
func (l *FindMarketAppLogic) FindMarketApp(in *market.MarketAppReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.AppName) {
		query = query.Where(asmarketapp.AppNameContains(in.AppName))
	}
	if len(in.SaleStatus) > 0 {
		query = query.Where(asmarketapp.SaleStatusIn(in.SaleStatus...))
	}
	if len(in.Status) > 0 {
		query = query.Where(asmarketapp.StatusIn(in.Status...))
	}
	if len(in.TargetUser) > 0 {
		query = query.Where(asmarketapp.TargetUserIn(in.TargetUser...))
	}
	if !tools.IsNilOrEmpty(in.TenantId) {
		query = query.Where(asmarketapp.TenantID(in.TenantId))
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	appEntitys, err := query.Order(schema.Asc(asmarketapp.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	tenantCodes, err := query.Select(asmarketapp.FieldTenantID).Strings(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	var appIds []int64
	for _, item := range appEntitys {
		appIds = append(appIds, item.ID)
	}
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(tenantCodes...), astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	subQuery := l.svcCtx.MarketStore.AsMarketAppPurchase.Query().Where(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.AppIDIn(appIds...))
	if in.GroupId == -1 {
		subQuery = subQuery.Where(asmarketapppurchase.TenantID(in.TenantCode))
	} else {
		subQuery = subQuery.Where(asmarketapppurchase.GroupID(in.GroupId))
	}
	purchaseEntitys, err := subQuery.All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	type AppModel struct {
		*schema.AsMarketApp
		TenantName  string `json:"tenantName"`
		UnitAddress string `json:"unitAddress"`
		IfPurchase  bool   `json:"ifPurchase"`
	}
	appModels := make([]*AppModel, 0)
	linq.From(appEntitys).SelectT(func(a *schema.AsMarketApp) *AppModel {
		appModel := new(AppModel)
		for _, t := range tenantEntitys {
			if t.TenantCode == a.TenantID && t.Edges.Unit != nil {
				appModel.TenantName = t.Edges.Unit.UnitName
				appModel.UnitAddress = t.Edges.Unit.StreetAddress
			}
		}
		appModel.IfPurchase = false
		for _, t := range purchaseEntitys {
			if t.AppID == a.ID {
				appModel.IfPurchase = true
			}
		}
		a.Edges = schema.AsMarketAppEdges{}
		appModel.AsMarketApp = a
		return appModel
	}).ToSlice(&appModels)
	return market.PageResult(in.Page, total, appModels, err)
}
