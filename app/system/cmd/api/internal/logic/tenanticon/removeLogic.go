package tenanticon

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req types.RemoveReq5) (resp *types.CommonResponse, err error) {
	rpcRep, err := l.svcCtx.SystemRpc.RemoveTenantIcon(l.ctx, &system.CommonRpcReq{Json: req.Ids})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	urls := make([]string, 0)
	json.Unmarshal([]byte(rpcRep.Json), &urls)
	for _, v := range urls {
		os.Remove(v)
	}
	return types.JsonResult(true, err)
}
