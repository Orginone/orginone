package marketmenuusersort

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenuSortListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenuSortListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserMenuSortListLogic {
	return GetUserMenuSortListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenuSortListLogic) GetUserMenuSortList(req types.GetUserMenuSortListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
