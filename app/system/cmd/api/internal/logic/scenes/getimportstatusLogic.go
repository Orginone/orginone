package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetimportstatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetimportstatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetimportstatusLogic {
	return GetimportstatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetimportstatusLogic) Getimportstatus(req types.GetImportStatusReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
