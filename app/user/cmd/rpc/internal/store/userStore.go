package store

import (
	"context"
	"errors"
	"fmt"
	"orginone/common"
	"orginone/common/models"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asrole"
	"orginone/common/schema/astenant"
	"orginone/common/schema/astenantattr"
	"orginone/common/schema/asunit"
	"orginone/common/tools"
)

type UserStore struct {
	*common.DbStore
}

func NewUserStore(c models.DbStoreConfig) *UserStore {
	return &UserStore{
		DbStore: common.NewDbStore(c),
	}
}

// 查询全新人员编号
func (s *UserStore) CreateUserCode(ctx context.Context, tenantCode string) (string, error) {
	unitEntity, err := s.AsUnit.Query().Where(asunit.IsDeleted(0), asunit.HasTenantxWith(astenant.TenantCode(tenantCode))).First(ctx)
	if err != nil {
		return "DEF", err
	}
	if unitEntity == nil {
		return "DEF", errors.New("unit not found !")
	}
	if tools.IsNilOrEmpty(unitEntity.UnitCode) {
		if len(unitEntity.SocialCreditCode) < 17 {
			return "DEF", err
		}
		unitEntity.UnitCode = unitEntity.SocialCreditCode[8:17]
		unitEntity.Update().SetUnitCode(unitEntity.UnitCode).Save(ctx)
	}
	userCode := "000"
	if len(unitEntity.UnitCode) >= 3 {
		userCode = unitEntity.UnitCode[0:3]
	}
	maxNum, err := s.AsPerson.Query().Where(asperson.TenantCode(tenantCode), asperson.UserCodeHasPrefix(userCode)).Count(ctx)
	if err != nil {
		return "DEF", err
	}
	return fmt.Sprintf("%s%06d", userCode, maxNum+1), err
}

//获取租户类型Id
//tenantType 1:普通租户,2:平台租户,3:默认租户,4:集团租户,5:开发租户
func (s *UserStore) GetTenantTypeId(ctx context.Context, tenantType int64) (int64, error) {

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
func (s *UserStore) GetRoleId(ctx context.Context, roleType int64) (int64, error) {
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
