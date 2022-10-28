package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchpassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchpassLogic(ctx context.Context, svcCtx *svc.ServiceContext) BatchpassLogic {
	return BatchpassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchpassLogic) Batchpass(req types.BatchpassReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
