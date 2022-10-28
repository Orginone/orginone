package marketapp

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.AppListReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	saleStatus := make([]int64, 0)
	if req.SaleStatus > -1 {
		saleStatus = append(saleStatus, req.SaleStatus)
	}
	return types.PageResult(l.svcCtx.MarketRpc.FindMarketApp(l.ctx, &market.MarketAppReq{
		AppName:    req.AppName,
		SaleStatus: saleStatus,
		Status:     []int64{},
		TargetUser: []int64{},
		GroupId:    -1,
		TenantId:   req.TenantId,
		TenantCode: tenantCode,
		Page: &market.PageRequest{
			Offset: (req.Size * (req.Current - 1)),
			Limit:  req.Size,
		},
	}))
}
