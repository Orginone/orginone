package marketmenuusersort

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortUserMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSortUserMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) SortUserMenuLogic {
	return SortUserMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SortUserMenuLogic) SortUserMenu(req types.SortUserMenuReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
