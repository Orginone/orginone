package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllgroupInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllgroupInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllgroupInfoByIdLogic {
	return &GetAllgroupInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取集团信息
func (l *GetAllgroupInfoByIdLogic) GetAllgroupInfoById(in *system.PrimaryKeyReq) (*system.CommonRpcRes, error) {
	groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.Id)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if groupEntity == nil {
		return &system.CommonRpcRes{}, errors.New("group not found.")
	}
	unitEntity := l.svcCtx.SystemStore.GetUnitNameByTenantCode(l.ctx, groupEntity.TenantCode)
	return system.JsonResult(&model.AllGroupRecord{
		AsAllGroup: groupEntity,
		IsCreate:   1,
		UnitName:   unitEntity.UnitName,
		LinkMan:    unitEntity.LinkMan,
		LinkPhone:  unitEntity.LinkPhone,
	}, err)
}
