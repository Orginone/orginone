package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokendetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokendetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokendetailLogic {
	return TokendetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokendetailLogic) Tokendetail(req types.TokenDetailReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
