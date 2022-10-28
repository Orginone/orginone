package propertiesdistribution

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DistributepropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDistributepropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) DistributepropertiesLogic {
	return DistributepropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DistributepropertiesLogic) Distributeproperties(req types.DistributePropertiesReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.DistributeProperties(l.ctx, &system.DistributePropertiesReq{
		TenantIds:      tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
		PropertiesId:   req.PropertiesId,
		PropertiesName: req.PropertiesName,
	}))
}
