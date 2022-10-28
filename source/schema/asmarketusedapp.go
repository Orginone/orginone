package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketUsedApp holds the schema definition for the AsMarketUsedApp entity.
type AsMarketUsedApp struct {
	ent.Schema
}

// Table Name
func (AsMarketUsedApp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_used_app"},
	}
}

// Fields of the AsMarketUsedApp.
func (AsMarketUsedApp) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("应用id"),
		field.Int64("user_id").Optional().Comment("用户id"),
		field.Int64("sort").Default(0).Optional().Comment("排序"))
}

// Edges of the AsMarketUsedApp.
func (AsMarketUsedApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("useds").Field("app_id").Unique(),
		edge.From("userx", AsUser.Type).Ref("usedapps").Field("app_id").Unique(),
	}
}
