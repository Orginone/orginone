package roles

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRolesToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRolesToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateRolesToUsersLogic {
	return UpdateRolesToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRolesToUsersLogic) UpdateRolesToUsers(req types.UpdateRolesToUsersReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.UpdateUserRole(l.ctx, &system.UpdateUserRoleReq{
		UserIds: tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
		RoleIds: tools.ArrayStrToInt64(strings.Split(req.RoleIds, ",")),
	}))
}
