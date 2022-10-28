package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetpasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPassword1Logic(ctx context.Context, svcCtx *svc.ServiceContext) ResetpasswordLogic {
	return ResetpasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetpasswordLogic) ResetPassword1(req types.ResetPasswordReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
