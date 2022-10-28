package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferauthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferauthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) TransferauthorityLogic {
	return TransferauthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferauthorityLogic) Transferauthority(req types.TransferAuthorityReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
