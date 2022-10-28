package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketroledistribution"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnitAppUnsubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnitAppUnsubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnitAppUnsubscribeLogic {
	return &UnitAppUnsubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 单位应用退订
func (l *UnitAppUnsubscribeLogic) UnitAppUnsubscribe(in *market.UnitAppUnsubscribeReq) (*market.CommonRpcRes, error) {
	_, err := l.svcCtx.MarketStore.AsMarketAppPurchase.Update().Where(asmarketapppurchase.IsDeleted(0),
		asmarketapppurchase.AppIDIn(in.AppIds...), asmarketapppurchase.TenantID(in.TenantCode)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = l.svcCtx.MarketStore.AsMarketRoleDistribution.Update().Where(asmarketroledistribution.IsDeleted(0),
		asmarketroledistribution.TenantCode(in.TenantCode), asmarketroledistribution.HasRolexWith(
			asmarketapprole.IsDeleted(0), asmarketapprole.AppIDIn(in.AppIds...))).SetIsDeleted(1).Save(l.ctx)
	return market.Result("true", err)
}
