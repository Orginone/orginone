package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnitAppPurchasLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnitAppPurchasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnitAppPurchasLogic {
	return &UnitAppPurchasLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 单位获取应用
func (l *UnitAppPurchasLogic) UnitAppPurchas(in *market.UnitAppPurchasReq) (*market.CommonRpcRes, error) {
	_, err := l.svcCtx.MarketStore.AsMarketAppPurchase.Create().SetAppID(in.AppId).SetTenantID(in.TenantCode).SetUseStatus(1).Save(l.ctx)
	return &market.CommonRpcRes{Json: "true"}, err
}
