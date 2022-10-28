package api

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantlistLogic {
	return TenantlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantlistLogic) Tenantlist(req types.TenantListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
