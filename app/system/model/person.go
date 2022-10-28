package model

import (
	"orginone/common/schema"
)

//人员完整信息
type PersonInfo struct {
	*schema.AsPerson
	RoleId     string             `json:"roleId"`
	RoleName   string             `json:"roleName"`
	IfPurchase bool               `json:"ifPurchase"`
	IsCreated  int64              `json:"isCreated"`
	RoleList   []*schema.AsRole   `json:"roleList"`
	DeptList   []*InnerAgencyTree `json:"deptList"`
}

type UpdatePersonMasterData struct {
	PersonName  string `json:"personName"`
	TenantCode  string `json:"tenantCode"`
	PhoneNumber string `json:"phoneNumber"`
}
