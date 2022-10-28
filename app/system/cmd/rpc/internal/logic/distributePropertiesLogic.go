package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/aspropertiesdistribution"

	"github.com/zeromicro/go-zero/core/logx"
)

type DistributePropertiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDistributePropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DistributePropertiesLogic {
	return &DistributePropertiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分配性质给租户
func (l *DistributePropertiesLogic) DistributeProperties(in *system.DistributePropertiesReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsProperties.UpdateOneID(in.PropertiesId).
		SetPropertiesName(in.PropertiesName).
		Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsPropertiesDistribution.Update().Where(
		aspropertiesdistribution.PropertiesID(in.PropertiesId),
		aspropertiesdistribution.IsDeleted(0),
	).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	bulk := make([]*schema.AsPropertiesDistributionCreate, len(in.TenantIds))
	for i, id := range in.TenantIds {
		bulk[i] = tx.AsPropertiesDistribution.Create().
			SetPropertiesID(in.PropertiesId).
			SetTenantID(id)
	}
	_, err = tx.AsPropertiesDistribution.CreateBulk(bulk...).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult("", tx.Commit())
}
