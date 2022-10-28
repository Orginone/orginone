package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDistributePropertiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDistributePropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDistributePropertiesLogic {
	return &CreateDistributePropertiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增租户性质
func (l *CreateDistributePropertiesLogic) CreateDistributeProperties(in *system.CreateDistributePropertiesReq) (*system.CommonRpcRes, error) {
	bulk := make([]*schema.AsPropertiesDistributionCreate, len(in.TenantIds))
	for i, id := range in.TenantIds {
		bulk[i] = l.svcCtx.SystemStore.AsPropertiesDistribution.Create().
			SetPropertiesID(in.PropertiesId).
			SetTenantID(id)
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsPropertiesDistribution.CreateBulk(bulk...).Save(l.ctx))
}
