package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asjob"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/asmarketrolemenu"
	"orginone/common/schema/asmarketusedapp"
	"orginone/common/schema/asperson"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMarketAppRoleMenusByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMarketAppRoleMenusByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMarketAppRoleMenusByUserLogic {
	return &FindMarketAppRoleMenusByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户应用菜单树
func (l *FindMarketAppRoleMenusByUserLogic) FindMarketAppRoleMenusByUser(in *system.RoleMenusByUserReq) (*system.CommonRpcRes, error) {
	marketMenuEntitys, err := l.svcCtx.SystemStore.AsMarketMenu.Query().Where(asmarketmenu.HasRoleMenusWith(
		asmarketrolemenu.HasRolexWith(asmarketapprole.HasRoleDistribsWith(asmarketroledistribution.Or(
			//分配给个人担任的岗位的
			asmarketroledistribution.HasJobxWith(asjob.HasPersonsWith(asperson.UserID(in.UserId))),
			//分配给个人所在的机构的
			asmarketroledistribution.HasAgencyxWith(asinneragency.HasPersonsWith(asperson.UserID(in.UserId))),
			//分配给个人所在的租户然后又分配给个人权限的
			asmarketroledistribution.And(
				asmarketroledistribution.UserID(in.UserId),
				asmarketroledistribution.HasRolexWith(asmarketapprole.HasAppxWith(
					asmarketapp.IsDeleted(0), asmarketapp.Platform(in.PlatformId), asmarketapp.Or(
						//租户自身的应用
						asmarketapp.HasAppPurchasesWith(asmarketapppurchase.TenantID(in.TenantCode),
							asmarketapppurchase.IsDeleted(0), asmarketapppurchase.UseStatus(1)),
						//租户所在的组织
						asmarketapp.HasAppGroupDistribsWith(asmarketappgroupdistribution.TenantID(in.TenantCode),
							asmarketappgroupdistribution.IsDeleted(0), asmarketappgroupdistribution.UseStatus(1))))))),
			asmarketroledistribution.IsDeleted(0)), asmarketapprole.IsDeleted(0)), asmarketrolemenu.IsDeleted(0)),
		asmarketmenu.IsDeleted(0)).WithParent().All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(l.menuSortOfUsedAppSort(loopMakeMarketAppTreeByIds(l.ctx, marketMenuEntitys, 0), in.UserId), err)
}

//根据用户常用排序
func (l *FindMarketAppRoleMenusByUserLogic) menuSortOfUsedAppSort(menuTrees []*model.MarketAppMenuTree, userId int64) []*model.MarketAppMenuTree {
	if len(menuTrees) > 0 {
		appIds := make([]int64, 0)
		linq.From(menuTrees).SelectT(func(m *model.MarketAppMenuTree) int64 {
			return m.AppID
		}).ToSlice(&appIds)
		l.svcCtx.SystemStore.AsMarketUsedApp.Delete().Where(asmarketusedapp.UserID(userId), asmarketusedapp.AppIDNotIn(appIds...)).Exec(l.ctx)
		appIds, err := l.svcCtx.SystemStore.AsMarketUsedApp.Query().Where(asmarketusedapp.IsDeleted(0), asmarketusedapp.UserID(userId), asmarketusedapp.AppIDIn(appIds...)).
			Order(schema.Asc(asmarketusedapp.FieldSort)).Select(asmarketusedapp.FieldAppID).Int64s(l.ctx)
		if err == nil {
			for index, appId := range appIds {
				for mIndex, menu := range menuTrees {
					if appId == menu.AppID {
						menuTrees[index], menuTrees[mIndex] = menuTrees[mIndex], menuTrees[index]
						break
					}
				}
			}
		}
	}
	return menuTrees
}

func loopMakeMarketAppTreeByIds(ctx context.Context, array []*schema.AsMarketMenu, parentID int64) []*model.MarketAppMenuTree {
	menuTrees := make([]*model.MarketAppMenuTree, 0)
	if parentID == 0 {
		newArray := make([]*schema.AsMarketMenu, 0)
		for _, v := range array {
			newArray = append(newArray, loopFindParentNode(ctx, v, array)...)
		}
		linq.From(newArray).DistinctByT(func(i *schema.AsMarketMenu) int64 {
			return i.ID
		}).ToSlice(&array)
	}
	for _, item := range array {
		if item.ParentID == parentID {
			menuTree := new(model.MarketAppMenuTree)
			menuTree.AsMarketMenu = item
			menuTree.Children = loopMakeMarketAppTreeByIds(ctx, array, item.ID)
			menuTrees = append(menuTrees, menuTree)
		}
	}
	return menuTrees
}

func loopFindParentNode(ctx context.Context, menu *schema.AsMarketMenu, array []*schema.AsMarketMenu) []*schema.AsMarketMenu {
	items := make([]*schema.AsMarketMenu, 0)
	items = append(items, menu)
	if menu.ParentID > 0 {
		if menu.Edges.Parent == nil {
			for _, item := range array {
				if item.ID == menu.ParentID {
					menu.Edges.Parent = item
				}
			}
			if menu.Edges.Parent == nil {
				temp, _ := menu.QueryParent().First(ctx)
				menu.Edges.Parent = temp
				array = append(array, temp)
			}
		}
		if menu.Edges.Parent != nil {
			items = append(items, loopFindParentNode(ctx, menu.Edges.Parent, array)...)
		}
	}
	return items
}
