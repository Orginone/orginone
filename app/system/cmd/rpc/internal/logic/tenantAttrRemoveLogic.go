package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenantattr"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantAttrRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantAttrRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantAttrRemoveLogic {
	return &TenantAttrRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除分类信息
func (l *TenantAttrRemoveLogic) TenantAttrRemove(in *system.TenantAttrRemoveReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsTenantAttr.Update().Where(astenantattr.IDIn(in.Id...)).ClearAttrRoles().SetIsDeleted(1).Save(l.ctx))
}
