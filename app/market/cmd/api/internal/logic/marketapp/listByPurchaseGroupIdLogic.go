package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByPurchaseGroupIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByPurchaseGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListByPurchaseGroupIdLogic {
	return ListByPurchaseGroupIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByPurchaseGroupIdLogic) ListByPurchaseGroupId(req types.ListByPurchaseGroupIdReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.FintGroupPurchase(l.ctx, &market.GroupPurchaseReq{
		AppName: req.AppName,
		GroupId: req.GroupId,
		Page: &market.PageRequest{
			Offset:      int64(req.Size * (req.Current - 1)),
			Limit:       req.Size,
			SearchCount: true,
			Filter:      req.AppName,
		},
	}))
}
