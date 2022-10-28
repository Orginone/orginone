package logic

import (
	"context"
	"orginone/common/schema/asproperties"
	"orginone/common/schema/aspropertiesdistribution"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePropertiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemovePropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePropertiesLogic {
	return &RemovePropertiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除性质
func (l *RemovePropertiesLogic) RemoveProperties(in *system.RemovePropertiesReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	_, err = tx.AsProperties.Update().Where(asproperties.IsDeleted(0), asproperties.IDIn(in.PropertiesIds...)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	_, err = tx.AsPropertiesDistribution.Update().Where(
		aspropertiesdistribution.PropertiesIDIn(in.PropertiesIds...),
		aspropertiesdistribution.IsDeleted(0),
	).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, nil
	}
	tx.Commit()
	return &system.CommonRpcRes{}, nil
}
