package marketusedapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveByRelationLogic {
	return RemoveByRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByRelationLogic) RemoveByRelation(req types.RemoveByRelationReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.RemoveByRelation(l.ctx, &market.RemoveByRelationReq{
		AppId:  req.AppId,
		UserId: req.UserId,
	}))
}
