package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppGroupDistribution holds the schema definition for the AsMarketAppGroupDistribution entity.
type AsMarketAppGroupDistribution struct {
	ent.Schema
}

// Table Name
func (AsMarketAppGroupDistribution) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_group_distribution"},
	}
}

// Fields of the AsMarketAppGroupDistribution.
func (AsMarketAppGroupDistribution) Fields() []ent.Field {
	return ExpandFields(
		field.String("remark").Optional().Comment("备注"),
		field.Int64("app_id").Optional().Comment("应用id"),
		field.String("tenant_id").Optional().Comment("被分配租户id"),
		field.Int64("group_id").Optional().Comment("集团id"),
		field.Int64("use_status").Default(1).Optional().Comment("0;停用;1;启用"))
}

// Edges of the AsMarketAppGroupDistribution.
func (AsMarketAppGroupDistribution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appGroupDistribs").Field("app_id").Unique(),
		edge.From("groupx", AsAllGroup.Type).Ref("appGroupDistribs").Field("group_id").Unique(),
	}
}
