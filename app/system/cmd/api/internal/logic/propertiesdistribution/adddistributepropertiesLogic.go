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

type AdddistributepropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdddistributepropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdddistributepropertiesLogic {
	return AdddistributepropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdddistributepropertiesLogic) Adddistributeproperties(req types.AddDistributePropertiesReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.CreateDistributeProperties(l.ctx, &system.CreateDistributePropertiesReq{
		TenantIds:    tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
		PropertiesId: req.PropertiesId,
	}))
}
