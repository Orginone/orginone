package logic

import (
	"context"
	"encoding/json"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asperson"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTenantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTenantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTenantsLogic {
	return &AddTenantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 租户导入
func (l *AddTenantsLogic) AddTenants(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	tenants := make([]*schema.AsTenant, 0)
	err := json.Unmarshal([]byte(in.Json), &tenants)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	ScCodes := make([]string, 0)
	tenantNames := make([]string, 0)
	for _, v := range tenants {
		tenantNames = append(tenantNames, v.TenantName)
		if !tools.IsNilOrEmpty(v.Edges.Unit.SocialCreditCode) {
			ScCodes = append(ScCodes, v.Edges.Unit.SocialCreditCode)
		}
	}
	tenantNames, err = tx.AsTenant.Query().Where(astenant.TenantNameIn(tenantNames...), astenant.IsDeleted(0)).Select(astenant.FieldTenantName).Strings(l.ctx)
	if len(ScCodes) > 0 {
		ScCodes, err = tx.AsUnit.Query().Where(asunit.SocialCreditCodeIn(ScCodes...), asunit.IsDeleted(0)).Select(asunit.FieldSocialCreditCode).Strings(l.ctx)
	}
	//查询普通租户类型
	tenantAttrId, err := l.svcCtx.SystemStore.GetTenantTypeId(l.ctx, 1)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	//查询单位管理员角色
	unitAdminAttrId, err := l.svcCtx.SystemStore.GetRoleId(l.ctx, 2)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	//查询默认租户类型
	defaultTenantAttrId, err := l.svcCtx.SystemStore.GetTenantTypeId(l.ctx, 3)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	for _, v := range tenants {
		if !tools.IsSocialCreditCode(v.Edges.Unit.SocialCreditCode) {
			continue
		}
		if tools.ArrIndexOf(tenantNames, v.TenantName) > 0 ||
			tools.ArrIndexOf(ScCodes, v.Edges.Unit.SocialCreditCode) > 0 {
			continue
		}
		//查询用户
		userEntity, err := tx.AsUser.Query().Where(asuser.IsDeleted(0), asuser.PhoneNumber(v.Edges.Unit.PhoneNumber), asuser.HasPersonWith(asperson.IsDeleted(0))).WithPerson().First(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if userEntity == nil {
			continue
		}
		//获取开发中心组织编号
		centerTenantEntity, err := l.svcCtx.SystemStore.AsTenant.Find().Where(astenant.TenantType(defaultTenantAttrId)).First(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}

		//开启事务
		tenantCode := tools.GenTenantCode()
		_, err = tx.AsUser.Update().Where(asuser.PhoneNumber(userEntity.PhoneNumber)).SetIsMaster(0).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("update user error. error message:" + err.Error())
		}
		//新增用户
		newUserEntity, err := tx.AsUser.Create().
			SetUserName(userEntity.UserName).
			SetPhoneNumber(userEntity.PhoneNumber).
			SetTenantCode(tenantCode).
			SetPwd(userEntity.Pwd).
			AddRoleIDs(unitAdminAttrId).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("create user error. error message:" + err.Error())
		}
		//新增人员
		newPersonEntity, err := tx.AsPerson.Create().
			SetIDCard(userEntity.Edges.Person.IDCard).
			SetNillableGender(&userEntity.Edges.Person.Gender).
			SetNillableUserEmail(&userEntity.Edges.Person.UserEmail).
			SetNillableUserPhoto(&userEntity.Edges.Person.UserPhoto).
			SetNillableProvince(&userEntity.Edges.Person.Province).
			SetNillableCity(&userEntity.Edges.Person.City).
			SetNillableStreetAddress(&userEntity.Edges.Person.StreetAddress).
			SetNillableUserCode(&userEntity.Edges.Person.UserCode).
			SetRealName(userEntity.Edges.Person.RealName).
			SetPhoneNumber(userEntity.Edges.Person.PhoneNumber).
			SetUserx(newUserEntity).
			SetTenantCode(tenantCode).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("create person user error. error message:" + err.Error())
		}
		//新建Tenant
		newTenantEntity, err := tx.AsTenant.Create().
			SetTenantCode(tenantCode).
			SetTenantName(v.TenantName).
			SetTenantType(tenantAttrId).
			SetCreateUser(newUserEntity.ID).
			SetUpdateUser(newUserEntity.ID).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("create tenant error. error message:" + err.Error())
		}
		//新建单位
		tx.AsUnit.Create().
			SetSocialCreditCode(v.Edges.Unit.SocialCreditCode).
			SetUnitName(newTenantEntity.TenantName).
			SetTenantx(newTenantEntity).
			SetLinkMan(newPersonEntity.RealName).
			SetLinkPhone(newPersonEntity.PhoneNumber).
			Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("create tenant error. error message:" + err.Error())
		}
		//删除默认租户人员和用户
		if userEntity.TenantCode == centerTenantEntity.TenantCode && userEntity.IsAdmin == 0 {
			_, err = tx.AsPerson.Delete().Where(asperson.UserID(userEntity.ID)).Exec(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
			}
			err = tx.AsUser.DeleteOneID(userEntity.ID).Exec(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
			}
		}
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
		}
	}
	return system.JsonResult(true, tx.Commit())
}
