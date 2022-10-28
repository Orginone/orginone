package token

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserupdatepwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserupdatepwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserupdatepwdLogic {
	return UserupdatepwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserupdatepwdLogic) Userupdatepwd(req types.UserUpdatePwdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
