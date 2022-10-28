package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/astenantattrrole"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantAttrRoleAllocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantAttrRoleAllocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantAttrRoleAllocLogic {
	return &TenantAttrRoleAllocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 租户组角色分配
func (l *TenantAttrRoleAllocLogic) TenantAttrRoleAlloc(in *system.TenantAttrRoleAllocReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsTenantAttrRole.Delete().Where(astenantattrrole.AttrID(in.AttrId)).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	bulk := make([]*schema.AsTenantAttrRoleCreate, len(in.RoleIds))
	for i, id := range in.RoleIds {
		bulk[i] = l.svcCtx.SystemStore.AsTenantAttrRole.Create().SetAttrID(in.AttrId).SetRoleID(id).SetIsDefault(0)
	}
	_, err = tx.AsTenantAttrRole.CreateBulk(bulk...).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return system.Result("true", err)
}
