package logic

import (
	"context"
	"strings"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinneragency"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindInnerAgencyTreeByTenantCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindInnerAgencyTreeByTenantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindInnerAgencyTreeByTenantCodeLogic {
	return &FindInnerAgencyTreeByTenantCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//查询租户内设机构列表(树)
func (l *FindInnerAgencyTreeByTenantCodeLogic) FindInnerAgencyTreeByTenantCode(in *system.TenantCodeReq) (*system.CommonRpcRes, error) {
	allInnerAgencyEntitys, err := l.svcCtx.SystemStore.AsInnerAgency.Query().
		Where(asinneragency.IsDeleted(0)).
		Where(asinneragency.Or(asinneragency.TenantCode(in.TenantCode), asinneragency.ParentIDIsNil(), asinneragency.ParentID(0))).
		Order(schema.Asc(asinneragency.FieldAgencyCode)).All(l.ctx)
	data, _ := loopMakeAgencyTreeTreeByIds(allInnerAgencyEntitys, 0, in.Filter)
	return system.JsonResult(data, err)
}

func loopMakeAgencyTreeTreeByIds(array []*schema.AsInnerAgency, parentID int64, filter string) ([]*model.InnerAgencyTree, bool) {
	isadd := false
	menuTrees := make([]*model.InnerAgencyTree, 0)
	for _, item := range array {
		if item.ParentID == parentID {
			menuTree := new(model.InnerAgencyTree)
			menuTree.AsInnerAgency = item
			childData, isCExist := loopMakeAgencyTreeTreeByIds(array, item.ID, filter)
			if strings.Contains(menuTree.AgencyName, filter) || isCExist {
				isadd = true
				if isCExist {
					menuTree.Children = childData
				}
				menuTrees = append(menuTrees, menuTree)
			}
		}
	}
	return menuTrees, isadd
}
