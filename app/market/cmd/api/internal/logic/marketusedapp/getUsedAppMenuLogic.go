package marketusedapp

import (
	"context"
	"encoding/json"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsedAppMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsedAppMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUsedAppMenuLogic {
	return GetUsedAppMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsedAppMenuLogic) GetUsedAppMenu(req types.GetUsedAppMenuReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.MarketRpc.GetUsedAppMenu(l.ctx, &market.GetUsedAppMenuReq{
		UserId: req.UserId,
	})
	menuTrees := make([]*schema.AsMarketMenu, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &menuTrees)
	for _, menu := range menuTrees {
		menu.ID = menu.AppID
	}
	return types.JsonResult(menuTrees, err)
}
