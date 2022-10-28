package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsRedeployData holds the schema definition for the AsRedeployData entity.
type AsRedeployData struct {
	ent.Schema
}

// Table Name
func (AsRedeployData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_redeploy_data"},
	}
}

// Fields of the AsRedeployData.
func (AsRedeployData) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("应用id"),
		field.String("application").Optional().Comment("应用详情"),
	)
}

// Edges of the AsRedeployData.
func (AsRedeployData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appRedeploys").Field("app_id").Unique(),
	}
}
