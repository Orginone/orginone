package person

import (
	"context"
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

func (l *RemoveLogic) Remove(req types.RemoveReq2) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.RemovePersonByIds(l.ctx, &system.RemovePersonByIdsReq{
		Ids: tools.ArrayStrToInt64(strings.Split(req.Ids, ",")),
	}))
}
