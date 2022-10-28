package logic

import (
	"context"
	"errors"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserReRegLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserReRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserReRegLogic {
	return &UserReRegLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *UserReRegLogic) UserReReg(in *user.UserReRegReq) (*user.CommonRpcRes, error) {
	count, err := l.svcCtx.UserStore.AsUser.Query().Where(asuser.PhoneNumber(in.PhoneNumber)).Count(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if count > 0 {
		return &user.CommonRpcRes{}, errors.New("user is existed. please to login.")
	}
	roleId, err := l.svcCtx.UserStore.GetRoleId(l.ctx, 3)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	defaultTenantType, err := l.svcCtx.UserStore.GetTenantTypeId(l.ctx, 3)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	defaultTenant, err := l.svcCtx.UserStore.AsTenant.Query().Where(astenant.TenantType(defaultTenantType)).First(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	password, err := tools.BcryptEncryptPwd(in.Pwd)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	//开启事务
	tx, err := l.svcCtx.UserStore.Tx(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, errors.New("starting a transaction: " + err.Error())
	}
	userEntity, err := tx.AsUser.Create().
		SetUserName(in.RealName).SetPhoneNumber(in.PhoneNumber).
		SetTenantCode(defaultTenant.TenantCode).SetPwd(password).AddRoleIDs(roleId).Save(l.ctx)
	if err != nil {
		//回滚
		tx.Rollback()
		return &user.CommonRpcRes{}, errors.New("save user error. error message:" + err.Error())
	}
	_, err = tx.AsPerson.Create().
		SetRealName(in.RealName).SetPhoneNumber(in.PhoneNumber).
		SetUserID(userEntity.ID).SetTenantCode(defaultTenant.TenantCode).Save(l.ctx)
	if err != nil {
		//回滚
		tx.Rollback()
		return &user.CommonRpcRes{}, errors.New("save person error. error message:" + err.Error())
	}
	//提交事务
	err = tx.Commit()
	return user.Result("true", err)
}
