package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportNoAdminUnitTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportNoAdminUnitTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportNoAdminUnitTenantLogic {
	return ImportNoAdminUnitTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportNoAdminUnitTenantLogic) ImportNoAdminUnitTenant(req types.ImportNoAdminUnitTenantReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
