package marketapproledistributionnew

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppIdByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppIdByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAppIdByUserIdLogic {
	return GetAppIdByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppIdByUserIdLogic) GetAppIdByUserId(req types.GetAppIdByUserIdReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.MarketRpc.GetUserAppIds(l.ctx, &market.GetUserAppIdsReq{
		UserId: req.UserId,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	ids := make([]int, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &ids)
	return types.JsonResult(tools.NumArrayToStrArray(ids), err)
}
