package marketapp

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/app/market/model"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllLogic {
	return GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogic) GetAll(req types.GetAllReq) (resp *types.CommonResponse, err error) {
	// 查询人员相关
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	repres, err := l.svcCtx.MarketRpc.GetMarketAppInfo(l.ctx, &market.PrimaryKeyReq{Id: req.Id, TenantCode: tenantCode})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	appModel := new(model.AppModelDetail)
	err = json.Unmarshal([]byte(repres.Json), &appModel)
	return types.JsonResult(appModel, err)
}
