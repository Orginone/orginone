package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByPurchaseTenantIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByPurchaseTenantIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListByPurchaseTenantIdLogic {
	return ListByPurchaseTenantIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByPurchaseTenantIdLogic) ListByPurchaseTenantId(req types.ListByPurchaseTenantIdReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.FintTenantPurchase(l.ctx, &market.TenantPurchaseReq{
		AppName:    req.AppName,
		TenantCode: req.TenantId,
		Page: &market.PageRequest{
			Offset:      int64(req.Size * (req.Current - 1)),
			Limit:       req.Size,
			SearchCount: true,
			Filter:      req.AppName,
		},
	}))
}
