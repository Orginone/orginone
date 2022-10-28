package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asperson"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPersonsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonsLogic {
	return &AddPersonsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量新增组织人员
func (l *AddPersonsLogic) AddPersons(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	persons := make([]*schema.AsPerson, 0)
	err := json.Unmarshal([]byte(in.Json), &persons)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	var phones []string
	for _, person := range persons {
		phones = append(phones, person.PhoneNumber)
	}
	phones, err = l.svcCtx.SystemStore.AsPerson.Query().
		Where(asperson.PhoneNumberIn(phones...)).
		Select(asperson.FieldPhoneNumber).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	linq.From(persons).WhereT(func(i *schema.AsPerson) bool {
		for _, phone := range phones {
			if i.PhoneNumber == phone {
				return false
			}
		}
		return true
	}).ToSlice(&persons)
	if len(persons) <= 0 {
		return &system.CommonRpcRes{}, err
	}
	userCodePrefix, index, err := l.svcCtx.SystemStore.GetNewUserCodeInfo(l.ctx, persons[0].TenantCode)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	roleId, err := l.svcCtx.SystemStore.GetRoleId(l.ctx, 3)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	for i, person := range persons {
		isExist, err := tx.AsPerson.Query().Where(
			asperson.TenantCode(person.TenantCode),
			asperson.PhoneNumber(person.PhoneNumber),
			asperson.IsDeleted(0)).
			Exist(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, err
		}
		if !isExist {
			pwd, err := tools.BcryptEncryptPwd(person.PhoneNumber[5:])
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			userEntity, err := tx.AsUser.Create().
				SetUserName(person.RealName).
				SetPhoneNumber(person.PhoneNumber).
				SetStatus(3).
				SetIsCreated(3).
				SetTenantApplyingState(2).
				SetTenantCode(person.TenantCode).
				SetPwd(pwd).
				SetIsMaster(1).
				AddRoleIDs(roleId).
				Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			person.Status = 3
			person.UserCode = fmt.Sprintf("%s%06d", userCodePrefix, index+int64(i))
			personCreater := tx.AsPerson.Create()
			tools.SchemaUpdateAny(personCreater.Mutation(), person, "id")
			_, err = personCreater.SetUserx(userEntity).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
		}
	}
	return system.Result("true", tx.Commit())
}
