package logic

import (
	"context"
	"strings"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantByNamesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantByNamesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantByNamesLogic {
	return &FindTenantByNamesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过名字查找租户列表
func (l *FindTenantByNamesLogic) FindTenantByNames(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantNameIn(strings.Split(in.Json, ",")...)).WithUnit().All(l.ctx))
}
