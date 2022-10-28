package marketapppurchase

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnsubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnsubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UnsubscribeLogic {
	return UnsubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnsubscribeLogic) Unsubscribe(req types.UnsubscribeReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	_, err = l.svcCtx.MarketRpc.UnitAppUnsubscribe(l.ctx, &market.UnitAppUnsubscribeReq{AppIds: req.AppIds, TenantCode: tenantCode})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful("退订成功")
}
