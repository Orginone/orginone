package marketappalert

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyAlertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyAlertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MyAlertListLogic {
	return MyAlertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyAlertListLogic) MyAlertList(req types.MyAlertListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
