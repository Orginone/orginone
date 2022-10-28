package marketapproledistributionnew

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppDistributedRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppDistributedRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAppDistributedRoleLogic {
	return GetAppDistributedRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppDistributedRoleLogic) GetAppDistributedRole(req types.GetAppDistributedRoleReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.MarketRpc.GetAppRoleDistribe(l.ctx, &market.GetAppRoleDistribeReq{
		AppId:      req.AppId,
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	data := make([]*model.AppRoleDistribeRes, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &data)
	return types.JsonResult(data, err)
}
