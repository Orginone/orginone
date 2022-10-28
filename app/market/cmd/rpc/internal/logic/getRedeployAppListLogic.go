package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asredeploydata"
	"orginone/common/schema/astenant"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeployAppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRedeployAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeployAppListLogic {
	return &GetRedeployAppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取重新发布应用列表
func (l *GetRedeployAppListLogic) GetRedeployAppList(in *market.GetRedeployReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.AppNameContains(in.AppName), asmarketapp.IsDeleted(0), asmarketapp.HasAppRedeploysWith(
		asredeploydata.IsDeleted(0), asredeploydata.StatusIn(in.Status...)))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	query = query.Order(schema.Asc(asmarketapp.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).WithAppRedeploys(func(ardq *schema.AsRedeployDataQuery) {
		ardq.Where(asredeploydata.IsDeleted(0), asredeploydata.StatusIn(in.Status...))
	})
	appTencodes, err := query.Clone().Select(asmarketapp.FieldTenantID).Strings(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	appEntitys, err := query.All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0),
		astenant.TenantCodeIn(appTencodes...), astenant.HasUnit()).WithUnit().All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	models := make([]*model.AppRedeployModel, 0)
	linq.From(appEntitys).SelectT(func(a *schema.AsMarketApp) *model.AppRedeployModel {
		m := new(model.AppRedeployModel)
		m.AppModel = new(model.AppModel)
		m.AsMarketApp = a
		if len(a.Edges.AppRedeploys) > 0 {
			m.RedeployId = a.Edges.AppRedeploys[0].ID
			m.RedeployStatus = a.Edges.AppRedeploys[0].Status
		}
		for _, t := range tenantEntitys {
			if t.TenantCode == a.TenantID && t.Edges.Unit != nil {
				m.TenantName = t.TenantName
				m.City = t.Edges.Unit.City
				m.UnitName = t.Edges.Unit.UnitName
				m.Province = t.Edges.Unit.Province
				m.UnitAddress = t.Edges.Unit.StreetAddress
			}
		}
		return m
	}).ToSlice(&models)
	return market.PageResult(in.Page, total, models, err)
}
