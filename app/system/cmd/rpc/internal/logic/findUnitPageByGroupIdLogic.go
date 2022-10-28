package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUnitPageByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUnitPageByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUnitPageByGroupIdLogic {
	return &FindUnitPageByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单位列表根据集团id
func (l *FindUnitPageByGroupIdLogic) FindUnitPageByGroupId(in *system.FindUnitPageReq) (*system.CommonRpcRes, error) {
	group, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.GroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	query := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0),
		asgrouptenantrelations.Type(2), asgrouptenantrelations.ExpiresTimeIsNil(), asgrouptenantrelations.StatusIn(102, 106),
		asgrouptenantrelations.GroupCodeHasPrefix(group.GroupCode)).Order(schema.Asc(asgrouptenantrelations.FieldSort)).
		QueryTenant().Where(astenant.IsDeleted(0), astenant.HasUnitWith(asunit.IsDeleted(0), asunit.UnitNameContains(in.Page.Filter))).
		WithUnit()
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantEntitys, err := query.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if in.Page.Offset == 0 && len(tenantEntitys) == 0 {
		adminTenants, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantCode(group.TenantCode)).WithUnit().All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		tenantEntitys = append(adminTenants, tenantEntitys...)
	}
	var unitModels = make([]*model.UnitModel, 0)
	linq.From(tenantEntitys).SelectT(func(t *schema.AsTenant) *model.UnitModel {
		return &model.UnitModel{
			AsUnit:     t.Edges.Unit,
			TenantType: t.TenantType,
			TenantCode: t.TenantCode,
		}
	}).ToSlice(&unitModels)
	return system.PageResult(in.Page, total+1, unitModels, err)
}
