package marketusedapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitLogic {
	return SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req types.SubmitReq15) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.SubmitMarketUsedApp(l.ctx, &market.SubmitMarketUsedAppReq{
		Id:     -1,
		UserId: req.UserId,
		AppId:  req.AppId,
		Sort:   -1,
	}))
}
