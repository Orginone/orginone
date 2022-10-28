package marketappalert

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAlertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAlertLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendAlertLogic {
	return SendAlertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAlertLogic) SendAlert(req types.SendAlertReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
