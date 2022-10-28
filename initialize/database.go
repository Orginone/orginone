package main

import (
	"context"
	"encoding/json"
	"log"
	"orginone/common"
	"orginone/common/cache"
	"orginone/common/models"
	"orginone/common/schema"
	"orginone/common/schema/asmenu"
	"orginone/common/schema/astenant"
	"orginone/common/tools"
	"orginone/common/tools/date"
	"orginone/common/tools/snowflake"
	"os"

	schemax "entgo.io/ent/dialect/sql/schema"
	"github.com/google/uuid"
)

func main() {
	store := common.NewDbStore(models.DbStoreConfig{
		Driver:       "mysql",
		Mode:         "dev",
		DataCenterId: 1,
		WorkerId:     1,
		DataSource:   "root:123456@(localhost)/test?charset=utf8&parseTime=true&loc=Local",
		Cache: cache.CacheConfig{
			Addr: "localhost:6379",
		},
	})
	defer store.Client.Close()
	ctx := context.Background()
	// Run the auto migration tool create database.
	if err := store.Schema.Create(ctx, schemax.WithForeignKeys(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// initialize as_menu
	userId := snowflake.SnowGen.NextVal()
	menuKeyMap := make(map[int64]int64)
	count := store.AsMenu.Query().CountX(ctx)
	if count < 1 {
		//初始化菜单表
		menus := make([]*schema.AsMenu, 0)
		buffer, err := os.ReadFile("menus.json")
		if err == nil {
			err = json.Unmarshal(buffer, &menus)
			if err == nil && len(menus) > 0 {
				for _, m := range menus {
					create := store.AsMenu.Create()
					tools.SchemaUpdateAny(create.Mutation(), m)
					if m.ParentID > 0 {
						create.SetParentID(menuKeyMap[m.ParentID])
					}
					create.SetCreateUser(userId).SetUpdateUser(userId)
					item := create.SaveX(ctx)
					menuKeyMap[m.ID] = item.ID
				}
			}
		}
	}
	// initialize as_role
	roleKeyMap := make(map[int64]int64)
	count = store.AsRole.Query().CountX(ctx)
	if count < 1 {
		//初始化角色表
		roles := make([]*schema.AsRole, 0)
		buffer, err := os.ReadFile("roles.json")
		if err == nil {
			err = json.Unmarshal(buffer, &roles)
			if err == nil && len(roles) > 0 {
				for _, r := range roles {
					create := store.AsRole.Create()
					tools.SchemaUpdateAny(create.Mutation(), r)
					create.SetCreateUser(userId).SetUpdateUser(userId)
					item := create.SaveX(ctx)
					roleKeyMap[r.ID] = item.ID
				}
			}
		}
	}
	if len(roleKeyMap) > 0 {
		// initialize as_menu_role
		roleEntity := store.AsRole.GetX(ctx, roleKeyMap[1])
		if count = roleEntity.QueryMenus().CountX(ctx); count == 0 {
			menuEntitys := store.AsMenu.Query().Where(asmenu.IDNotIn(menuKeyMap[1])).AllX(ctx)
			roleEntity.Update().AddMenus(menuEntitys...).SaveX(ctx)
		}
		roleEntity = store.AsRole.GetX(ctx, roleKeyMap[2])
		if count = roleEntity.QueryMenus().CountX(ctx); count == 0 {
			menuEntitys := store.AsMenu.Query().Where(asmenu.Or(asmenu.ParentID(menuKeyMap[3]), asmenu.IDIn(menuKeyMap[2], menuKeyMap[3]))).AllX(ctx)
			roleEntity.Update().AddMenus(menuEntitys...).SaveX(ctx)
		}
		roleEntity = store.AsRole.GetX(ctx, roleKeyMap[7])
		if count = roleEntity.QueryMenus().CountX(ctx); count == 0 {
			menuEntitys := store.AsMenu.Query().Where(asmenu.Or(asmenu.ParentID(menuKeyMap[4]), asmenu.IDIn(menuKeyMap[2], menuKeyMap[4]))).AllX(ctx)
			roleEntity.Update().AddMenus(menuEntitys...).SaveX(ctx)
		}
	}
	// initialize as_tenant_attr
	attrKeyMap := make(map[int64]int64)
	count = store.AsTenantAttr.Query().CountX(ctx)
	if count < 1 {
		//初始化租户特性表
		tenantAttrs := make([]*schema.AsTenantAttr, 0)
		buffer, err := os.ReadFile("tenantAttrs.json")
		if err == nil {
			err = json.Unmarshal(buffer, &tenantAttrs)
			if err == nil && len(tenantAttrs) > 0 {
				for _, a := range tenantAttrs {
					create := store.AsTenantAttr.Create()
					tools.SchemaUpdateAny(create.Mutation(), a)
					create.SetCreateUser(userId).SetUpdateUser(userId)
					item := create.SaveX(ctx)
					attrKeyMap[a.ID] = item.ID
				}
			}
		}
	}
	// initialize as_tenant_attr_role
	count = store.AsTenantAttrRole.Query().CountX(ctx)
	if count < 1 {
		//初始化租户特性角色表
		tenantAttrRoles := make([]*schema.AsTenantAttrRole, 0)
		buffer, err := os.ReadFile("tenantAttrRoles.json")
		if err == nil {
			err = json.Unmarshal(buffer, &tenantAttrRoles)
			if err == nil && len(tenantAttrRoles) > 0 {
				creates := make([]*schema.AsTenantAttrRoleCreate, 0)
				for _, m := range tenantAttrRoles {
					m.AttrID = attrKeyMap[m.AttrID]
					m.RoleID = roleKeyMap[m.RoleID]
					create := store.AsTenantAttrRole.Create()
					tools.SchemaUpdateAny(create.Mutation(), m)
					creates = append(creates, create)
				}
				store.AsTenantAttrRole.CreateBulk(creates...).SaveX(ctx)
			}
		}
	}
	if len(attrKeyMap) > 0 {
		// initialize as_tenant
		if exist := store.AsTenant.Query().Where(astenant.TenantType(attrKeyMap[2])).ExistX(ctx); !exist {
			adminTenant := store.AsTenant.Create().SetTenantName("平台管理租户").SetTenantCode("000000").SetTenantType(attrKeyMap[2]).
				SetCreateUser(userId).SetUpdateUser(userId).SaveX(ctx)
			store.AsUnit.Create().SetTenantx(adminTenant).SetUnitName("平台管理租户").SetIsVirtual("1").
				SetSocialCreditCode("000000").SetUnitNameEn("orginone").SetLinkMan("超管").SetLinkPhone("202005201314").
				SetCreateUser(userId).SetUpdateUser(userId).SetUnitType(attrKeyMap[2]).
				SetUnitCode("000000").SetUnitRemark("system auto gen").SaveX(ctx)
			defaultTenant := store.AsTenant.Create().SetTenantName("默认租户").SetTenantCode("999999").SetTenantType(attrKeyMap[3]).
				SetCreateUser(userId).SetUpdateUser(userId).SaveX(ctx)
			store.AsUnit.Create().SetTenantx(defaultTenant).SetUnitName("默认租户").SetIsVirtual("1").
				SetSocialCreditCode("999999").SetUnitNameEn("orginone").SetLinkMan("超管").SetLinkPhone("202005201314").
				SetCreateUser(userId).SetUpdateUser(userId).SetUnitType(attrKeyMap[3]).
				SetUnitCode("999999").SetUnitRemark("system auto gen").SaveX(ctx)
		}
	}
	// initialize as_user
	if exist := store.AsUser.Query().ExistX(ctx); !exist {
		pwd, _ := tools.BcryptEncryptPwd("1234qwer")
		user := store.AsUser.Create().SetID(userId).SetPwd(pwd).SetTenantCode("000000").SetIsAdmin(1).SetIsCreated(3).SetIsMaster(1).
			SetUserName("超管").SetPhoneNumber("202005201314").SetCreateUser(userId).SetUpdateUser(userId).
			AddRoleIDs(roleKeyMap[1]).SetOpenID(uuid.NewString()).SaveX(ctx)
		store.AsPerson.Create().SetTenantCode("000000").SetIsMaster(1).SetRealName("超管").SetPhoneNumber("202005201314").
			SetCreateUser(userId).SetUpdateUser(userId).SetUserx(user).SetGender(1).SetIDCard(uuid.NewString()).SetUserCode("Admin").
			SetUserBirthday(date.Love()).SaveX(ctx)
	}
	// initialize as_inner_agency
	if exist := store.AsInnerAgency.Query().ExistX(ctx); !exist {
		store.AsInnerAgency.Create().SetAgencyName("虚拟根节点").SetAgencyCode("IG").SetCreateUser(userId).SetUpdateUser(userId).SetTenantCode("000000").SaveX(ctx)
	}
	// initialize baseinfo_administrative_area_all
	areaKeyMap := make(map[int64]int64)
	count = store.Baseinfoadministrativeareaall.Query().CountX(ctx)
	if count < 1 {
		//初始化行政区域表
		areas := make([]*schema.Baseinfoadministrativeareaall, 0)
		buffer, err := os.ReadFile("cityarea.json")
		if err == nil {
			err = json.Unmarshal(buffer, &areas)
			if err == nil && len(areas) > 0 {
				for _, m := range areas {
					create := store.Baseinfoadministrativeareaall.Create()
					tools.SchemaUpdateAny(create.Mutation(), m)
					if m.Pid > 0 {
						create.SetParentxID(areaKeyMap[m.Pid])
					}
					create.SetCreateUser(userId).SetUpdateUser(userId)
					item := create.SaveX(ctx)
					areaKeyMap[m.ID] = item.ID
				}
			}
		}
	}
	// initialize as_dict
	dictKeyMap := make(map[int64]int64)
	count = store.AsDict.Query().CountX(ctx)
	if count < 1 {
		//初始化代码字典
		dicts := make([]*schema.AsDict, 0)
		buffer, err := os.ReadFile("dicts.json")
		if err == nil {
			err = json.Unmarshal(buffer, &dicts)
			if err == nil && len(dicts) > 0 {
				for _, d := range dicts {
					create := store.AsDict.Create()
					tools.SchemaUpdateAny(create.Mutation(), d)
					if d.ParentID > 0 {
						create.SetParentxID(dictKeyMap[d.ParentID])
					}
					create.SetCreateUser(userId).SetUpdateUser(userId)
					item, err := create.Save(ctx)
					if err == nil {
						dictKeyMap[d.ID] = item.ID
					}
				}
			}
		}
	}
}
