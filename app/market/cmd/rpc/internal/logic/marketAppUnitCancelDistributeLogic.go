package logic

import (
	"context"
	"orginone/common/schema/asmarketappgroupdistribution"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppUnitCancelDistributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppUnitCancelDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppUnitCancelDistributeLogic {
	return &MarketAppUnitCancelDistributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团取消分发
func (l *MarketAppUnitCancelDistributeLogic) MarketAppUnitCancelDistribute(in *market.MarketAppUnitCancelDistributeReq) (*market.CommonRpcRes, error) {
	return market.JsonResult(l.svcCtx.MarketStore.AsMarketAppGroupDistribution.Update().Where(asmarketappgroupdistribution.GroupID(in.GroupId), asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.AppID(in.AppId), asmarketappgroupdistribution.TenantID(in.TenantId), asmarketappgroupdistribution.IsDeleted(0)).SetIsDeleted(1).Save(l.ctx))
}
