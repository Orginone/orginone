package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExportExcelLogic {
	return ExportExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportExcelLogic) ExportExcel(req types.ExportExcelReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
