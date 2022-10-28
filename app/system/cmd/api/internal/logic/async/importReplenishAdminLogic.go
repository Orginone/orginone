package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportReplenishAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportReplenishAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportReplenishAdminLogic {
	return ImportReplenishAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportReplenishAdminLogic) ImportReplenishAdmin(req types.ImportReplenishAdminReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
