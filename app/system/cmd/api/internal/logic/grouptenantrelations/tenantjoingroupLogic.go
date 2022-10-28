package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantjoingroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantjoingroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantjoingroupLogic {
	return TenantjoingroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantjoingroupLogic) Tenantjoingroup(req types.TenantJoinGroupReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.TenantsJoinGroup(l.ctx, &system.TenantsJoinGroupReq{
		GroupId:        req.GroupId,
		Flag:           req.Flag,
		SourcegGroupId: -1,
		TenantIds:      []int64{req.TenantId},
	}))
}
