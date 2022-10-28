package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportExcel2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportExcel2Logic(ctx context.Context, svcCtx *svc.ServiceContext) ImportExcel2Logic {
	return ImportExcel2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportExcel2Logic) ImportExcel2(req types.ImportExcel2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
