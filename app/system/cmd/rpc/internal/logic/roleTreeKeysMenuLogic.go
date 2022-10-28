package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asmenu"
	"orginone/common/schema/asrole"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleTreeKeysMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleTreeKeysMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleTreeKeysMenuLogic {
	return &RoleTreeKeysMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 角色权限
func (l *RoleTreeKeysMenuLogic) RoleTreeKeysMenu(in *system.RoleTreeKeysMenuReq) (*system.CommonRpcRes, error) {
	mIds, err := l.svcCtx.SystemStore.AsRole.Query().Where(asrole.IDIn(in.RoleIds...)).
		QueryMenus().Where(asmenu.IsDeleted(0)).Select(asmenu.FieldID).Strings(l.ctx)
	if mIds == nil {
		mIds = make([]string, 0)
	}
	return system.JsonResult(mIds, err)
}
