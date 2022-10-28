package logic

import (
	"context"
	"errors"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswdLogic {
	return &UpdatePasswdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改密码
func (l *UpdatePasswdLogic) UpdatePasswd(in *user.UpdatePasswdReq) (*user.CommonRpcRes, error) {
	userEntity, err := l.svcCtx.UserStore.AsUser.Get(l.ctx, in.UserId)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if tools.ValidateBcryptPwd(userEntity.Pwd, in.NewPwd) {
		return user.Result("true", nil)
	}
	if !tools.ValidateBcryptPwd(userEntity.Pwd, in.OldPwd) {
		return &user.CommonRpcRes{}, errors.New("password is error.")
	}
	pwd, err := tools.BcryptEncryptPwd(in.NewPwd)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	_, err = userEntity.Update().SetPwd(pwd).Save(l.ctx)
	return user.Result("true", err)
}
