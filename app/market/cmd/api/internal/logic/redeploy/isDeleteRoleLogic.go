package redeploy

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsDeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) IsDeleteRoleLogic {
	return IsDeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsDeleteRoleLogic) IsDeleteRole(req types.IsDeleteRoleReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(l.svcCtx.MarketRpc.CheckRoleCanDelete(l.ctx, &market.CheckRoleReq{
		AppId:    req.AppId,
		RoleName: req.RoleName,
	}))
}
