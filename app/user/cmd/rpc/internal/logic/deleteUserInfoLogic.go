package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserInfoLogic {
	return &DeleteUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteUserInfoLogic) DeleteUserInfo(in *user.DeleteUserInfoReq) (*user.CommonRpcRes, error) {
	tx, err := l.svcCtx.UserStore.Tx(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	users, err := l.svcCtx.UserStore.AsUser.Query().Where(asuser.And(asuser.IDIn(in.UserIds...), asuser.IsCreated(3),
		asuser.Not(asuser.HasWorkingDatasWith(asworkingdata.IsDeleted(0), asworkingdata.Type(1))))).WithPerson().All(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	if len(users) > 0 {
		for _, i := range users {
			//删除用户
			_, err = i.Update().SetIsDeleted(1).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &user.CommonRpcRes{}, err
			}
			//删除人员
			_, err = i.Edges.Person.Update().ClearAgencys().ClearJobs().SetIsDeleted(1).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &user.CommonRpcRes{}, err
			}
			samePhoneUsers, err := tx.AsUser.Query().Where(asuser.PhoneNumber(i.PhoneNumber), asuser.IsDeleted(0)).Order(schema.Desc(asuser.FieldCreateTime)).All(l.ctx)
			if err != nil {
				tx.Rollback()
				return &user.CommonRpcRes{}, err
			}
			if len(samePhoneUsers) > 0 && i.IsMaster == 1 {
				_, err = tx.AsUser.UpdateOne(samePhoneUsers[0]).SetIsMaster(1).Save(l.ctx)
				if err != nil {
					tx.Rollback()
					return &user.CommonRpcRes{}, err
				}
			} else {
				pwd, err := tools.BcryptEncryptPwd(i.PhoneNumber[5:])
				userEntity, err := tx.AsUser.Create().
					SetUserName(i.UserName).
					SetPhoneNumber(i.PhoneNumber).
					SetStatus(2).
					SetIsAdmin(0).
					SetIsCreated(3).
					SetTenantApplyingState(0).
					SetTenantCode("ZH200205214156yVEQnZ").
					SetPwd(pwd).
					SetIsMaster(1).
					Save(l.ctx)
				if err != nil {
					tx.Rollback()
					return &user.CommonRpcRes{}, err
				}
				userCode, _ := l.svcCtx.UserStore.CreateUserCode(l.ctx, "ZH200205214156yVEQnZ")
				_, err = tx.AsPerson.Create().
					SetPhoneNumber(i.PhoneNumber).
					SetRealName(i.UserName).
					SetTenantCode("ZH200205214156yVEQnZ").
					SetUserCode(userCode).
					SetUserx(userEntity).
					SetStatus(2).
					Save(l.ctx)
				if err != nil {
					tx.Rollback()
					return &user.CommonRpcRes{}, err
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	return &user.CommonRpcRes{}, nil
}
