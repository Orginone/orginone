package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ResetPasswordLogic {
	return ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req types.ResetPasswordReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.UserRpc.ResetPassWord(l.ctx, &user.CommonRpcRes{Json: req.PhoneNumber}))
}
