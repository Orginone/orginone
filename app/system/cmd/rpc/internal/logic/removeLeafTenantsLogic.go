package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/astenant"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLeafTenantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLeafTenantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLeafTenantsLogic {
	return &RemoveLeafTenantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除实集团叶子节点租户
func (l *RemoveLeafTenantsLogic) RemoveLeafTenants(in *system.RemoveLeafTenantsReq) (*system.CommonRpcRes, error) {
	allGroupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.ID(in.GroupId),
		asallgroup.IsDeleted(0), asallgroup.TypeNEQ(2)).First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if allGroupEntity == nil {
		return &system.CommonRpcRes{}, errors.New("该集团不存在或不是实集团!")
	}
	tenantEntity, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCode(allGroupEntity.TenantCode)).First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if tenantEntity != nil {
		index := tools.ArrIndexOf(in.TenantIds, tenantEntity.ID)
		if index != -1 {
			return &system.CommonRpcRes{}, errors.New("禁止移出集团管理单位！")
		}
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsGroupTenantRelations.Update().Where(
		asgrouptenantrelations.Type(2),
		asgrouptenantrelations.IsDeleted(0),
		asgrouptenantrelations.StatusIn(102, 106),
		asgrouptenantrelations.SonIDIn(in.TenantIds...),
		asgrouptenantrelations.GroupCodeContains(allGroupEntity.GroupCode),
	).SetIsDeleted(1).Save(l.ctx)
	tCodes, err := tx.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.IDIn(in.TenantIds...)).Select(astenant.FieldTenantCode).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppGroupDistribution.Update().Where(
		asmarketappgroupdistribution.GroupID(in.GroupId),
		asmarketappgroupdistribution.TenantIDIn(tCodes...),
	).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &system.CommonRpcRes{}, err
}
