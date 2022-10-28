package logic

import (
	"context"
	"errors"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/schema/asuser"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByAccountLogic {
	return &FindUserByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *FindUserByAccountLogic) FindUserByAccount(in *user.UserByAccountReq) (*user.CommonRpcRes, error) {
	query := l.svcCtx.UserStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.Status(2))
	if in.Id > 0 {
		query = query.Where(asuser.ID(in.Id))
	} else {
		query = query.Where(asuser.PhoneNumber(in.Account))
	}
	userEntitys, err := query.All(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if len(userEntitys) < 1 {
		return &user.CommonRpcRes{}, errors.New("user not found!")
	}
	if !tools.IsNilOrEmpty(in.Password) {
		//验证密码
		linq.From(userEntitys).WhereT(func(u *schema.AsUser) bool {
			return tools.ValidateBcryptPwd(u.Pwd, in.Password)
		}).ToSlice(&userEntitys)
	}
	if len(userEntitys) < 1 {
		err = errors.New("password incorrect!")
	}
	return user.JsonResult(userEntitys, err)
}
