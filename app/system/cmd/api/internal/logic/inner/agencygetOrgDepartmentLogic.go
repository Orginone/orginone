package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetOrgDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetOrgDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetOrgDepartmentLogic {
	return AgencygetOrgDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetOrgDepartmentLogic) AgencygetOrgDepartment(req types.AgencyGetOrgDepartmentReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
