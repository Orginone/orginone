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

type ListByGroupIdVOLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByGroupIdVOLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListByGroupIdVOLogic {
	return ListByGroupIdVOLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByGroupIdVOLogic) ListByGroupIdVO(req types.ListByGroupIdVOReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.PageResult(l.svcCtx.MarketRpc.FindMarketApp(l.ctx, &market.MarketAppReq{
		AppName:    req.AppName,
		SaleStatus: []int64{1, 3},
		Status:     []int64{},
		TargetUser: []int64{},
		GroupId:    req.GroupId,
		TenantCode: tenantCode,
		Page: &market.PageRequest{
			Offset: (req.Size * (req.Current - 1)),
			Limit:  req.Size,
		},
	}))
}
