package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteAppLogic {
	return DeleteAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAppLogic) DeleteApp(req types.DeleteAppReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppDelete(l.ctx, &market.DeleteAppReq{
		AppId:       req.AppId,
		RemoveAdmin: req.RemoveAdmin,
	}))
}
