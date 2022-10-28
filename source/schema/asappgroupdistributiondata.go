package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsAppGroupDistributionData holds the schema definition for the AsAppGroupDistributionData entity.
type AsAppGroupDistributionData struct {
	ent.Schema
}

// Table Name
func (AsAppGroupDistributionData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_app_group_distribution_data"},
	}
}

// Fields of the AsAppGroupDistributionData.
func (AsAppGroupDistributionData) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("应用id"),
		field.Int64("group_id").Optional().Comment("集团id"),
		field.String("config").Optional().Comment("配置Json"))
}

// Edges of the AsAppGroupDistributionData.
func (AsAppGroupDistributionData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appGroupDistribConfigs").Field("app_id").Unique(),
		edge.From("groupx", AsAllGroup.Type).Ref("appGroupDistribConfigs").Field("group_id").Unique(),
	}
}
