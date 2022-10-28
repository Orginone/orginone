package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantByAccountLogic {
	return &FindTenantByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据账户查找租户
func (l *FindTenantByAccountLogic) FindTenantByAccount(in *system.ByAccountReq) (*system.CommonRpcRes, error) {
	tenantCodes, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.PhoneNumber(in.Account),
		asuser.IsDeleted(0), asuser.TenantApplyingStateNEQ(1)).Select(asuser.FieldTenantCode).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if len(tenantCodes) < 1 {
		return &system.CommonRpcRes{}, errors.New("user not join tenant.")
	}
	tenantEntitys, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantCodeIn(tenantCodes...), astenant.IsDeleted(0)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if len(tenantEntitys) < 1 {
		return &system.CommonRpcRes{}, errors.New("tenant is deleted.")
	}
	return system.JsonResult(tenantEntitys, err)
}
