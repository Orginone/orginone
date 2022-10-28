package redeploy

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) RedeployLogic {
	return RedeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedeployLogic) Redeploy(req []byte) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.MarketAppPublish(l.ctx, &market.CommonRpcReq{Json: string(req)}))
}
