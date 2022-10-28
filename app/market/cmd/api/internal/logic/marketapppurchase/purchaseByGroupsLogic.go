package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseByGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseByGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PurchaseByGroupsLogic {
	return PurchaseByGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseByGroupsLogic) PurchaseByGroups(req types.PurchaseByGroupsReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.GroupAppPurchas(l.ctx, &market.GroupAppPurchasReq{
		GroupIds: tools.ArrayStrToInt64(req.GroupIds),
		AppId:    req.AppId,
	}))
}
