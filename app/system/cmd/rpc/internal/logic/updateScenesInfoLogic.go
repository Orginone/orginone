package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/astenant"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScenesInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateScenesInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScenesInfoLogic {
	return &UpdateScenesInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新租户信息
func (l *UpdateScenesInfoLogic) UpdateScenesInfo(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var tenantEntity *schema.AsTenant
	err := json.Unmarshal([]byte(in.Json), &tenantEntity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	local, err := l.svcCtx.SystemStore.AsTenant.Get(l.ctx, tenantEntity.ID)
	updater := l.svcCtx.SystemStore.AsTenant.UpdateOne(local)
	tools.SchemaUpdateAny(updater.Mutation(), tenantEntity, astenant.FieldID)
	return system.JsonResult(updater.Save(l.ctx))
}
