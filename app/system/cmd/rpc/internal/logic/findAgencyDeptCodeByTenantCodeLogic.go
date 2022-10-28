package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAgencyDeptCodeByTenantCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAgencyDeptCodeByTenantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAgencyDeptCodeByTenantCodeLogic {
	return &FindAgencyDeptCodeByTenantCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询全新的部门编码
func (l *FindAgencyDeptCodeByTenantCodeLogic) FindAgencyDeptCodeByTenantCode(in *system.TenantCodeReq) (*system.CommonRpcRes, error) {
	return system.Result(l.svcCtx.SystemStore.GetAgencyInnerCode(l.ctx, in.TenantCode))
}
