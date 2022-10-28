package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOnSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetOnSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetOnSaleLogic {
	return SetOnSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetOnSaleLogic) SetOnSale(req types.SetOnSaleReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppOnOrOffSale(l.ctx, &market.AppOnOrOffSaleReq{
		Flag:   1,
		AppIds: req.IdList,
	}))
}
