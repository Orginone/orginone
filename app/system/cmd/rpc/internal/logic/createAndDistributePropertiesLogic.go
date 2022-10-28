package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asproperties"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAndDistributePropertiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAndDistributePropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAndDistributePropertiesLogic {
	return &CreateAndDistributePropertiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团创建性质并进行分配
func (l *CreateAndDistributePropertiesLogic) CreateAndDistributeProperties(in *system.CreateAndDistributePropertiesReq) (*system.CommonRpcRes, error) {
	datas, err := l.svcCtx.SystemStore.AsProperties.Query().Where(asproperties.GroupID(in.GroupId), asproperties.PropertiesName(in.PropertiesName), asproperties.IsDeleted(0)).All(l.ctx)
	if len(datas) != 0 {
		return &system.CommonRpcRes{}, err
	}
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	igEntity, err := tx.AsProperties.Create().SetGroupID(in.GroupId).SetPropertiesName(in.PropertiesName).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	bulk := make([]*schema.AsPropertiesDistributionCreate, len(in.TenantIds))
	for i, id := range in.TenantIds {
		bulk[i] = tx.AsPropertiesDistribution.Create().
			SetPropertiesID(igEntity.ID).
			SetTenantID(id)
	}
	tx.AsPropertiesDistribution.CreateBulk(bulk...).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &system.CommonRpcRes{}, err
}
