package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVOLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVOLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListVOLogic {
	return ListVOLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVOLogic) ListVO(req types.ListVOReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
