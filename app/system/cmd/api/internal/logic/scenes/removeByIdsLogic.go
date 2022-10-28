package scenes

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveByIdsLogic {
	return RemoveByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByIdsLogic) RemoveByIds(req types.RemoveByIdsReq3) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.RemoveScenesByIds(l.ctx, &system.RemoveScenesByIdsReq{
		Ids: tools.ArrayStrToInt64(strings.Split(req.Ids, ",")),
	}))
}
