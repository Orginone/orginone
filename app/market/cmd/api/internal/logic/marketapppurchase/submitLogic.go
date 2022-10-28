package marketapppurchase

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitLogic {
	return SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req types.MarketAppPurchasSubmitReq) (resp *types.CommonResponse, err error) {
	if !tools.IsNilOrEmpty(req.GroupId) {
		groupId, err := strconv.ParseInt(req.GroupId, 10, 64)
		if err != nil {
			return types.Failed(http.StatusBadRequest, errors.New("请求参数错误"))
		}
		return types.CommonResult(l.svcCtx.MarketRpc.GroupAppPurchas(l.ctx, &market.GroupAppPurchasReq{
			GroupIds: []int64{groupId},
			AppId:    req.AppId,
		}))
	} else if !tools.IsNilOrEmpty(req.TenantCode) {
		return types.CommonResult(l.svcCtx.MarketRpc.UnitAppPurchas(l.ctx, &market.UnitAppPurchasReq{
			TenantCode: req.TenantCode,
			AppId:      req.AppId,
		}))
	}
	return types.Failed(http.StatusBadRequest, errors.New("请求参数错误"))
}
