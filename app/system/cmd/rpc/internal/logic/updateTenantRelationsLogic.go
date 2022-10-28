package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools/linq"

	linqx "github.com/ahmetb/go-linq/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTenantRelationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTenantRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTenantRelationsLogic {
	return &UpdateTenantRelationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type TenantRelationData struct {
	socialCreditCode string
	parentGroupName  string
}

// 批量导入租户与集团关系
func (l *UpdateTenantRelationsLogic) UpdateTenantRelations(in *system.UpdateTenantRelationsReq) (*system.CommonRpcRes, error) {
	type GroupTenant struct {
		GroupName         string
		SocialCreditCodes []string
	}
	groupTenants := make([]*GroupTenant, 0)
	linq.From(in.TenantRelationDatas).GroupByT(
		func(e *system.TenantRelationData) string {
			return e.ParentGroupName
		}, func(e *system.TenantRelationData) string {
			return e.SocialCreditCode
		}).SelectT(func(g linqx.Group) *GroupTenant {
		group := new(GroupTenant)
		group.SocialCreditCodes = make([]string, 0)
		group.GroupName = g.Key.(string)
		for _, code := range g.Group {
			group.SocialCreditCodes = append(group.SocialCreditCodes, code.(string))
		}
		return group
	}).ToSlice(&groupTenants)
	for _, item := range groupTenants {
		groupId, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.GroupName(item.GroupName)).Limit(1).Select(asallgroup.FieldID).Int(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		tenantIds, err := l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.SocialCreditCodeIn(item.SocialCreditCodes...)).QueryTenantx().Where(astenant.IsDeleted(0)).Select(astenant.FieldID).Int64s(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		_, err = l.svcCtx.SystemStore.GroupPullTenants(l.ctx, int64(groupId), in.SourceGroupId, tenantIds)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
	}
	return system.Result("true", nil)
}
