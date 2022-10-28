package propertiesdistribution

import (
	"context"
	"orginone/common/rpcs/system"
	"orginone/common/tools"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovePropertiesLogic {
	return RemovePropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePropertiesLogic) RemoveProperties(req types.RemovePropertiesReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.RemoveProperties(l.ctx, &system.RemovePropertiesReq{
		PropertiesIds: tools.ArrayStrToInt64(strings.Split(req.PropertiesIds, ",")),
	}))
}
