package menus

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllRoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllRoutesLogic {
	return AllRoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllRoutesLogic) AllRoutes(req types.AllRoutesReq) (resp *types.CommonResponse, err error) {
	// 查询登录的人员信息
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	var roleIds []int64
	linq.From(strings.Split(req.RoleIds, ",")).SelectT(func(s string) int64 {
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			id = 0
		}
		return id
	}).ToSlice(&roleIds)
	//查询系统菜单
	rpcres, err := l.svcCtx.SystemRpc.FindRoleMenusByUser(l.ctx, &system.RoleMenusByUserReq{UserId: userId, PlatformId: req.Platform, TenantCode: tenantCode, RoleIds: roleIds})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var menuTrees []*model.MenuTree
	err = json.Unmarshal([]byte(rpcres.Json), &menuTrees)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	result := make([]interface{}, 0)
	for _, tree := range menuTrees {
		result = append(result, tree)
	}
	//查询应用菜单
	rpcres, err = l.svcCtx.SystemRpc.FindMarketAppRoleMenusByUser(l.ctx, &system.RoleMenusByUserReq{UserId: userId, PlatformId: req.Platform, TenantCode: tenantCode, RoleIds: roleIds})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var marketMenuTrees []*model.MarketAppMenuTree
	err = json.Unmarshal([]byte(rpcres.Json), &marketMenuTrees)
	if err == nil {
		for _, tree := range marketMenuTrees {
			result = append(result, tree)
		}
	}
	return types.Successful(result)
}
