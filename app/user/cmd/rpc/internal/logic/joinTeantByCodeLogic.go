package logic

import (
	"context"
	"errors"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinTeantByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinTeantByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinTeantByCodeLogic {
	return &JoinTeantByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 加入租户
func (l *JoinTeantByCodeLogic) JoinTeantByCode(in *user.TenantCodeReq) (*user.CommonRpcRes, error) {
	userCount, err := l.svcCtx.UserStore.AsUser.Query().Where(asuser.PhoneNumber(in.Account), asuser.TenantCode(in.TenantCode), asuser.TenantApplyingStateNEQ(3)).Count(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if userCount > 0 {
		return &user.CommonRpcRes{}, errors.New("请勿重复加入")
	}
	dbUserEntity, err := l.svcCtx.UserStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.ID(in.UserId), asuser.HasPersonWith(asperson.IsDeleted(0))).WithPerson().Only(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	//开启事务
	tx, err := l.svcCtx.UserStore.Tx(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, errors.New("starting a transaction: " + err.Error())
	}
	newUserEntity, err := tx.AsUser.Create().
		SetUserName(dbUserEntity.UserName).SetPhoneNumber(dbUserEntity.PhoneNumber).
		SetIsMaster(0).SetTenantApplyingState(1).SetStatus(1).
		SetTenantCode(in.TenantCode).SetPwd(dbUserEntity.Pwd).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &user.CommonRpcRes{}, err
	}
	_, err = tx.AsPerson.Create().SetUserx(newUserEntity).
		SetTenantCode(newUserEntity.TenantCode).
		SetRealName(dbUserEntity.Edges.Person.RealName).SetNillableCity(&dbUserEntity.Edges.Person.City).
		SetCreateUser(dbUserEntity.ID).SetNillableGender(&dbUserEntity.Edges.Person.Gender).
		SetUpdateUser(dbUserEntity.ID).SetNillableProvince(&dbUserEntity.Edges.Person.Province).
		SetIsMaster(dbUserEntity.Edges.Person.IsMaster).SetNillableUserCode(&dbUserEntity.Edges.Person.UserCode).
		SetNillableStreetAddress(&dbUserEntity.Edges.Person.StreetAddress).SetNillableIDCard(&dbUserEntity.Edges.Person.IDCard).
		SetNillableUserPhoto(&dbUserEntity.Edges.Person.UserPhoto).SetNillableUserEmail(&dbUserEntity.Edges.Person.UserEmail).
		SetPhoneNumber(dbUserEntity.Edges.Person.PhoneNumber).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &user.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return user.Result("true", err)
}
