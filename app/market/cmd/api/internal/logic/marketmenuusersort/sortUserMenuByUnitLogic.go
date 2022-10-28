package marketmenuusersort

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortUserMenuByUnitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSortUserMenuByUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SortUserMenuByUnitLogic {
	return SortUserMenuByUnitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SortUserMenuByUnitLogic) SortUserMenuByUnit(req types.SortUserMenuByUnitReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
