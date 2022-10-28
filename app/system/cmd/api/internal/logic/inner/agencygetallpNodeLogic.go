package inner

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetallpNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetallpNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetallpNodeLogic {
	return AgencygetallpNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetallpNodeLogic) AgencygetallpNode(req types.AgencyGetAllPNodeReq) (resp *types.CommonResponse, err error) {
	_, _, TenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	rpcRes, err := l.svcCtx.SystemRpc.FindInnerAgencyTreeByTenantCode(l.ctx, &system.TenantCodeReq{
		TenantCode: TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var innerAgencyTrees []*model.InnerAgencyTree
	err = json.Unmarshal([]byte(rpcRes.Json), &innerAgencyTrees)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(loopMakeAgencyTreeTreeByDeptID(innerAgencyTrees, req.DeptId))
}

func loopMakeAgencyTreeTreeByDeptID(array []*model.InnerAgencyTree, deptID int64) []*model.InnerAgencyTree {
	menuTrees := make([]*model.InnerAgencyTree, 0)
	for _, item := range array {
		if item.ID != deptID {
			menuTree := new(model.InnerAgencyTree)
			menuTree.AsInnerAgency = item.AsInnerAgency
			if len(item.Children) > 0 {
				menuTree.Children = loopMakeAgencyTreeTreeByDeptID(item.Children, deptID)
			}
			menuTrees = append(menuTrees, menuTree)
		}
	}
	return menuTrees
}
