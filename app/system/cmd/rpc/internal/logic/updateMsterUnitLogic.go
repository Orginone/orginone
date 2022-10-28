package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMsterUnitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMsterUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMsterUnitLogic {
	return &UpdateMsterUnitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 主单位导入
func (l *UpdateMsterUnitLogic) UpdateMsterUnit(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	units := make([]model.UpdatePersonMasterData, 0)
	err := json.Unmarshal([]byte(in.Json), &units)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if len(units) > 0 {
		tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		phones := make([]string, 0)
		for _, v := range units {
			phones = append(phones, v.PhoneNumber)
		}
		_, err = tx.AsPerson.Update().Where(asperson.PhoneNumberIn(phones...), asperson.TenantCode(units[0].TenantCode)).SetIsMaster(0).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, err
		}
		for _, v := range units {
			_, err = tx.AsPerson.Update().Where(asperson.PhoneNumber(v.PhoneNumber), asperson.TenantCode(v.TenantCode)).SetIsMaster(1).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
		}
		return system.JsonResult(true, tx.Commit())
	}
	return &system.CommonRpcRes{}, nil
}
