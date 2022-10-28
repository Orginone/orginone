package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppDisGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppDisGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppDisGroupListLogic {
	return &GetAppDisGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用已购集团
func (l *GetAppDisGroupListLogic) GetAppDisGroupList(in *market.AppLinkReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0),
		asallgroup.HasAppPurchasesWith(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.AppID(in.AppId)))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	groupEntitys, err := query.WithAppPurchases(func(amapq *schema.AsMarketAppPurchaseQuery) {
		amapq.Where(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.AppID(in.AppId))
	}).Order(schema.Asc(asallgroup.FieldTenantCode)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit) * 2).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	tenantCodes := make([]string, 0)
	linq.From(groupEntitys).SelectT(func(g *schema.AsAllGroup) string { return g.TenantCode }).ToSlice(&tenantCodes)
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(tenantCodes...),
		astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().Order(schema.Asc(astenant.FieldID)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	groupModels := make([]*model.GroupModel, 0)
	for _, t := range tenantEntitys {
		groupModel := new(model.GroupModel)
		for _, g := range groupEntitys {
			if g.TenantCode == t.TenantCode {
				groupModel.AsAllGroup = g
				groupModel.UseStatus = g.Edges.AppPurchases[0].UseStatus
				groupModel.UnitName = t.Edges.Unit.UnitName
				groupModel.LinkMan = t.Edges.Unit.LinkMan
				groupModel.LinkPhone = t.Edges.Unit.LinkPhone
				groupModels = append(groupModels, groupModel)
				break
			}
		}
	}
	return market.PageResult(in.Page, total, groupModels, err)
}
