package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type DistributeDeptPersonsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDistributeDeptPersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DistributeDeptPersonsLogic {
	return &DistributeDeptPersonsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分配部门人员
func (l *DistributeDeptPersonsLogic) DistributeDeptPersons(in *system.DistributeDeptPersonsReq) (*system.CommonRpcRes, error) {
	agencyExist, err := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.ID(in.DeptId), asinneragency.TenantCode(in.TenantCode)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if agencyExist {
		_, err = l.svcCtx.SystemStore.AsPerson.Update().Where(asperson.IDIn(in.PersonIdList...),
			asperson.HasUserxWith(asuser.Not(asuser.HasWorkingDatasWith(asworkingdata.IsDeleted(0), asworkingdata.Type(in.Type))))).ClearAgencys().AddAgencyIDs(in.DeptId).Save(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		return system.JsonResult(true, nil)
	}
	return system.JsonResult(false, errors.New("部门与组织不匹配"))
}
