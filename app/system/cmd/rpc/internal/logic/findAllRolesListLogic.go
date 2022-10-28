package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asrole"
	"orginone/common/schema/astenant"
	"orginone/common/schema/astenantattrrole"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllRolesListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllRolesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllRolesListLogic {
	return &FindAllRolesListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有角色列表
func (l *FindAllRolesListLogic) FindAllRolesList(in *system.FindAllRolesListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsTenantAttrRole.Query()
	if !tools.IsNilOrEmpty(in.TenantCode) {
		tenantEntity, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantCode(in.TenantCode), astenant.IsDeleted(0)).First(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		query = query.Where(astenantattrrole.AttrID(tenantEntity.TenantType))
	}
	return system.JsonResult(query.QueryRolex().Where(asrole.IsDeleted(0)).Order(schema.Desc(asrole.FieldCreateTime)).All(l.ctx))
}
