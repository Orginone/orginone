package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListPageLogic {
	return ListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPageLogic) ListPage(req types.ListPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
