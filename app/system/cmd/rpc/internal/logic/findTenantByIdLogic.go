package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantByIdLogic {
	return &FindTenantByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取租户详情
func (l *FindTenantByIdLogic) FindTenantById(in *system.TenantByIdReq) (*system.CommonRpcRes, error) {
	tenantEntity, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.ID(in.Id), astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(tenantEntity, err)
}
