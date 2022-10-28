package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketMenu holds the schema definition for the AsMarketMenu entity.
type AsMarketMenu struct {
	ent.Schema
}

// Table Name
func (AsMarketMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_menu"},
	}
}

// Fields of the AsMarketMenu.
func (AsMarketMenu) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").StructTag(`json:"appId,string"`).Optional().Comment("应用主键"),
		field.Int64("parent_id").Optional().StructTag(`json:"parentId,string"`).Comment("上级菜单"),
		field.String("menu_name").Optional().Comment("菜单名称"),
		field.String("menu_url").Optional().Comment("菜单url"),
		field.String("menu_column").Optional().Comment("菜单字段"),
		field.String("icon").Optional().Comment("菜单图标"),
		field.Int64("sort").Default(0).Optional().Comment("排序字段"),
		field.String("https_menu_url").Optional().Comment("菜单url"),
		field.Int64("reform_status").Default(0).Comment("整改状态;0-已认证,1-整改中"),
		field.String("out_ip_menu_url").Optional().Comment("外网IPurl"))
}

// Edges of the AsMarketMenu.
func (AsMarketMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", AsMarketMenu.Type).From("parent").Field("parent_id").Unique(),
		edge.From("appx", AsMarketApp.Type).Ref("appMenus").Field("app_id").Unique(),
		edge.To("roleMenus", AsMarketRoleMenu.Type),
		edge.From("roles", AsMarketAppRole.Type).Ref("menus"),
		edge.To("UserSorts", AsMarketMenuUserSort.Type),
	}
}
