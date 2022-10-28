package model

type AgencyData struct {
	ParentAgencyName string `json:"parentAgencyName,optional"`
	AgencyName       string `json:"agencyName,optional"`
	TenantCode       string `json:"tenantCode,optional"`
	Id               int64  `json:"id,optional"`
}

type GroupRelationData struct {
	GroupName        string `json:"groupName"`
	SocialCreditCode string `json:"socialCreditCode"`
	ParentGroupName  string `json:"parentGroupName"`
	TenantCode       string `json:"tenantCode"`
	SourceGroupId    int64  `json:"sourceGroupId"`
	UserId           int64  `json:"userId"`
}
