package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListAppListLogic {
	return ListAppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAppListLogic) ListAppList(req types.ListAppListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
