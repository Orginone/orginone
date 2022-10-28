package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrgansInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrgansInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrgansInfoLogic {
	return &UpdateOrgansInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新单位信息
func (l *UpdateOrgansInfoLogic) UpdateOrgansInfo(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var unitEntity *schema.AsUnit
	err := json.Unmarshal([]byte(in.Json), &unitEntity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	local, err := l.svcCtx.SystemStore.AsUnit.Get(l.ctx, unitEntity.ID)
	updater := l.svcCtx.SystemStore.AsUnit.UpdateOne(local)
	tools.SchemaUpdateAny(updater.Mutation(), unitEntity, "id")
	return system.JsonResult(updater.Save(l.ctx))
}
