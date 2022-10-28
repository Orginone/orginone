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

type ListUnitAndPersonalVOLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUnitAndPersonalVOLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListUnitAndPersonalVOLogic {
	return ListUnitAndPersonalVOLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUnitAndPersonalVOLogic) ListUnitAndPersonalVO(req types.ListUnitAndPersonalVOReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.PageResult(l.svcCtx.MarketRpc.FindMarketApp(l.ctx, &market.MarketAppReq{
		AppName:    req.Name,
		SaleStatus: []int64{1},
		Status:     []int64{7},
		TargetUser: []int64{1, 2},
		GroupId:    -1,
		TenantCode: tenantCode,
		Page: &market.PageRequest{
			Offset: (req.Size * (req.Current - 1)),
			Limit:  req.Size,
		},
	}))
}
