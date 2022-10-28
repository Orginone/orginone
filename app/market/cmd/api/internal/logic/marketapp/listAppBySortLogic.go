package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAppBySortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAppBySortLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListAppBySortLogic {
	return ListAppBySortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAppBySortLogic) ListAppBySort(req types.ListAppBySortReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
