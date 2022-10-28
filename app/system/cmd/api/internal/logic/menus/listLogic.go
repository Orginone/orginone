package menus

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq5) (resp *types.CommonResponse, err error) {
	// 查询登录的人员信息
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	//查询系统菜单
	rpcres, err := l.svcCtx.SystemRpc.FindRoleMenusByUser(l.ctx, &system.RoleMenusByUserReq{UserId: userId, PlatformId: 1, TenantCode: tenantCode, RoleIds: make([]int64, 0)})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var menuTrees []*model.MenuTree
	err = json.Unmarshal([]byte(rpcres.Json), &menuTrees)
	return types.JsonResult(loopSearchMenuByName(menuTrees, req.Name), err)
}

func loopSearchMenuByName(menuTrees []*model.MenuTree, name string) []*model.MenuTree {
	if tools.IsNilOrEmpty(name) {
		return menuTrees
	}
	newMenus := make([]*model.MenuTree, 0)
	for _, v := range menuTrees {
		if strings.Contains(v.Name, name) {
			newMenus = append(newMenus, v)
		} else if len(v.Children) > 0 {
			v.Children = loopSearchMenuByName(v.Children, name)
			if len(v.Children) > 0 {
				newMenus = append(newMenus, v)
			}
		}
	}
	return newMenus
}
