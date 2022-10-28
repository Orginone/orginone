package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAgencyInnerUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAgencyInnerUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAgencyInnerUserLogic {
	return &UpdateAgencyInnerUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改部门层级关系及用户所在部门
func (l *UpdateAgencyInnerUserLogic) UpdateAgencyInnerUser(in *system.UpdateAgencyInnerUserReq) (*system.CommonRpcRes, error) {
	local, err := l.svcCtx.SystemStore.AsInnerAgency.Get(l.ctx, in.DeptId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if local == nil {
		return &system.CommonRpcRes{}, errors.New("inner agency not found !")
	}
	query := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.TenantCode(local.TenantCode), asinneragency.IDNEQ(local.ID))
	//检查部门名称是否重复
	if !tools.IsNilOrEmpty(in.DeptName) {
		nameExist, err := query.Clone().Where(asinneragency.AgencyName(in.DeptName)).Exist(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if nameExist {
			return &system.CommonRpcRes{}, errors.New("dept name has Exist !")
		}
		local.AgencyName = in.DeptName
	}
	//检查部门编号是否重复
	if !tools.IsNilOrEmpty(in.AgencyCode) {
		codeExist, err := query.Clone().Where(asinneragency.AgencyCode(in.AgencyCode)).Exist(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if codeExist {
			return &system.CommonRpcRes{}, errors.New("dept code has Exist !")
		}
		local.AgencyCode = in.AgencyCode
	} else if tools.IsNilOrEmpty(local.AgencyCode) {
		newCode, err := l.svcCtx.SystemStore.GetAgencyInnerCode(l.ctx, local.TenantCode)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		local.AgencyCode = newCode
	}
	if in.PId != -1 {
		local.ParentID = in.PId
	}
	_, err = l.svcCtx.SystemStore.AsInnerAgency.UpdateOne(local).
		SetAgencyName(local.AgencyName).
		SetAgencyCode(local.AgencyCode).
		SetParentID(local.ParentID).ClearPersons().AddPersonIDs(in.PersonIds...).Save(l.ctx)
	if err != nil {
		return system.JsonResult(false, err)
	}
	return system.JsonResult(true, nil)
}
