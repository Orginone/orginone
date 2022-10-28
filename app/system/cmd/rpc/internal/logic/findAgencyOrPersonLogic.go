package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAgencyOrPersonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAgencyOrPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAgencyOrPersonLogic {
	return &FindAgencyOrPersonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取机构或人员信息
func (l *FindAgencyOrPersonLogic) FindAgencyOrPerson(in *system.AgencyPersonReq) (*system.CommonRpcRes, error) {
	var result interface{}
	var err error
	if in.PersonId > 0 {
		agencyEntitys, err := l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.ID(in.PersonId)).QueryAgencys().All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		result = agencyEntitys
	}
	if in.AgencyId > 0 {
		personEntitys, err := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.ID(in.AgencyId)).QueryPersons().All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		result = personEntitys
	}
	return system.JsonResult(result, err)
}
