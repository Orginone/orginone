package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitSecondLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitSecondLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitSecondLogic {
	return SubmitSecondLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitSecondLogic) SubmitSecond(req types.SubmitSecondReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppDeploy(l.ctx, &market.AppDeployReq{AppId: req.AppId}))
}
