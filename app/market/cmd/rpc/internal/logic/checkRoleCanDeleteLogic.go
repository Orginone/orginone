package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapprole"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckRoleCanDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckRoleCanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckRoleCanDeleteLogic {
	return &CheckRoleCanDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询app角色是否可以删除(发布时)
func (l *CheckRoleCanDeleteLogic) CheckRoleCanDelete(in *market.CheckRoleReq) (*market.CommonRpcRes, error) {
	entity, err := l.svcCtx.MarketStore.AsMarketAppRole.Query().
		Where(
			asmarketapprole.AppID(in.AppId),
			asmarketapprole.RoleName(in.RoleName),
		).
		QueryRoleDistribs().
		All(l.ctx)
	return market.JsonResult(entity == nil, err)
}
