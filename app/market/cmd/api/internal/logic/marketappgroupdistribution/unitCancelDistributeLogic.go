package marketappgroupdistribution

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnitCancelDistributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnitCancelDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UnitCancelDistributeLogic {
	return UnitCancelDistributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnitCancelDistributeLogic) UnitCancelDistribute(req types.UnitCancelDistributeReq) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.MarketRpc.MarketAppUnitCancelDistribute(l.ctx, &market.MarketAppUnitCancelDistributeReq{
		AppId:    req.AppId,
		GroupId:  req.GroupId,
		TenantId: req.TenantId,
	})
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.Successful(true)
}
