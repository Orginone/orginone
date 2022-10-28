package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTeantByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTeantByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTeantByNameLogic {
	return &SearchTeantByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据名称查询账户
func (l *SearchTeantByNameLogic) SearchTeantByName(in *system.SearchTeantReq) (*system.CommonRpcRes, error) {
	tenantEntitys, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantNameContains(in.TenantName)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(tenantEntitys, err)
}
