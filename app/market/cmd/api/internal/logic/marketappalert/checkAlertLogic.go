package marketappalert

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAlertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckAlertLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckAlertLogic {
	return CheckAlertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckAlertLogic) CheckAlert(req types.CheckAlertReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
