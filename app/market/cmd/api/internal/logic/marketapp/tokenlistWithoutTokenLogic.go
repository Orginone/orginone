package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenlistWithoutTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenlistWithoutTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenlistWithoutTokenLogic {
	return TokenlistWithoutTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenlistWithoutTokenLogic) TokenlistWithoutToken(req types.TokenListWithoutTokenReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
