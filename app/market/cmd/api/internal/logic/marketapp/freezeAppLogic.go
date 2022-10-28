package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreezeAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreezeAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) FreezeAppLogic {
	return FreezeAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreezeAppLogic) FreezeApp(req types.FreezeAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
