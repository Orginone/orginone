package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppRole holds the schema definition for the AsMarketAppRole entity.
type AsMarketAppRole struct {
	ent.Schema
}

// Table Name
func (AsMarketAppRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_role"},
	}
}

// Fields of the AsMarketAppRole.
func (AsMarketAppRole) Fields() []ent.Field {
	return ExpandFields(
		field.String("role_name").Optional().Comment("角色名称"),
		field.Int64("app_id").Optional().StructTag(`json:"appId,string"`).Comment("应用主键"))
}

// Edges of the AsMarketAppRole.
func (AsMarketAppRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roleDistribs", AsMarketRoleDistribution.Type),
		edge.To("roleMenus", AsMarketRoleMenu.Type),
		edge.To("menus", AsMarketMenu.Type).StorageKey(
			edge.Table("as_market_role_menu"),
			edge.Columns("role_id", "menu_id"),
		),
		edge.From("appx", AsMarketApp.Type).Ref("appRoles").Field("app_id").Unique(),
	}
}
