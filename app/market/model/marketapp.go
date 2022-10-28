package model

import (
	"orginone/common/schema"
)

type AppModel struct {
	*schema.AsMarketApp
	TenantName  string `json:"tenantName"`
	UnitName    string `json:"unitName"`
	UnitAddress string `json:"unitAddress"`
	Province    string `json:"province"`
	City        string `json:"city"`
	AppKey      string `json:"appKey"`
	AppSecret   string `json:"appSecret"`
}

type AppRedeployModel struct {
	*AppModel
	RedeployId     int64 `json:"redeployId,string"`
	RedeployStatus int64 `json:"redeployStatus"`
}

type AppMarketMenuModel struct {
	*schema.AsMarketMenu
	RoleList     []*schema.AsMarketAppRole `json:"roleList"`
	RoleNameList []string                  `json:"roleNameList"`
	Children     []*AppMarketMenuModel     `json:"children"`
}

type AsMarketAppRoleModel struct {
	*schema.AsMarketAppRole
	DistributionDept   int64 `json:"distributionDept"`
	DistributionJob    int64 `json:"distributionJob"`
	DistributionPerson int64 `json:"distributionPerson"`
}

type AppModelDetail struct {
	App struct {
		*schema.AsMarketApp
		TenantName     string `json:"tenantName"`
		UnitName       string `json:"unitName"`
		UnitAddress    string `json:"unitAddress"`
		Province       string `json:"province"`
		City           string `json:"city"`
		RedeployId     int64  `json:"redeployId,string"`
		RedeployStatus int64  `json:"redeployStatus"`
		AppKey         string `json:"appKey"`
		AppSecret      string `json:"appSecret"`
	} `json:"app"`
	RoleList     []*AsMarketAppRoleModel `json:"roleList"`
	RoleNameList []string                `json:"roleNameList"`
	Menu         []*AppMarketMenuModel   `json:"menu"`
}

type UnitModel struct {
	*schema.AsUnit
	VirtualUnit bool   `json:"virtualUnit"`
	IsVirtual   int64  `json:"isVirtual"`
	TenantCode  string `json:"tenantCode"`
	TenantType  int64  `json:"tenantType"`
}

type GroupModel struct {
	*schema.AsAllGroup
	UseStatus int64  `json:"useStatus"`
	LinkMan   string `json:"linkMan"`
	LinkPhone string `json:"linkPhone"`
	UnitName  string `json:"unitName"`
}

type RoleModel struct {
	RoleId   int64  `json:"id,omitempt,string"`
	RoleName string `json:"roleName"`
	IsDelete int64  `json:"isDelete,omitempt"`
}

type MenuTreeModel struct {
	Icon         string           `json:"icon"`
	MenuUrl      string           `json:"menuUrl"`
	MenuName     string           `json:"menuName"`
	InnerUrl     string           `json:"innerUrl,omitempty"`
	HttpsMenuUrl string           `json:"httpsMenuUrl,omitempty"`
	RoleList     []*RoleModel     `json:"roleList"`
	Children     []*MenuTreeModel `json:"children"`
}

type AppMarketMenuTreeModel struct {
	*schema.AsMarketMenu
	Children []*MenuTreeModel `json:"children"`
}

type PublishAppModel struct {
	AppId         int64                          `json:"appId,string"`
	MenuList      []*MenuTreeModel               `json:"menu"`
	RoleList      []*RoleModel                   `json:"roleList"`
	ComponentList []*schema.AsMarketAppComponent `json:"componentList"`
}

type AppRoleDistribeRes struct {
	RoleId  int64                    `json:"roleId,string"`
	Jobs    []map[string]interface{} `json:"jobIdNameList"`
	Agencys []map[string]interface{} `json:"agencyIdNameList"`
	Users   []map[string]interface{} `json:"userIdNameList"`
}
