package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppDisTenantListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppDisTenantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppDisTenantListLogic {
	return &GetAppDisTenantListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用已购租户
func (l *GetAppDisTenantListLogic) GetAppDisTenantList(in *market.AppLinkReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppPurchase.Query().Where(asmarketapppurchase.IsDeleted(0),
		asmarketapppurchase.AppID(in.AppId), asmarketapppurchase.TenantIDNotNil())
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	codes, err := query.Order(schema.Asc(asmarketapppurchase.FieldTenantID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit) * 2).Select(asmarketapppurchase.FieldTenantID).Strings(l.ctx)
	if len(codes) == 0 {
		return market.PageResult(in.Page, 0, codes, nil)
	}
	tenantEntitys, err := l.svcCtx.MarketStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(codes...), astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().Order(schema.Asc(astenant.FieldID)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	unitModels := make([]*model.UnitModel, 0)
	for _, t := range tenantEntitys {
		unitModel := new(model.UnitModel)
		unitModel.AsUnit = t.Edges.Unit
		unitModel.TenantCode = t.TenantCode
		unitModel.TenantType = t.TenantType
		if t.Edges.Unit.IsVirtual == "1" {
			unitModel.IsVirtual = 1
			unitModel.VirtualUnit = true
		} else {
			unitModel.IsVirtual = 0
			unitModel.VirtualUnit = false
		}
		unitModels = append(unitModels, unitModel)
	}
	return market.PageResult(in.Page, total, unitModels, nil)
}
