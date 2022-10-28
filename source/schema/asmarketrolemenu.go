package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketRoleMenu holds the schema definition for the AsMarketRoleMenu entity.
type AsMarketRoleMenu struct {
	ent.Schema
}

// Table Name
func (AsMarketRoleMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_role_menu"},
	}
}

// Fields of the AsMarketRoleMenu.
func (AsMarketRoleMenu) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("role_id").Optional().Comment("角色id"),
		field.Int64("menu_id").Optional().Comment("菜单id"))
}

// Edges of the AsMarketRoleMenu.
func (AsMarketRoleMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menux", AsMarketMenu.Type).Ref("roleMenus").Field("menu_id").Unique(),
		edge.From("rolex", AsMarketAppRole.Type).Ref("roleMenus").Field("role_id").Unique(),
	}
}
