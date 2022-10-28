package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInnerUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInnerUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInnerUserLogic {
	return &AddInnerUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增内部机构并分配人员
func (l *AddInnerUserLogic) AddInnerUser(in *system.AddInnerUserReq) (*system.CommonRpcRes, error) {
	unitEntity, err := l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.HasTenantxWith(astenant.TenantCode(in.TenantCode))).First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if unitEntity == nil {
		return &system.CommonRpcRes{}, errors.New("unit not found by tenantcode :" + in.TenantCode)
	}
	if tools.IsNilOrEmpty(in.AgencyCode) {
		in.AgencyCode, err = l.svcCtx.SystemStore.GetAgencyInnerCode(l.ctx, in.TenantCode)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
	}
	isExist, err := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.AgencyCode(in.AgencyCode), asinneragency.IsDeleted(0), asinneragency.TenantCode(in.TenantCode)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if isExist {
		return &system.CommonRpcRes{}, errors.New("agencycode is Exist")
	}
	parent, err := l.svcCtx.SystemStore.AsInnerAgency.Get(l.ctx, in.ParentId)
	if err != nil {
		parent, err = l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.ParentIDIsNil()).Limit(1).First(l.ctx)
	}
	if err != nil {
		return &system.CommonRpcRes{}, errors.New("unkown error.")
	}
	igEntity, err := l.svcCtx.SystemStore.AsInnerAgency.Create().
		SetAgencyCode(in.AgencyCode).
		SetAgencyName(in.AgencyName).
		SetParent(parent).
		SetTenantCode(in.TenantCode).
		AddPersonIDs(in.PersonIds...).
		Save(l.ctx)
	return system.JsonResult(igEntity, err)
}
