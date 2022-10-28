package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRoleLogic {
	return &DeleteUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户角色
func (l *DeleteUserRoleLogic) DeleteUserRole(in *system.UpdateUserRoleReq) (*system.CommonRpcRes, error) {
	count, err := l.svcCtx.SystemStore.AsUser.Update().Where(asuser.IDIn(in.UserIds...)).RemoveRoleIDs(in.RoleIds...).Save(l.ctx)
	return system.JsonResult(count > 0, err)
}
