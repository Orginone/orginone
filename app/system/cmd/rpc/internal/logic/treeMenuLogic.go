package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asmenu"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTreeMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeMenuLogic {
	return &TreeMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取菜单树形结构
func (l *TreeMenuLogic) TreeMenu(in *system.Nil) (*system.CommonRpcRes, error) {
	menuEntitys, err := l.svcCtx.SystemStore.AsMenu.Query().Where(asmenu.IsDeleted(0), asmenu.Category(1)).All(l.ctx)
	return system.JsonResult(loopBuildTreeByIds(l.ctx, menuEntitys, 0), err)
}

func loopBuildTreeByIds(ctx context.Context, array []*schema.AsMenu, parentID int64) []*model.MenuTree {
	menuTrees := make([]*model.MenuTree, 0)
	for _, item := range array {
		if item.ParentID == parentID {
			menuTree := new(model.MenuTree)
			tools.Copy(menuTree, item)
			menuTree.Children = loopBuildTreeByIds(ctx, array, item.ID)
			menuTrees = append(menuTrees, menuTree)
		}
	}
	return menuTrees
}
