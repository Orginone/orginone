package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortLazyTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSortLazyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) SortLazyTreeLogic {
	return SortLazyTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SortLazyTreeLogic) SortLazyTree(req types.SortLazyTreeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
