package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapprole"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApiRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiRoleListLogic {
	return &ApiRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据应用token获取应用角色信息
func (l *ApiRoleListLogic) ApiRoleList(in *market.ApiRoleLitReq) (*market.CommonRpcRes, error) {
	return market.JsonResult(l.svcCtx.MarketStore.AsMarketAppRole.Query().Where(asmarketapprole.IsDeleted(0), asmarketapprole.AppID(in.AppId)).All(l.ctx))
}
