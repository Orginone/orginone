package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListcheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListcheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListcheckLogic {
	return ListcheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListcheckLogic) Listcheck(req types.ListCheckReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.FindMarketAppCheckList(l.ctx, &market.FindMarketAppCheckListReq{
		Page: &market.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		Type:    req.Count,
		AppName: req.AppName,
	}))
}
