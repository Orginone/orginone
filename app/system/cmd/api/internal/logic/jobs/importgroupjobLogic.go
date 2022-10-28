package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportgroupjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportgroupjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportgroupjobLogic {
	return ImportgroupjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportgroupjobLogic) Importgroupjob(req types.ImportGroupjobReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
