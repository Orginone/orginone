package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOffSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetOffSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetOffSaleLogic {
	return SetOffSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetOffSaleLogic) SetOffSale(req types.SetOffSaleReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppOnOrOffSale(l.ctx, &market.AppOnOrOffSaleReq{
		Flag:   0,
		AppIds: req.IdList,
	}))
}
