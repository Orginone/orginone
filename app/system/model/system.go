package model

import "orginone/common/schema"

//系统菜单树
type MenuTree struct {
	*schema.AsMenu
	Children []*MenuTree `json:"children"`
}

//组织菜单树
type InnerAgencyTree struct {
	*schema.AsInnerAgency
	Children []*InnerAgencyTree `json:"children"`
}

//行政规划树
type AdministrativeTree struct {
	Lable    string                `json:"label"`
	Value    string                `json:"value"`
	Children []*AdministrativeTree `json:"children"`
}

//行政规划树
type UnitInfo struct {
	Lable    string                `json:"label"`
	Value    string                `json:"value"`
	Children []*AdministrativeTree `json:"children"`
}

//单位信息类
type UnitModel struct {
	*schema.AsUnit
	Flag       int64  `json:"flag"`
	TenantType int64  `json:"tenantType"`
	TenantCode string `json:"tenantCode"`
}

//集团集合
type AllGroupRecord struct {
	*schema.AsAllGroup
	IsCreate         int64  `json:"isCreate,default=-1"`
	UnitName         string `json:"unitName"`
	LinkMan          string `json:"linkMan"`
	LinkPhone        string `json:"linkPhone"`
	SocialCreditCode string `json:"socialCreditCode"`
}

type TenantModel struct {
	*schema.AsTenant
	IsMaster            int64  `json:"isMaster"`
	IsCreated           int64  `json:"isCreated"`
	TenantApplyingState int64  `json:"tenantApplyingState"`
	TenantTypeName      string `json:"tenantTypeName"`
	RealName            string `json:"realName"`
	SocialCreCode       string `json:"socialCreCode"`
	PhoneNumber         string `json:"phoneNumber"`
}
