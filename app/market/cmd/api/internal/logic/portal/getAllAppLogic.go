package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllAppLogic {
	return GetAllAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllAppLogic) GetAllApp(req types.GetAllAppReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
