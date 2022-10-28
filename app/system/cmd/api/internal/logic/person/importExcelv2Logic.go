package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportExcelv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportExcelv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) ImportExcelv2Logic {
	return ImportExcelv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportExcelv2Logic) ImportExcelv2(req types.ImportExcelV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
