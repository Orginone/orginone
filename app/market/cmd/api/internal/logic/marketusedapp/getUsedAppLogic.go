package marketusedapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsedAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsedAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUsedAppLogic {
	return GetUsedAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsedAppLogic) GetUsedApp(req types.GetUsedAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
