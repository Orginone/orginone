package grouptenantrelations

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantsjoingroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantsjoingroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantsjoingroupLogic {
	return TenantsjoingroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantsjoingroupLogic) Tenantsjoingroup(req types.TenantsJoinGroupReq) (resp *types.CommonResponse, err error) {
	tenantIds := tools.ArrayStrToInt64(strings.Split(req.TenantIds, ","))
	return types.CommonResult(l.svcCtx.SystemRpc.TenantsJoinGroup(l.ctx, &system.TenantsJoinGroupReq{
		GroupId:        req.GroupId,
		Flag:           -1,
		SourcegGroupId: req.SourceGroupId,
		TenantIds:      tenantIds,
	}))
}
