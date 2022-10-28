package marketappalert

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCheckAlertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelCheckAlertLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelCheckAlertLogic {
	return CancelCheckAlertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelCheckAlertLogic) CancelCheckAlert(req types.CancelCheckAlertReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
