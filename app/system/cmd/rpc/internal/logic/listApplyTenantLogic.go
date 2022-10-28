package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListApplyTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListApplyTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApplyTenantLogic {
	return &ListApplyTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 申请加入集团租户列表
func (l *ListApplyTenantLogic) ListApplyTenant(in *system.ListApplyTenantReq) (*system.CommonRpcRes, error) {
	flags := make([]int64, 0)
	switch in.Flag {
	case 101, 102, 103, 104, 106: //101-待审核,102-审核通过,103-审核拒绝,104-取消申请,106-创建者
		flags = append(flags, in.Flag)
		break
	case 105: //105-全部(101-104)
		flags = append(flags, 101, 102, 103, 104)
		break
	case 107: //107-已办(102-104)
		flags = append(flags, 102, 103, 104)
		break
	default:
		return &system.CommonRpcRes{}, errors.New("flag is not supported.")
	}
	query := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0),
		asgrouptenantrelations.StatusIn(flags...), asgrouptenantrelations.ParentID(in.GroupId))
	if in.Flag == 105 || in.Flag == 107 {
		query = query.Where(asgrouptenantrelations.Type(2))
	}
	queryTenant := query.QueryTenant().Where(astenant.IsDeleted(0), astenant.HasUnitWith(asunit.IsDeleted(0), asunit.UnitNameContains(in.Page.Filter)))
	total, err := queryTenant.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	tenantEntitys, err := queryTenant.WithUnit().WithAllGroups(func(agtrq *schema.AsGroupTenantRelationsQuery) {
		agtrq.Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.StatusIn(flags...), asgrouptenantrelations.ParentID(in.GroupId))
		if in.Flag == 105 || in.Flag == 107 {
			agtrq.Where(asgrouptenantrelations.Type(2))
		}
	}).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	unitModels := make([]*model.UnitModel, 0)
	for _, tenant := range tenantEntitys {
		unitModel := new(model.UnitModel)
		unitModel.AsUnit = tenant.Edges.Unit
		if len(tenant.Edges.AllGroups) > 0 {
			unitModel.Flag = tenant.Edges.AllGroups[0].Status
		}
		unitModels = append(unitModels, unitModel)
	}
	return system.PageResult(in.Page, total, unitModels, err)
}
