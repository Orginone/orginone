package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportreplenishadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportreplenishadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportreplenishadminLogic {
	return ImportreplenishadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportreplenishadminLogic) Importreplenishadmin(req types.ImportReplenishAdminReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
