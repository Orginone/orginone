package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetlistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetlistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetlistsLogic {
	return GetlistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetlistsLogic) Getlists(req types.GetListsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
