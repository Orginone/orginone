package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenanticon"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantIconListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantIconListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantIconListLogic {
	return &FindTenantIconListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询租户图标
func (l *FindTenantIconListLogic) FindTenantIconList(in *system.FindTenantIconListReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsTenantIcon.Query().Where(astenanticon.TenantCode(in.TenantCode),
		astenanticon.IsDeleted(0), astenanticon.IconContains(in.TenantIcon)).All(l.ctx))
}
