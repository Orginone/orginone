package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMasterUnitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetMasterUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMasterUnitLogic {
	return &SetMasterUnitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置主租户
func (l *SetMasterUnitLogic) SetMasterUnit(in *system.SetMasterUnitReq) (*system.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return system.Result("false", err)
	}
	_, err = tx.AsPerson.Update().Where(asperson.PhoneNumber(in.Account)).SetIsMaster(0).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	_, err = tx.AsPerson.Update().Where(asperson.PhoneNumber(in.Account), asperson.TenantCode(in.TenantCode)).SetIsMaster(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	err = tx.Commit()
	return system.Result("true", err)
}
