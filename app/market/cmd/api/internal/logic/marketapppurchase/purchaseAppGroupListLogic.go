package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseAppGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseAppGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PurchaseAppGroupListLogic {
	return PurchaseAppGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseAppGroupListLogic) PurchaseAppGroupList(req types.PurchaseAppGroupListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.GetAppDisGroupList(l.ctx, &market.AppLinkReq{
		AppId: req.AppId,
		Page: &market.PageRequest{
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
