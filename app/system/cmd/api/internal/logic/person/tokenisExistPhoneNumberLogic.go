package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenisExistPhoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenisExistPhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenisExistPhoneNumberLogic {
	return TokenisExistPhoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenisExistPhoneNumberLogic) TokenisExistPhoneNumber(req types.TokenIsExistPhoneNumberReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
