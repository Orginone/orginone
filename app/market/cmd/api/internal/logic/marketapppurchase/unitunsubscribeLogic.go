package marketapppurchase

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnitunsubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnitunsubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UnitunsubscribeLogic {
	return UnitunsubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnitunsubscribeLogic) Unitunsubscribe(req types.UnitunsubscribeReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	_, err = l.svcCtx.MarketRpc.GroupAppUnsubscribe(l.ctx, &market.GroupAppUnsubscribeReq{
		GroupId:    req.GroupId,
		TenantCode: tenantCode,
		AppIds:     tools.ArrayStrToInt64(req.AppIds),
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful("退订成功")
}
