package menus

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoutesLogic {
	return RoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoutesLogic) Routes(req types.RoutesReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
