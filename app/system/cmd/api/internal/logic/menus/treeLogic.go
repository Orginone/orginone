package menus

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) TreeLogic {
	return TreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeLogic) Tree(req types.Nil) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.TreeMenu(l.ctx, &system.Nil{})
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	var menuTree []*model.MenuTree
	err = json.Unmarshal([]byte(rpcres.Json), &menuTree)
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.JsonResult(menuTree, err)
}
