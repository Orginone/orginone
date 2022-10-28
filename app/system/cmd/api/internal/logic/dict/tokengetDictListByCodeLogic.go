package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokengetDictListByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokengetDictListByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokengetDictListByCodeLogic {
	return TokengetDictListByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokengetDictListByCodeLogic) TokengetDictListByCode(req types.DicTokenGetDictListByCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
