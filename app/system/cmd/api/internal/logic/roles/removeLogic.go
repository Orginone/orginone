package roles

import (
	"context"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

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

func (l *RemoveLogic) Remove(req types.RemoveReq3) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.SystemRpc.RemoveRoleByIds(l.ctx, &system.RemoveRoleByIdsReq{
		RoleIds: tools.ArrayStrToInt64(strings.Split(req.Ids, ",")),
	})
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.Successful(true)
}
