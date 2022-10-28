package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthoritytransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthoritytransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthoritytransferLogic {
	return AuthoritytransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthoritytransferLogic) Authoritytransfer(req types.AuthorityTransferReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
