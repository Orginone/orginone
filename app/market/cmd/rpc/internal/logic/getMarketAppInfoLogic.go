package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappkeysecret"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/asmarketrolemenu"
	"orginone/common/schema/asredeploydata"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMarketAppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMarketAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMarketAppInfoLogic {
	return &GetMarketAppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用详情
func (l *GetMarketAppInfoLogic) GetMarketAppInfo(in *market.PrimaryKeyReq) (*market.CommonRpcRes, error) {
	appModel := new(model.AppModelDetail)
	appModel.Menu = make([]*model.AppMarketMenuModel, 0)
	appModel.RoleList = make([]*model.AsMarketAppRoleModel, 0)
	appModel.RoleNameList = make([]string, 0)
	appEntity, err := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.ID(in.Id)).WithAppRedeploys(func(ardq *schema.AsRedeployDataQuery) {
		ardq.Where(asredeploydata.Status(0), asredeploydata.IsDeleted(0)).Limit(1)
	}).WithAppKeys(func(amaksq *schema.AsMarketAppKeySecretQuery) {
		amaksq.Where(asmarketappkeysecret.IsDeleted(0)).Limit(1)
	}).WithAppRoles(func(amarq *schema.AsMarketAppRoleQuery) {
		amarq.Where(asmarketapprole.IsDeleted(0))
	}).WithAppMenus(func(ammq *schema.AsMarketMenuQuery) {
		ammq.Where(asmarketmenu.IsDeleted(0), asmarketmenu.HasRoleMenusWith(asmarketrolemenu.IsDeleted(0))).WithRoleMenus()
	}).First(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if len(appEntity.Edges.AppRedeploys) > 0 {
		appModel.App.RedeployId = appEntity.Edges.AppRedeploys[0].ID
		appModel.App.RedeployStatus = appEntity.Edges.AppRedeploys[0].Status
	}
	if len(appEntity.Edges.AppKeys) > 0 {
		appModel.App.AppKey = appEntity.Edges.AppKeys[0].AppKey
		appModel.App.AppSecret = appEntity.Edges.AppKeys[0].AppSecret
	}
	if len(appEntity.Edges.AppMenus) > 0 && len(appEntity.Edges.AppRoles) > 0 {
		linq.From(appEntity.Edges.AppMenus).SelectT(func(m *schema.AsMarketMenu) *model.AppMarketMenuModel {
			appMenuModel := new(model.AppMarketMenuModel)
			appMenuModel.RoleList = make([]*schema.AsMarketAppRole, 0)
			if len(m.Edges.RoleMenus) > 0 {
				for _, m := range m.Edges.RoleMenus {
					for _, r := range appEntity.Edges.AppRoles {
						if m.RoleID == r.ID {
							appMenuModel.RoleList = append(appMenuModel.RoleList, r)
							appMenuModel.RoleNameList = append(appMenuModel.RoleNameList, r.RoleName)
						}
					}
				}
			}
			m.Edges = schema.AsMarketMenuEdges{}
			appMenuModel.AsMarketMenu = m
			appMenuModel.Children = make([]*model.AppMarketMenuModel, 0)
			return appMenuModel
		}).ToSlice(&appModel.Menu)
	}
	if len(appEntity.Edges.AppRoles) > 0 {
		linq.From(appEntity.Edges.AppRoles).SelectT(func(r *schema.AsMarketAppRole) *model.AsMarketAppRoleModel {
			model := new(model.AsMarketAppRoleModel)
			roleDistributions, err := l.svcCtx.MarketStore.AsMarketRoleDistribution.Query().Where(asmarketroledistribution.RoleID(r.ID),
				asmarketroledistribution.TenantCode(in.TenantCode), asmarketroledistribution.IsDeleted(0)).All(l.ctx)
			if err == nil && len(roleDistributions) > 0 {
				for _, d := range roleDistributions {
					if d.UserID > 0 {
						model.DistributionPerson++
					}
					if d.AgencyID > 0 {
						model.DistributionDept++
					}
					if d.JobID > 0 {
						model.DistributionJob++
					}
				}
			}
			model.AsMarketAppRole = r
			return model
		}).ToSlice(&appModel.RoleList)
		appModel.RoleNameList = make([]string, 0)
		for _, v := range appModel.RoleList {
			appModel.RoleNameList = append(appModel.RoleNameList, v.RoleName)
		}
	}
	appEntity.Edges = schema.AsMarketAppEdges{} //清空边数据
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0),
		astenant.TenantCode(appEntity.TenantID), astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().Limit(1).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if len(tenantEntitys) > 0 {
		appModel.App.TenantName = tenantEntitys[0].TenantName
		appModel.App.City = tenantEntitys[0].Edges.Unit.City
		appModel.App.UnitName = tenantEntitys[0].Edges.Unit.UnitName
		appModel.App.Province = tenantEntitys[0].Edges.Unit.Province
		appModel.App.UnitAddress = tenantEntitys[0].Edges.Unit.StreetAddress
	}
	appModel.App.AsMarketApp = appEntity
	return market.JsonResult(appModel, err)
}
