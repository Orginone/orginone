package marketappapply

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMyListLogic {
	return GetMyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyListLogic) GetMyList(req types.GetMyListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
