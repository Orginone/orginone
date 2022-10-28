package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppOnOrOffSaleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppOnOrOffSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppOnOrOffSaleLogic {
	return &MarketAppOnOrOffSaleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用上下架
func (l *MarketAppOnOrOffSaleLogic) MarketAppOnOrOffSale(in *market.AppOnOrOffSaleReq) (*market.CommonRpcRes, error) {
	_, err := l.svcCtx.MarketStore.AsMarketApp.Update().Where(asmarketapp.IDIn(in.AppIds...), asmarketapp.IsDeleted(0)).SetSaleStatus(in.Flag).Save(l.ctx)
	return market.Result("success", err)
}
