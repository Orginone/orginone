package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantpersonLogic {
	return TenantpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantpersonLogic) Tenantperson(req types.TenantPersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
