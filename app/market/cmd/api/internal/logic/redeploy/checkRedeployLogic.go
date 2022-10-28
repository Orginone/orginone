package redeploy

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckRedeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckRedeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckRedeployLogic {
	return CheckRedeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckRedeployLogic) CheckRedeploy(req types.CheckRedeployReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(l.svcCtx.MarketRpc.CheckRedeploy(l.ctx, &market.CheckRedeployReq{
		RedeployIds: tools.ArrayStrToInt64(req.RedeployIds),
		Pass:        req.CheckStatus,
	}))
}
