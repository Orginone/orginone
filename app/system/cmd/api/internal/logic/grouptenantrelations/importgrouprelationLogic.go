package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportgrouprelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportgrouprelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportgrouprelationLogic {
	return ImportgrouprelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportgrouprelationLogic) Importgrouprelation(req types.ImportGroupRelationReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
