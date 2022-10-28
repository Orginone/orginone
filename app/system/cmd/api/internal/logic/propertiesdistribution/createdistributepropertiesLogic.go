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

type CreatedistributepropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatedistributepropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreatedistributepropertiesLogic {
	return CreatedistributepropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatedistributepropertiesLogic) Createdistributeproperties(req types.CreateDistributePropertiesReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.CreateAndDistributeProperties(l.ctx, &system.CreateAndDistributePropertiesReq{
		TenantIds:      tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
		GroupId:        req.GroupId,
		PropertiesName: req.PropertiesName,
	}))
}
