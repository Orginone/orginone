package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportnoadminunittenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportnoadminunittenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportnoadminunittenantLogic {
	return ImportnoadminunittenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportnoadminunittenantLogic) Importnoadminunittenant(req types.ImportNoAdminUnitTenantReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
