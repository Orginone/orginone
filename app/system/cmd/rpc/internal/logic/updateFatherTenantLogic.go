package logic

import (
	"context"
	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFatherTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFatherTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFatherTenantLogic {
	return &UpdateFatherTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量调整租户上级
func (l *UpdateFatherTenantLogic) UpdateFatherTenant(in *system.UpdateFatherTenantReq) (*system.CommonRpcRes, error) {
	sourceGroupId, err := l.svcCtx.SystemStore.GetTopGroupId(l.ctx, in.GroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return l.svcCtx.SystemStore.GroupPullTenants(l.ctx, in.GroupId, sourceGroupId, in.TenantIds)
}
