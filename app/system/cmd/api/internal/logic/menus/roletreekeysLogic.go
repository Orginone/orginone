package menus

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoletreekeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoletreekeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoletreekeysLogic {
	return RoletreekeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoletreekeysLogic) Roletreekeys(req types.RoleTreeKeysReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.RoleTreeKeysMenu(l.ctx, &system.RoleTreeKeysMenuReq{
		RoleIds: tools.ArrayStrToInt64(strings.Split(req.RoleIds, ",")),
	})
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	var keys []string
	err = json.Unmarshal([]byte(rpcres.Json), &keys)
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.JsonResult(keys, err)
}
