package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelUniAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelUniAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelUniAppLogic {
	return CancelUniAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelUniAppLogic) CancelUniApp(req types.CancelUniAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
