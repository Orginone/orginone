package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/astenant"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllGroupByTentantCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllGroupByTentantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllGroupByTentantCodeLogic {
	return &FindAllGroupByTentantCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取租户下集团信息
func (l *FindAllGroupByTentantCodeLogic) FindAllGroupByTentantCode(in *system.FindAllGroupByTenantCodeReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0),
		asallgroup.HasAllTenantsWith(
			asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.TypeIn(in.RelationTypeRange...),
			asgrouptenantrelations.StatusIn(in.StatusRange...), asgrouptenantrelations.HasGroupWith(),
			asgrouptenantrelations.HasTenantWith(astenant.TenantCode(in.TenantCode), astenant.IsDeleted(0))))
	switch in.Type {
	case 1:
		query = l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.Or(asallgroup.TenantCode(in.TenantCode),
			asallgroup.HasAllTenantsWith(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.TypeIn(in.RelationTypeRange...),
				asgrouptenantrelations.StatusIn(in.StatusRange...), asgrouptenantrelations.HasGroupWith(),
				asgrouptenantrelations.HasTenantWith(astenant.TenantCode(in.TenantCode), astenant.IsDeleted(0)))))
		break
	case 2:
		query = query.Where(asallgroup.TenantCode(in.TenantCode))
		break
	case 3:
		query = query.Where(asallgroup.TenantCodeNEQ(in.TenantCode))
		break
	}
	if !tools.IsNilOrEmpty(in.Page.Filter) {
		query = query.Where(asallgroup.GroupNameContains(in.Page.Filter))
	}
	count, err := query.Count(l.ctx)
	groupEntitys, err := query.WithAllTenants(func(agtrq *schema.AsGroupTenantRelationsQuery) {
		agtrq.Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.TypeIn(in.RelationTypeRange...),
			asgrouptenantrelations.StatusIn(in.StatusRange...), asgrouptenantrelations.HasTenantWith(astenant.TenantCode(in.TenantCode), astenant.IsDeleted(0)))
	}).Order(schema.Desc(asallgroup.FieldCreateTime)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	records := make([]*model.AllGroupRecord, 0)
	linq.From(groupEntitys).SelectT(func(g *schema.AsAllGroup) *model.AllGroupRecord {
		var isCreate int64 = 0
		if g.TenantCode == in.TenantCode {
			isCreate = 1
		}
		if len(g.Edges.AllTenants) > 0 {
			g.Status = g.Edges.AllTenants[0].Status
		}
		g.Edges = schema.AsAllGroupEdges{}
		unitEntity := l.svcCtx.SystemStore.GetUnitNameByTenantCode(l.ctx, g.TenantCode)
		return &model.AllGroupRecord{
			AsAllGroup: g,
			IsCreate:   isCreate,
			UnitName:   unitEntity.UnitName,
			LinkMan:    unitEntity.LinkMan,
			LinkPhone:  unitEntity.LinkPhone,
		}
	}).ToSlice(&records)
	//FIXME find group of group topic
	return system.PageResult(in.Page, int64(count), records, err)
}
