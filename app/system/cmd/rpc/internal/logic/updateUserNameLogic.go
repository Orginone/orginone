package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asunit"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserNameLogic {
	return &UpdateUserNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户名
func (l *UpdateUserNameLogic) UpdateUserName(in *system.UpdateUserName) (*system.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return system.Result("false", err)
	}
	userEntity, err := tx.AsUser.Get(l.ctx, in.UserId)
	if err != nil {
		return system.Result("false", err)
	}
	_, err = tx.AsUser.Update().Where(asuser.PhoneNumber(userEntity.PhoneNumber)).SetUserName(in.UserName).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	updatePerson := tx.AsPerson.Update().Where(asperson.PhoneNumber(userEntity.PhoneNumber)).SetRealName(in.UserName)
	if !tools.IsNilOrEmpty(in.IdCard) {
		updatePerson = updatePerson.SetIDCard(in.IdCard)
	}
	if !tools.IsNilOrEmpty(in.UserCode) {
		updatePerson = updatePerson.SetUserCode(in.UserCode)
	}
	_, err = updatePerson.Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	_, err = tx.AsUnit.Update().Where(asunit.LinkPhone(userEntity.PhoneNumber)).SetLinkMan(in.UserName).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	err = tx.Commit()
	return system.Result("true", err)
}
