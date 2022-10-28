package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemainAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRemainAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRemainAppLogic {
	return GetRemainAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRemainAppLogic) GetRemainApp(req types.GetRemainAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
