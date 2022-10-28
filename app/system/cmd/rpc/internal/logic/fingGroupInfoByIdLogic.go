package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type FingGroupInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFingGroupInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FingGroupInfoByIdLogic {
	return &FingGroupInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取集团信息
func (l *FingGroupInfoByIdLogic) FingGroupInfoById(in *system.PrimaryKeyReq) (*system.CommonRpcRes, error) {
	groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.ID(in.Id)).First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantEntity, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCode(groupEntity.TenantCode),
		astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(model.NewGroupUnitInfo(1, true).SetParentId(0).
		SetTenant(tenantEntity).SetGroup(groupEntity).SetManagerCode(groupEntity.TenantCode), err)
}
