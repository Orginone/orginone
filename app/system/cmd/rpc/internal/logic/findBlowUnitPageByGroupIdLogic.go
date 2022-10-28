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

type FindBlowUnitPageByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBlowUnitPageByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBlowUnitPageByGroupIdLogic {
	return &FindBlowUnitPageByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有单位列表（不含子集团）根据集团id
func (l *FindBlowUnitPageByGroupIdLogic) FindBlowUnitPageByGroupId(in *system.FindUnitPageReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0),
		asgrouptenantrelations.Type(2), asgrouptenantrelations.ExpiresTimeIsNil(), asgrouptenantrelations.StatusIn(102, 106),
		asgrouptenantrelations.ParentID(in.GroupId)).Order(schema.Asc(asgrouptenantrelations.FieldSort)).
		QueryTenant().Where(astenant.IsDeleted(0), astenant.HasUnitWith(asunit.IsDeleted(0), asunit.UnitNameContains(in.Page.Filter))).
		WithUnit().Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	teantEntitys, err := query.All(l.ctx)
	var unitModels = make([]*model.UnitModel, 0)
	linq.From(teantEntitys).SelectT(func(t *schema.AsTenant) *model.UnitModel {
		return &model.UnitModel{
			AsUnit:     t.Edges.Unit,
			TenantType: t.TenantType,
			TenantCode: t.TenantCode,
		}
	}).ToSlice(&unitModels)
	return system.PageResult(in.Page, total, unitModels, err)
}
