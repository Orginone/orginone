package api

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/app/market/model"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) RolelistLogic {
	return RolelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolelistLogic) Rolelist(req types.Nil) (resp *types.CommonResponse, err error) {
	appId, _, err := common.GetAppTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.MarketRpc.ApiRoleList(l.ctx, &market.ApiRoleLitReq{
		AppId: appId,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	appRoles := make([]*model.AsMarketAppRoleModel, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &appRoles)
	return types.JsonResult(appRoles, err)
}
