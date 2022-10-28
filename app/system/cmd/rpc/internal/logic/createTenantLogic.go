package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建租户
func (l *CreateTenantLogic) CreateTenant(in *system.CreateTenantReq) (*system.CommonRpcRes, error) {
	if !tools.IsSocialCreditCode(in.SocialCreCode) {
		return &system.CommonRpcRes{}, errors.New("socialCreditCode is error!")
	}
	//查询租户是否存在
	count, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantName(in.TenantName)).Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if count > 0 {
		return &system.CommonRpcRes{}, errors.New("tenantName is existed. please confirm tenantName.")
	}
	//查询单位是否存在
	count, err = l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.SocialCreditCode(in.SocialCreCode)).Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if count > 0 {
		return &system.CommonRpcRes{}, errors.New("socialCreditCode is existed. please confirm socialCreditCode.")
	}
	//查询用户
	userEntity, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.ID(in.UserId), asuser.HasPersonWith(asperson.IsDeleted(0))).WithPerson().First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
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
	//获取开发中心组织编号
	centerTenantEntity, err := l.svcCtx.SystemStore.AsTenant.Find().Where(astenant.TenantType(defaultTenantAttrId)).First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}

	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
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
		SetTenantName(in.TenantName).
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
		SetSocialCreditCode(in.SocialCreCode).
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
		_, err = tx.AsPerson.Delete().Where(asperson.UserID(in.UserId)).Exec(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
		}
		err = tx.AsUser.DeleteOneID(in.UserId).Exec(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
		}
	}
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, errors.New("delete default user error. error message:" + err.Error())
	}
	//提交事务
	err = tx.Commit()
	return system.Result(newPersonEntity.TenantCode, err)
}
