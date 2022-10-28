package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUniAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUniAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUniAppLogic {
	return AddUniAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUniAppLogic) AddUniApp(req types.AddUniAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
