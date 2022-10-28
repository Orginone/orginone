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

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增或修改角色
func (l *UpdateRoleLogic) UpdateRole(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var entity *schema.AsRole
	err := json.Unmarshal([]byte(in.Json), &entity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	updater := l.svcCtx.SystemStore.AsRole.UpdateOneID(entity.ID)
	tools.SchemaUpdateAny(updater.Mutation(), entity, "id")
	obj, err := updater.Save(l.ctx)
	if obj == nil {
		obj, err = l.svcCtx.SystemStore.AsRole.Create().SetRoleAlias(entity.RoleAlias).
			SetRoleDescription(entity.RoleDescription).SetRoleName(entity.RoleName).
			SetSort(entity.Sort).SetStatus(entity.Status).Save(l.ctx)
	}
	return system.JsonResult(obj, err)

}
