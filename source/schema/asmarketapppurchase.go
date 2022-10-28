package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppPurchase holds the schema definition for the AsMarketAppPurchase entity.
type AsMarketAppPurchase struct {
	ent.Schema
}

// Table Name
func (AsMarketAppPurchase) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_purchase"},
	}
}

// Fields of the AsMarketAppPurchase.
func (AsMarketAppPurchase) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("应用主键"),
		field.String("tenant_id").Optional().Comment("租户id"),
		field.Int64("group_id").Optional().Comment("集团id"),
		field.Int64("use_status").Default(1).Optional().Comment("0;停用;1;启用"),
		field.String("remark").Optional().Comment("备注"))
}

// Edges of the AsMarketAppPurchase.
func (AsMarketAppPurchase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appPurchases").Field("app_id").Unique(),
		edge.From("groupx", AsAllGroup.Type).Ref("appPurchases").Field("group_id").Unique(),
	}
}
