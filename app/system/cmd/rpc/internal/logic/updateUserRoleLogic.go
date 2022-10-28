package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asrole"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新人员角色
func (l *UpdateUserRoleLogic) UpdateUserRole(in *system.UpdateUserRoleReq) (*system.CommonRpcRes, error) {
	count, err := l.svcCtx.SystemStore.AsRole.Query().Where(asrole.IDIn(in.RoleIds...), asrole.IsDeleted(0)).Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if count != int64(len(in.RoleIds)) {
		return &system.CommonRpcRes{}, errors.New("roleIds中包含不存在的roleId!")
	}
	_, err = l.svcCtx.SystemStore.AsUser.Update().Where(asuser.IDIn(in.UserIds...)).ClearRoles().AddRoleIDs(in.RoleIds...).Save(l.ctx)
	return &system.CommonRpcRes{}, err
}
