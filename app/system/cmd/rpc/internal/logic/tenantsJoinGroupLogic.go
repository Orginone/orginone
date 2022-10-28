package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantsJoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantsJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantsJoinGroupLogic {
	return &TenantsJoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 租户加入集团
func (l *TenantsJoinGroupLogic) TenantsJoinGroup(in *system.TenantsJoinGroupReq) (*system.CommonRpcRes, error) {
	if in.SourcegGroupId > 0 {
		return l.svcCtx.SystemStore.GroupResetTenants(l.ctx, in.GroupId, in.SourcegGroupId, in.TenantIds)
	}
	switch in.Flag {
	case 101:
		tenantIds, err := l.tenantIdDistinct(in)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		in.TenantIds = tenantIds
		break
	case 102:
		return system.Result("true", l.svcCtx.SystemStore.CreateGroupTenantRelation(l.ctx, in.GroupId, in.TenantIds, in.Flag))
	case 103:
		_, err := l.svcCtx.SystemStore.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.Status(101),
			asgrouptenantrelations.Type(2), asgrouptenantrelations.ParentID(in.GroupId), asgrouptenantrelations.SonIDIn(in.TenantIds...),
			asgrouptenantrelations.ExpiresTimeIsNil()).SetStatus(in.Flag).Save(l.ctx)
		return system.Result("true", err)
	case 104:
		localIds, err := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.SonIDIn(in.TenantIds...),
			asgrouptenantrelations.Status(101), asgrouptenantrelations.Type(2), asgrouptenantrelations.ParentID(in.GroupId)).
			Select(asgrouptenantrelations.FieldSonID).Int64s(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if len(localIds) < 1 {
			return &system.CommonRpcRes{}, errors.New("tenants is not request.")
		}
		in.TenantIds = localIds
		break
	default:
		return &system.CommonRpcRes{}, errors.New("flag is not supported.")
	}
	if len(in.TenantIds) < 1 {
		return &system.CommonRpcRes{}, errors.New("tanents join group faild.")
	}
	err := l.svcCtx.SystemStore.TenantsJoinGroup(l.ctx, in.TenantIds, in.GroupId, in.Flag)
	return system.Result("true", err)
}

//去重
func (l *TenantsJoinGroupLogic) tenantIdDistinct(in *system.TenantsJoinGroupReq) ([]int64, error) {
	localIds, err := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.SonIDIn(in.TenantIds...),
		asgrouptenantrelations.StatusIn(101, 102, 106), asgrouptenantrelations.Type(2), asgrouptenantrelations.ParentID(in.GroupId)).
		Select(asgrouptenantrelations.FieldSonID).Int64s(l.ctx)
	if err != nil {
		return in.TenantIds, err
	}
	if len(localIds) > 0 {
		linq.From(in.TenantIds).WhereT(func(i int64) bool {
			ok := true
			for _, id := range localIds {
				if id == i {
					return false
				}
			}
			return ok
		}).Distinct().ToSlice(in.TenantIds)
	}
	return in.TenantIds, nil
}
