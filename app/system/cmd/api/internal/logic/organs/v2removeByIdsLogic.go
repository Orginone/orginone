package organs

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2removeByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2removeByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2removeByIdsLogic {
	return V2removeByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2removeByIdsLogic) V2removeByIds(req types.V2RemoveByIdsReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.RemoveOrgansByIds(l.ctx, &system.V2RemoveByIdsReq{
		Ids: tools.ArrayStrToInt64(strings.Split(req.Ids, ",")),
	}))
}
