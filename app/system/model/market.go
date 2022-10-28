package model

import "orginone/common/schema"

//应用菜单树
type MarketAppMenuTree struct {
	*schema.AsMarketMenu
	Children []*MarketAppMenuTree `json:"children"`
}
