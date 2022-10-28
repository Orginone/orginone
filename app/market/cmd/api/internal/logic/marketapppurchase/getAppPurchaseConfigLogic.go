package marketapppurchase

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppPurchaseConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppPurchaseConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAppPurchaseConfigLogic {
	return GetAppPurchaseConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppPurchaseConfigLogic) GetAppPurchaseConfig(req types.GetAppPurchaseConfigReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.MarketRpc.GetAppPurchasConfig(l.ctx, &market.AppPurchasConfigReq{
		AppId:      req.AppId,
		TenantCode: tenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	config := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &config)
	return types.JsonResult(config, err)
}
