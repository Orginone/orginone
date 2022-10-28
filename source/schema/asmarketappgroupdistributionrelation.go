package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppGroupDistributionRelation holds the schema definition for the AsMarketAppGroupDistributionRelation entity.
type AsMarketAppGroupDistributionRelation struct {
	ent.Schema
}

// Table Name
func (AsMarketAppGroupDistributionRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_group_distribution_relation"},
	}
}

// Fields of the AsMarketAppGroupDistributionRelation.
func (AsMarketAppGroupDistributionRelation) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("应用id"),
		field.Int64("relation_id").Optional().Comment("被分配租户id"),
		field.Int64("group_id").Optional().Comment("集团id"))
}

// Edges of the AsMarketAppGroupDistributionRelation.
func (AsMarketAppGroupDistributionRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appGroupDistribsRelation").Field("app_id").Unique(),
		edge.From("groupx", AsAllGroup.Type).Ref("appGroupDistribsRelation").Field("group_id").Unique(),
	}
}
