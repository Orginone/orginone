package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asmenu"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleMenusByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleMenusByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleMenusByUserLogic {
	return &FindRoleMenusByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 平台菜单查询
func (l *FindRoleMenusByUserLogic) FindRoleMenusByUser(in *system.RoleMenusByUserReq) (*system.CommonRpcRes, error) {
	allMenuEntitys, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.ID(in.UserId)).QueryRoles().
		QueryMenus().Where(asmenu.IsDeleted(0), asmenu.Category(1)).Order(schema.Asc(asmenu.FieldSort)).All(l.ctx)
	return system.JsonResult(loopMakeMenuTreeByIds(l.ctx, allMenuEntitys, 0), err)
}

func loopMakeMenuTreeByIds(ctx context.Context, array []*schema.AsMenu, parentID int64) []*model.MenuTree {
	menuTrees := make([]*model.MenuTree, 0)
	for _, item := range array {
		if item.ParentID == parentID {
			menuTree := new(model.MenuTree)
			menuTree.AsMenu = item
			menuTree.Children = loopMakeMenuTreeByIds(ctx, array, item.ID)
			menuTrees = append(menuTrees, menuTree)
		}
	}
	return menuTrees
}
