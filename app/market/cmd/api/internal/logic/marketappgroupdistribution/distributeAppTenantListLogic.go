package marketappgroupdistribution

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type DistributeAppTenantListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDistributeAppTenantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DistributeAppTenantListLogic {
	return DistributeAppTenantListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DistributeAppTenantListLogic) DistributeAppTenantList(req types.DistributeAppTenantListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.GetGroupDistributeAppTenantList(l.ctx, &market.GroupPurchaseAppReq{
		AppId:    req.AppId,
		GroupId:  req.GroupId,
		UnitName: req.UnitName,
		Page: &market.PageRequest{
			Offset: (req.Current - 1) * req.Size,
			Limit:  req.Size,
		},
	}))
}
