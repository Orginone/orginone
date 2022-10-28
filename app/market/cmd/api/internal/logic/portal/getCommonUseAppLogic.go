package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommonUseAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommonUseAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCommonUseAppLogic {
	return GetCommonUseAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommonUseAppLogic) GetCommonUseApp(req types.GetCommonUseAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
