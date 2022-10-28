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

type SubmitMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitMenuLogic {
	return &SubmitMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增或修改菜单
func (l *SubmitMenuLogic) SubmitMenu(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var entity *schema.AsMenu
	err := json.Unmarshal([]byte(in.Json), &entity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	updater := l.svcCtx.SystemStore.AsMenu.UpdateOneID(entity.ID)
	tools.SchemaUpdateAny(updater.Mutation(), entity, "id", "parent_id")
	obj, err := updater.SetStatus(2).Save(l.ctx)
	if obj == nil {
		roleId, err := l.svcCtx.SystemStore.GetRoleId(l.ctx, 1)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		create := l.svcCtx.SystemStore.AsMenu.Create().SetAlias(entity.Alias).
			SetCategory(entity.Category).SetIcon(entity.Icon).SetIsOpen(entity.IsOpen).SetName(entity.Name).
			SetPath(entity.Path).SetReformStatus(entity.ReformStatus).
			SetRemark(entity.Remark).SetSort(entity.Sort).AddRoleIDs(roleId)
		if entity.ParentID > 0 {
			create.SetParentID(entity.ParentID)
		}
		obj, err = create.Save(l.ctx)
	}
	return system.JsonResult(obj, err)
}
