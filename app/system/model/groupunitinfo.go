package model

import (
	"orginone/common/schema"
)

//集团信息
type GroupUnitInfo struct {
	Name             string           `json:"name"`
	Below            bool             `json:"below"`
	Children         []*GroupUnitInfo `json:"children"`
	Code             string           `json:"code"`
	GroupCode        string           `json:"groupCode"`
	GroupName        string           `json:"groupName"`
	GroupDescription string           `json:"groupDescription"`
	GroupType        int64            `json:"groupType"`
	Id               int64            `json:"id,string"`
	LinkMan          string           `json:"linkMan"`
	LinkPhone        string           `json:"linkPhone"`
	ManagerTenant    bool             `json:"managerTenant"`
	ParentId         int64            `json:"parentId,string"`
	SocialCreditCode string           `json:"socialCreditCode"`
	TenantCode       string           `json:"tenantCode"`
	TenantId         int64            `json:"tenantId,string"`
	Type             int64            `json:"type"`
	UnitName         string           `json:"unitName"`
	IsCreate         int64            `json:"isCreate,default=-1"`
}

func NewGroupUnitInfo(Type int64, Below bool) *GroupUnitInfo {
	return &GroupUnitInfo{
		Type:     Type,
		Below:    Below,
		Children: make([]*GroupUnitInfo, 0),
	}
}

func (m *GroupUnitInfo) SetManagerCode(ManagerCode string) *GroupUnitInfo {
	m.ManagerTenant = ManagerCode == m.TenantCode
	m.IsCreate = 0
	if m.ManagerTenant {
		m.IsCreate = 1
	}
	return m
}

func (m *GroupUnitInfo) SetChildrens(Childrens ...*GroupUnitInfo) *GroupUnitInfo {
	for _, Children := range Childrens {
		m.Children = append(m.Children, Children)
	}
	return m
}

func (m *GroupUnitInfo) SetParentId(ParentId int64) *GroupUnitInfo {
	m.ParentId = ParentId
	return m
}

func (m *GroupUnitInfo) SetGroup(Group *schema.AsAllGroup) *GroupUnitInfo {
	m.GroupType = Group.Type
	m.Id = Group.ID
	m.Name = Group.GroupName
	m.Code = Group.GroupCode
	m.GroupName = Group.GroupName
	m.GroupCode = Group.GroupCode
	m.GroupDescription = Group.GroupDescription
	return m
}

func (m *GroupUnitInfo) SetGroupRelations(Group *schema.AsGroupTenantRelations) *GroupUnitInfo {
	m.GroupType = Group.Type
	m.Id = Group.ID
	m.Name = Group.Edges.Tenant.Edges.Unit.UnitName
	m.Code = Group.GroupCode
	m.GroupDescription = ""
	return m
}

func (m *GroupUnitInfo) SetTenant(Tenant *schema.AsTenant) *GroupUnitInfo {
	m.TenantId = Tenant.ID
	m.TenantCode = Tenant.TenantCode
	if Tenant.Edges.Unit != nil {
		m.Id = Tenant.Edges.Unit.ID
		m.UnitName = Tenant.Edges.Unit.UnitName
		m.LinkMan = Tenant.Edges.Unit.LinkMan
		m.LinkPhone = Tenant.Edges.Unit.LinkPhone
		m.SocialCreditCode = Tenant.Edges.Unit.SocialCreditCode
	}
	return m
}
