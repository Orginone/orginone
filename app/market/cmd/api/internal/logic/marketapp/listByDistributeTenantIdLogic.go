package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByDistributeTenantIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByDistributeTenantIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListByDistributeTenantIdLogic {
	return ListByDistributeTenantIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByDistributeTenantIdLogic) ListByDistributeTenantId(req types.ListByDistributeTenantIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
