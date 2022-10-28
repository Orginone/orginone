package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseAppTenantListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseAppTenantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PurchaseAppTenantListLogic {
	return PurchaseAppTenantListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseAppTenantListLogic) PurchaseAppTenantList(req types.PurchaseAppTenantListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.GetAppDisTenantList(l.ctx, &market.AppLinkReq{
		AppId: req.AppId,
		Page: &market.PageRequest{
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
