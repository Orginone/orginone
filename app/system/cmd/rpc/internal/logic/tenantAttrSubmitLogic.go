package logic

import (
	"context"
	"encoding/json"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/astenantattr"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantAttrSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantAttrSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantAttrSubmitLogic {
	return &TenantAttrSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增或编辑分类信息
func (l *TenantAttrSubmitLogic) TenantAttrSubmit(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var entity *schema.AsTenantAttr
	err := json.Unmarshal([]byte(in.Json), &entity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	exist, err := l.svcCtx.SystemStore.AsTenantAttr.Query().Where(astenantattr.IsDeleted(0), astenantattr.AttrName(entity.AttrName), astenantattr.IDNEQ(entity.ID)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if exist {
		return system.Result("false", errors.New("租户分类名称已存在，更换后重试"))
	}
	updater := l.svcCtx.SystemStore.AsTenantAttr.UpdateOneID(entity.ID)
	tools.SchemaUpdateAny(updater.Mutation(), entity, "id")
	obj, err := updater.Save(l.ctx)
	if obj == nil {
		obj, err = l.svcCtx.SystemStore.AsTenantAttr.Create().SetAttrName(entity.AttrName).SetAttrRemark(entity.AttrRemark).SetParentID(entity.ParentID).Save(l.ctx)
	}
	return system.JsonResult(obj, nil)
}
