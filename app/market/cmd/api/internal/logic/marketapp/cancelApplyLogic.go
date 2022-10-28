package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelApplyLogic {
	return CancelApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelApplyLogic) CancelApply(req types.CancelApplyReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppCancelApply(l.ctx, &market.AppCancelApplyReq{AppIds: req.IdList}))
}
