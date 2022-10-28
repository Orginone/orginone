package token

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistrationretrieveidCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistrationretrieveidCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegistrationretrieveidCardLogic {
	return RegistrationretrieveidCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistrationretrieveidCardLogic) RegistrationretrieveidCard(req types.RegistrationRetrieveIdCardReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
