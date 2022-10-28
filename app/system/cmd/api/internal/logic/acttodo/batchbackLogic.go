package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) BatchbackLogic {
	return BatchbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchbackLogic) Batchback(req types.BatchbackReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
