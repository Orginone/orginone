package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppPurchasConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppPurchasConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppPurchasConfigLogic {
	return &GetAppPurchasConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用对当前租户的分发情况
func (l *GetAppPurchasConfigLogic) GetAppPurchasConfig(in *market.AppPurchasConfigReq) (*market.CommonRpcRes, error) {
	config := make(map[string]interface{}, 0)
	config["group"] = make([]string, 0)
	config["unit"] = market.Nil{}
	query := l.svcCtx.MarketStore.AsMarketAppPurchase.Query().Where(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.AppID(in.AppId))
	exist, err := query.Clone().Where(asmarketapppurchase.TenantID(in.TenantCode)).Exist(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if exist {
		tenantEntity, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCode(in.TenantCode),
			astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().First(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
		Unit := new(model.UnitModel)
		Unit.AsUnit = tenantEntity.Edges.Unit
		Unit.TenantCode = tenantEntity.TenantCode
		Unit.TenantType = tenantEntity.TenantType
		if tenantEntity.Edges.Unit.IsVirtual == "1" {
			Unit.IsVirtual = 1
			Unit.VirtualUnit = true
		} else {
			Unit.IsVirtual = 0
			Unit.VirtualUnit = false
		}
		config["unit"] = Unit
	}
	groupIds, err := query.QueryGroupx().Where(asallgroup.IsDeleted(0), asallgroup.TenantCode(in.TenantCode)).Select(asallgroup.FieldID).Strings(l.ctx)
	if len(groupIds) > 0 {
		config["group"] = groupIds
	}
	return market.JsonResult(config, err)
}
