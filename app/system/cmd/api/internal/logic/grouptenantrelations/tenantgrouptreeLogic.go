package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantgrouptreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantgrouptreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantgrouptreeLogic {
	return TenantgrouptreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantgrouptreeLogic) Tenantgrouptree(req types.TenantGroupTreeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
