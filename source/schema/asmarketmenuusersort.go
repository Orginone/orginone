package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketMenuUserSort holds the schema definition for the AsMarketMenuUserSort entity.
type AsMarketMenuUserSort struct {
	ent.Schema
}

// Table Name
func (AsMarketMenuUserSort) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_menu_user_sort"},
	}
}

// Fields of the AsMarketMenuUserSort.
func (AsMarketMenuUserSort) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("user_id").Optional().Comment("用户id"),
		field.Int64("menu_id").Optional().Comment("一级应用菜单"),
		field.Int64("sort").Default(0).Optional().Comment("排序字段"))
}

// Edges of the AsMarketMenuUserSort.
func (AsMarketMenuUserSort) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userx", AsUser.Type).Ref("appMenusUserSorts").Field("user_id").Unique(),
		edge.From("appmenux", AsMarketMenu.Type).Ref("UserSorts").Field("menu_id").Unique(),
	}
}
