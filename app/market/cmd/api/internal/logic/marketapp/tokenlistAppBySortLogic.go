package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenlistAppBySortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenlistAppBySortLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenlistAppBySortLogic {
	return TokenlistAppBySortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenlistAppBySortLogic) TokenlistAppBySort(req types.TokenListAppBySortReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
