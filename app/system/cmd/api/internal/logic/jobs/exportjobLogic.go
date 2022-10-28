package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExportjobLogic {
	return ExportjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportjobLogic) Exportjob(req types.ExportJobReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
