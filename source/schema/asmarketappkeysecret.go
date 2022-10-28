package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppKeySecret holds the schema definition for the AsMarketAppKeySecret entity.
type AsMarketAppKeySecret struct {
	ent.Schema
}

func (AsMarketAppKeySecret) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_key_secret"},
	}
}

// Fields of the AsMarketAppPurchase.
func (AsMarketAppKeySecret) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("app_id主键"),
		field.String("app_key").Optional().Comment("应用id,随机生成"),
		field.String("app_secret").Optional().Comment("应用加密密钥,随机生成"))
}

// Edges of the AsMarketAppKeySecret.
func (AsMarketAppKeySecret) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appKeys").Field("app_id").Unique(),
	}
}
