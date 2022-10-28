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

type DeleteRolesToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRolesToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteRolesToUsersLogic {
	return DeleteRolesToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRolesToUsersLogic) DeleteRolesToUsers(req types.DeleteRolesToUsersReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(l.svcCtx.SystemRpc.DeleteUserRole(l.ctx, &system.UpdateUserRoleReq{
		UserIds: tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
		RoleIds: tools.ArrayStrToInt64(strings.Split(req.RoleIds, ",")),
	}))
}
