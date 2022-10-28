package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplenishUserCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplenishUserCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReplenishUserCodeLogic {
	return ReplenishUserCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplenishUserCodeLogic) ReplenishUserCode(req types.ReplenishUserCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
