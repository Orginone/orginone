package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
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
func (l *JoinTeantByCodeLogic) JoinTeantByCode(in *system.TenantCodeReq) (*system.CommonRpcRes, error) {
	userCount, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.PhoneNumber(in.Account), asuser.TenantCode(in.TenantCode), asuser.TenantApplyingStateNEQ(3)).Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if userCount > 0 {
		return &system.CommonRpcRes{}, errors.New("has joined the tenant")
	}
	luserEntity, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.HasPersonWith(asperson.IsDeleted(0))).WithPerson().Where(asuser.ID(in.UserId)).Only(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, errors.New("starting a transaction: " + err.Error())
	}
	newUserEntity, err := tx.AsUser.Create().
		SetUserName(luserEntity.UserName).SetPhoneNumber(luserEntity.PhoneNumber).
		SetIsMaster(0).SetTenantApplyingState(1).SetStatus(1).
		SetTenantCode(in.TenantCode).SetPwd(luserEntity.Pwd).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsPerson.Create().SetUserx(newUserEntity).
		SetTenantCode(newUserEntity.TenantCode).
		SetRealName(luserEntity.Edges.Person.RealName).SetNillableCity(&luserEntity.Edges.Person.City).
		SetCreateUser(luserEntity.ID).SetNillableGender(&luserEntity.Edges.Person.Gender).
		SetUpdateUser(luserEntity.ID).SetNillableProvince(&luserEntity.Edges.Person.Province).
		SetIsMaster(luserEntity.Edges.Person.IsMaster).SetNillableUserCode(&luserEntity.Edges.Person.UserCode).
		SetNillableStreetAddress(&luserEntity.Edges.Person.StreetAddress).SetNillableIDCard(&luserEntity.Edges.Person.IDCard).
		SetNillableUserPhoto(&luserEntity.Edges.Person.UserPhoto).SetNillableUserEmail(&luserEntity.Edges.Person.UserEmail).
		SetPhoneNumber(luserEntity.Edges.Person.PhoneNumber).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return system.Result("true", err)
}
