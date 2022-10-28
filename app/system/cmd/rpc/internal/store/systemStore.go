package store

import (
	"context"
	"errors"
	"fmt"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/models"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asrole"
	"orginone/common/schema/astenant"
	"orginone/common/schema/astenantattr"
	"orginone/common/schema/asunit"
	"orginone/common/schema/asuser"
	"orginone/common/tools"
	"orginone/common/tools/linq"
	"strconv"
	"strings"
)

type SystemStore struct {
	*common.DbStore
}

func NewSystemStore(c models.DbStoreConfig) *SystemStore {
	return &SystemStore{
		DbStore: common.NewDbStore(c),
	}
}

// 查询人员详细信息
func (s *SystemStore) QueryPersonInfo(ctx context.Context, userIds ...int64) ([]*model.PersonInfo, error) {
	persons := make([]*model.PersonInfo, 0)
	userEntitys, err := s.AsUser.Query().Where(asuser.IDIn(userIds...)).
		Where(asuser.HasRolesWith(asrole.IsDeleted(0)), asuser.HasPerson()).
		WithPerson(func(apq *schema.AsPersonQuery) {
			apq.WithAgencys(func(aiaq *schema.AsInnerAgencyQuery) { aiaq.Where(asinneragency.IsDeleted(0)) })
		}).WithRoles().All(ctx)
	if err != nil {
		return persons, err
	}
	linq.From(userEntitys).SelectT(func(u *schema.AsUser) *model.PersonInfo {
		//角色
		var roleIds []string
		var roleNames []string
		for _, v := range u.Edges.Roles {
			roleId := strconv.Itoa(int(v.ID))
			if tools.ArrIndexOf(roleIds, roleId) == -1 {
				roleIds = append(roleIds, roleId)
				roleNames = append(roleNames, v.RoleName)
			}
		}
		//部门
		var deptIds []int64
		deptList := make([]*model.InnerAgencyTree, 0)
		for _, a := range u.Edges.Person.Edges.Agencys {
			if tools.ArrIndexOf(deptIds, a.ID) == -1 {
				agency := new(model.InnerAgencyTree)
				agency.AsInnerAgency = a
				agency.Children = make([]*model.InnerAgencyTree, 0)
				deptList = append(deptList, agency)
				deptIds = append(deptIds, a.ID)
			}
		}
		//当前逻辑：校验该租户是否以任何形式购买
		ifPurchase := false
		c, e := s.AsMarketApp.Query().Where(asmarketapp.IsDeleted(0), asmarketapp.Or(
			asmarketapp.HasAppPurchasesWith(asmarketapppurchase.TenantID(u.TenantCode), asmarketapppurchase.IsDeleted(0)),
			asmarketapp.HasAppGroupDistribsWith(asmarketappgroupdistribution.TenantID(u.TenantCode), asmarketappgroupdistribution.IsDeleted(0)))).Count(ctx)
		if e == nil && c > 0 {
			ifPurchase = true
		}
		return &model.PersonInfo{
			DeptList:   deptList,
			IsCreated:  u.IsCreated,
			IfPurchase: ifPurchase,
			RoleId:     strings.Join(roleIds, ","),
			RoleName:   strings.Join(roleNames, ","),
			AsPerson:   u.Edges.Person,
			RoleList:   u.Edges.Roles,
		}
	}).ToSlice(&persons)
	return persons, err
}

//根据集团Id查询单位类型的节点信息
func (s *SystemStore) FindUnitNodeInfoById(ctx context.Context, managerTenantCode string, groupId int64) ([]*model.GroupUnitInfo, error) {
	groupInfos := make([]*model.GroupUnitInfo, 0)
	gtrEntitys, err := s.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ParentID(groupId),
		asgrouptenantrelations.StatusIn(102, 106), asgrouptenantrelations.IsHide(0), asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.ExpiresTimeIsNil(),
		asgrouptenantrelations.Type(2), asgrouptenantrelations.HasTenant()).
		WithTenant(func(atq *schema.AsTenantQuery) {
			atq.Where(astenant.HasUnit()).WithUnit()
		}).All(ctx)
	if err != nil {
		return groupInfos, err
	}
	for _, gtr := range gtrEntitys {
		groupInfos = append(groupInfos, model.NewGroupUnitInfo(2, false).SetParentId(groupId).SetGroupRelations(gtr).SetTenant(gtr.Edges.Tenant).SetManagerCode(managerTenantCode))
	}
	return groupInfos, err
}

//根据集团Id查询集团类型的节点信息
func (s *SystemStore) FindGroupNodeInfoById(ctx context.Context, managerTenantCode string, groupId int64) ([]*model.GroupUnitInfo, error) {
	groupInfos := make([]*model.GroupUnitInfo, 0)
	groupIds, err := s.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ParentID(groupId),
		asgrouptenantrelations.StatusIn(102, 106), asgrouptenantrelations.IsHide(0), asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.ExpiresTimeIsNil(),
		asgrouptenantrelations.Type(1), asgrouptenantrelations.HasGroupWith(asallgroup.IsDeleted(0))).Select(asgrouptenantrelations.FieldSonID).Int64s(ctx)
	if err != nil {
		return groupInfos, err
	}
	if len(groupIds) > 0 {
		groupEntitys, err := s.AsAllGroup.Query().Where(asallgroup.IDIn(groupIds...)).All(ctx)
		if err != nil {
			return groupInfos, err
		}
		tenantCodes := make([]string, 0)
		linq.From(groupEntitys).SelectT(func(g *schema.AsAllGroup) string { return g.TenantCode }).ToSlice(&tenantCodes)
		tenantEntitys, err := s.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(tenantCodes...), astenant.HasUnit()).WithUnit().All(ctx)
		if err != nil {
			return groupInfos, err
		}
		for _, g := range groupEntitys {
			below, _ := s.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ParentID(g.ID), asgrouptenantrelations.StatusIn(102, 106), asgrouptenantrelations.IsHide(0),
				asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.ExpiresTimeIsNil()).Exist(ctx)
			for _, t := range tenantEntitys {
				if g.TenantCode == t.TenantCode {
					groupInfos = append(groupInfos, model.NewGroupUnitInfo(1, below).SetParentId(groupId).SetTenant(t).SetGroup(g).SetManagerCode(managerTenantCode))
				}
			}
		}
	}
	return groupInfos, err
}

//获取机构编号根据租户编号
func (s *SystemStore) GetAgencyInnerCode(ctx context.Context, tenantCode string) (string, error) {
	var agencyCode string
	unitEntity, err := s.AsUnit.Query().Where(asunit.IsDeleted(0), asunit.HasTenantxWith(astenant.TenantCode(tenantCode))).First(ctx)
	if err != nil {
		return agencyCode, err
	}
	if unitEntity == nil {
		return agencyCode, errors.New("tenant not found.")
	}
	if (tools.IsNilOrEmpty(unitEntity.UnitCode) || len(unitEntity.UnitCode) != 9) && len(unitEntity.SocialCreditCode) >= 17 {
		s.AsUnit.Update().SetUnitCode(unitEntity.SocialCreditCode[8:17])
		agencyCode = unitEntity.SocialCreditCode[8:11]
	} else if len(unitEntity.UnitCode) >= 3 {
		agencyCode = unitEntity.UnitCode[0:3]
	} else {
		return agencyCode, errors.New("unit code invalid.")
	}
	innerAgencyEntitys, err := s.AsInnerAgency.Query().Where(asinneragency.TenantCode(tenantCode), asinneragency.IsDeleted(0), asinneragency.AgencyCodeContains(agencyCode)).All(ctx)
	if innerAgencyEntitys == nil || len(innerAgencyEntitys) == 0 {
		return agencyCode + "0001", nil
	}
	value := linq.From(innerAgencyEntitys).SelectT(func(e *schema.AsInnerAgency) int64 {
		num, err := strconv.ParseInt(e.AgencyCode[3:7], 10, 64)
		if err != nil {
			return 0
		}
		return num
	}).Max()
	return fmt.Sprintf("%s%04d", agencyCode, value.(int64)+1), nil
}

// 查询全新人员编号
func (s *SystemStore) GetNewUserCodeInfo(ctx context.Context, tenantCode string) (code string, index int64, err error) {
	code = "DEF"
	unitEntity, err := s.AsUnit.Query().Where(asunit.IsDeleted(0), asunit.HasTenantxWith(astenant.TenantCode(tenantCode))).First(ctx)
	if err == nil && unitEntity != nil {
		if tools.IsNilOrEmpty(unitEntity.UnitCode) {
			if len(unitEntity.SocialCreditCode) >= 17 {
				unitEntity.UnitCode = unitEntity.SocialCreditCode[8:17]
				unitEntity.Update().SetUnitCode(unitEntity.UnitCode).Save(ctx)
			}
		}
		if len(unitEntity.UnitCode) >= 3 {
			code = unitEntity.UnitCode[0:3]
			index, err = s.AsPerson.Query().Where(asperson.TenantCode(tenantCode), asperson.UserCodeHasPrefix(code)).Count(ctx)
		}
	}
	index++
	return
}

//获取租户类型Id
//tenantType 1:普通租户,2:平台租户,3:默认租户,4:集团租户,5:开发租户
func (s *SystemStore) GetTenantTypeId(ctx context.Context, tenantType int64) (int64, error) {

	tenantTypeName := "普通租户"
	switch tenantType {
	case 1:
		tenantTypeName = "普通租户"
	case 2:
		tenantTypeName = "平台租户"
	case 3:
		tenantTypeName = "默认租户"
	case 4:
		tenantTypeName = "集团租户"
	case 5:
		tenantTypeName = "开发租户"
	default:
		return 0, errors.New("not found.")
	}
	attrEntity, err := s.AsTenantAttr.Query().Where(astenantattr.AttrName(tenantTypeName)).First(ctx)
	if err != nil {
		return 0, err
	}
	return attrEntity.ID, nil
}

//获取租户类型Id
//roleType 1:平台超级管理员,2:单位管理员,3:普通用户,4:应用管理员,5:集团管理员
func (s *SystemStore) GetRoleId(ctx context.Context, roleType int64) (int64, error) {
	roleName := "平台超级管理员"
	switch roleType {
	case 1:
		roleName = "平台超级管理员"
	case 2:
		roleName = "单位管理员"
	case 3:
		roleName = "普通用户"
	case 4:
		roleName = "应用管理员"
	case 5:
		roleName = "集团管理员"
	default:
		return 0, errors.New("not found.")
	}
	roleEntity, err := s.AsRole.Query().Where(asrole.RoleName(roleName)).First(ctx)
	if err != nil {
		return 0, err
	}
	return roleEntity.ID, nil
}

//根据租户编号获取单位名称
func (s *SystemStore) GetUnitNameByTenantCode(ctx context.Context, tenantCode string) *schema.AsUnit {
	unitEntitys, err := s.AsTenant.Query().Where(astenant.TenantCode(tenantCode)).QueryUnit().Limit(1).All(ctx)
	if err == nil && len(unitEntitys) > 0 {
		return unitEntitys[0]
	}
	return new(schema.AsUnit)
}

// 集团加入集团
func (s *SystemStore) GroupJoinGroup(ctx context.Context, groupEntity *schema.AsAllGroup, parentGroup *schema.AsAllGroup, tx *schema.Tx) error {
	serial, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ParentID(parentGroup.ID)).Count(ctx)
	if err != nil {
		return err
	}
	serial = serial + 1
	groupCode := parentGroup.GroupCode + fmt.Sprintf("%03d", serial)
	_, err = tx.AsGroupTenantRelations.Create().SetParentID(parentGroup.ID).SetSonID(groupEntity.ID).SetType(1).
		SetStatus(102).SetGroupCode(groupCode).SetSerial(serial).SetSort(serial).Save(ctx)
	if err != nil {
		return err
	}
	_, err = groupEntity.Update().SetGroupCode(groupCode).Save(ctx)
	groupEntity.GroupCode = groupCode
	if err != nil {
		return err
	}
	oldRelations, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.ExpiresTimeIsNil(), asgrouptenantrelations.ParentID(groupEntity.ID)).All(ctx)
	if err != nil {
		return err
	}
	for serial, r := range oldRelations {
		serial = serial + 1
		_, err = r.Update().SetSerial(int64(serial)).SetSort(int64(serial)).SetIsDeleted(1).SetGroupCode(groupCode + fmt.Sprintf("%03d", serial)).Save(ctx)
		if err != nil {
			return err
		}
	}
	return err
}

// 多租户加入集团
func (s *SystemStore) TenantsJoinGroup(ctx context.Context, tenantIds []int64, groupId int64, flag int64) error {
	tx, err := s.Tx(ctx)
	if err != nil {
		return err
	}
	for _, tenantId := range tenantIds {
		if flag == 101 {
			_, err = tx.AsGroupTenantRelations.Create().SetParentID(groupId).
				SetSonID(tenantId).SetStatus(flag).SetType(2).SetSerial(0).Save(ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			_, err = tx.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.ParentID(groupId),
				asgrouptenantrelations.SonID(tenantId), asgrouptenantrelations.Type(2),
				asgrouptenantrelations.Status(101)).SetStatus(flag).Save(ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Commit()
	return err
}

// 创建集团并设置租户为集团租户
func (s *SystemStore) CreateAndUpdateTenant(ctx context.Context, in *system.CreateGroupReq, tx *schema.Tx) (*schema.AsAllGroup, error) {
	groupTenantId, err := s.GetTenantTypeId(ctx, 4)
	if err != nil {
		return nil, err
	}
	adminTenantId, err := s.GetTenantTypeId(ctx, 2)
	if err != nil {
		return nil, err
	}
	_, err = tx.AsTenant.Update().Where(astenant.TenantCode(in.TenantCode), astenant.TenantTypeNEQ(adminTenantId)).SetTenantType(groupTenantId).Save(ctx)
	if err != nil {
		return nil, err
	}
	var groupType int64 = 1
	if tools.IsNilOrEmpty(in.GroupCode) {
		groupType = 2
	}
	groupEntity, err := tx.AsAllGroup.Create().SetGroupName(in.GroupName).SetGroupDescription(in.GroupDescription).
		SetTenantCode(in.TenantCode).SetNillableGroupCode(&in.GroupCode).SetType(groupType).SetDepth(int64(len(in.Shape))).Save(ctx)
	if err != nil {
		return nil, err
	}
	for index, layer := range in.Shape {
		_, err = tx.AsLayer.Create().SetGroup(groupEntity).SetLayer(int64(index)).SetWidth(layer).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	groupRoleId, err := s.GetRoleId(ctx, 5)
	if err != nil {
		return nil, err
	}
	adminRoleId, err := s.GetRoleId(ctx, 1)
	if err != nil {
		return nil, err
	}
	if exist, err := tx.AsUser.Query().Where(asuser.ID(in.UserId), asuser.HasRolesWith(asrole.IDIn(groupRoleId, adminRoleId))).Exist(ctx); err == nil && !exist {
		_, err = tx.AsUser.UpdateOneID(in.UserId).AddRoleIDs(groupRoleId).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	return groupEntity, err
}

// 创建租户组与集团的关系
func (s *SystemStore) CreateGroupTenantRelation(ctx context.Context, groupId int64, tenantIds []int64, flag int64) error {
	groupEntity, err := s.AsAllGroup.Get(ctx, groupId)
	if err != nil {
		return err
	}
	serial, err := groupEntity.QueryAllTenants().Where(asgrouptenantrelations.GroupCodeNotNil()).Count(ctx)
	if err != nil {
		return err
	}
	manageTenant, err := s.AsTenant.Query().Where(astenant.TenantCode(groupEntity.TenantCode)).First(ctx)
	if err != nil {
		return err
	}
	if len(tenantIds) == 0 {
		tenantIds = append(tenantIds, manageTenant.ID)
	}
	for _, tenantId := range tenantIds {
		if flag == 102 && manageTenant.ID == tenantId {
			flag = 106
		}
		serial = serial + 1
		groupCode := groupEntity.GroupCode + fmt.Sprintf("%03d", serial)
		num, err := s.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.Status(101),
			asgrouptenantrelations.Type(2), asgrouptenantrelations.ParentID(groupId), asgrouptenantrelations.SonID(tenantId),
			asgrouptenantrelations.ExpiresTimeIsNil()).SetSerial(serial).SetSort(serial).SetGroupCode(groupCode).SetStatus(flag).Save(ctx)
		if err != nil {
			return err
		}
		if num < 1 { // 不存在，则为集团分配
			_, err = s.AsGroupTenantRelations.Create().SetParentID(groupId).SetSonID(tenantId).
				SetSerial(serial).SetSort(serial).SetGroupCode(groupCode).SetStatus(flag).SetType(2).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return err
}

//根据账户查询账户租户
func (s *SystemStore) GetTenantByAccountAndCount(ctx context.Context, account string, count int64) ([]string, error) {
	query := s.AsUser.Query().Where(asuser.PhoneNumber(account), asuser.IsDeleted(0))
	switch count {
	case 4:
		query = query.Where(asuser.TenantApplyingStateIn(1, 2, 3))
		break
	case 5:
		query = query.Where(asuser.TenantApplyingStateIn(0, 2))
		break
	case 6:
		break
	case 7:
		query = query.Where(asuser.TenantApplyingStateIn(2, 3))
		break
	default:
		query = query.Where(asuser.TenantApplyingState(count))
		break
	}
	return query.Select(asuser.FieldTenantCode).Strings(ctx)
}

//集团发起租户加入集团中任意集团类型节点
func (s *SystemStore) GroupPullTenants(ctx context.Context, groupId int64, sourceGroupId int64, tenantIds []int64) (*system.CommonRpcRes, error) {
	sourceGroupEntity, err := s.AsAllGroup.Get(ctx, sourceGroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = s.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.GroupCodeHasPrefix(sourceGroupEntity.GroupCode),
		asgrouptenantrelations.Type(2), asgrouptenantrelations.SonIDIn(tenantIds...)).SetIsDeleted(1).Save(ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.Result("true", s.CreateGroupTenantRelation(ctx, groupId, tenantIds, 102))
}

//集团分配租户单位
func (s *SystemStore) GroupResetTenants(ctx context.Context, groupId int64, sourceGroupId int64, tenantIds []int64) (*system.CommonRpcRes, error) {
	_, err := s.AsAllGroup.Get(ctx, sourceGroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = s.AsGroupTenantRelations.Update().Where(
		asgrouptenantrelations.Type(2), asgrouptenantrelations.ParentID(sourceGroupId)).SetIsDeleted(1).Save(ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.Result("true", s.CreateGroupTenantRelation(ctx, groupId, tenantIds, 102))
}

// 查找树顶级节点
func (s *SystemStore) GetTopGroupId(ctx context.Context, groupId int64) (int64, error) {
	grl, err := s.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.SonID(groupId), asgrouptenantrelations.Type(1)).All(context.Background())
	if err != nil {
		return groupId, err
	}
	if len(grl) < 1 {
		return groupId, nil
	} else {
		return s.GetTopGroupId(ctx, grl[0].ParentID)
	}
}

//人员插入
func (s *SystemStore) CreatePerson(ctx context.Context, tx *schema.Tx, tenantCode string, name string, code string, phoneNumber string, roleIds []int64, deptIds []int64) error {
	var isMaster int64 = 1
	pwd, err := tools.BcryptEncryptPwd(phoneNumber[5:])
	if err != nil {
		return err
	}
	_, err = tx.AsUser.Update().Where(asuser.PhoneNumber(phoneNumber),
		asuser.IsDeleted(0), asuser.TenantCode("ZH200205214156yVEQnZ")).
		SetIsDeleted(1).Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.AsPerson.Update().Where(asperson.PhoneNumber(phoneNumber),
		asperson.IsDeleted(0), asperson.TenantCode("ZH200205214156yVEQnZ")).
		SetIsDeleted(1).Save(ctx)
	if err != nil {
		return err
	}
	userEntitys, err := s.AsUser.Query().Where(asuser.IsDeleted(0), asuser.PhoneNumber(phoneNumber)).All(ctx)
	if err != nil {
		return err
	}
	if len(userEntitys) > 0 {
		pwd = userEntitys[0].Pwd
		for _, u := range userEntitys {
			if u.IsMaster == 1 {
				isMaster = 0
				break
			}
		}
	}
	userEntity, err := tx.AsUser.Create().
		SetUserName(name).
		SetPhoneNumber(phoneNumber).
		SetStatus(3).
		SetIsCreated(3).
		SetTenantApplyingState(2).
		SetTenantCode(tenantCode).
		SetPwd(pwd).
		SetIsMaster(isMaster).
		AddRoleIDs(roleIds...).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.AsPerson.Create().
		SetPhoneNumber(phoneNumber).
		SetRealName(name).
		SetTenantCode(tenantCode).
		SetUserCode(code).
		SetUserx(userEntity).
		SetStatus(3).
		AddAgencyIDs(deptIds...).
		Save(ctx)
	return err
}
