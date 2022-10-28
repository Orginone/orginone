package logic

import (
	"context"
	"errors"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPassWordLogic {
	return &ResetPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置用户密码
func (l *ResetPassWordLogic) ResetPassWord(in *user.CommonRpcRes) (*user.CommonRpcRes, error) {
	if len(in.Json) != 11 {
		return &user.CommonRpcRes{}, errors.New("phonenumber invalid !")
	}
	newPwd, err := tools.BcryptEncryptPwd(in.Json[5:])
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	count, err := l.svcCtx.UserStore.AsUser.Update().Where(asuser.PhoneNumber(in.Json)).SetPwd(newPwd).Save(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if count <= 0 {
		return &user.CommonRpcRes{}, errors.New("update fail !")
	}
	return user.JsonResult(true, nil)
}
