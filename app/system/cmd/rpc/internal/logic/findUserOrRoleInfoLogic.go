package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asrole"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserOrRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserOrRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserOrRoleInfoLogic {
	return &FindUserOrRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户或角色信息
func (l *FindUserOrRoleInfoLogic) FindUserOrRoleInfo(in *system.RoleUserReq) (*system.CommonRpcRes, error) {
	var result interface{}
	var err error
	if in.UserId > 0 {
		roleEntitys, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.ID(in.UserId)).QueryRoles().All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		result = roleEntitys
	}
	if in.RoleId > 0 {
		userEntitys, err := l.svcCtx.SystemStore.AsRole.Query().Where(asrole.ID(in.RoleId)).QueryUsers().All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		result = userEntitys
	}
	return system.JsonResult(result, err)
}
