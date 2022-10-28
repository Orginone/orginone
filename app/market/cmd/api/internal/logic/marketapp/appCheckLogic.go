package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppCheckLogic {
	return AppCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppCheckLogic) AppCheck(req types.AppCheckReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.AppCheck(l.ctx, &market.AppCheckReq{
		Status: req.CheckStatus,
		AppIds: tools.ArrayStrToInt64(req.IdList),
	}))
}
