package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsLayer holds the schema definition for the AsLayer entity.
type AsLayer struct {
	ent.Schema
}

// Table Name
func (AsLayer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_layer"},
	}
}

// Fields of the AsLayer.
func (AsLayer) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("layer").Optional().Comment("层级"),
		field.Int64("width").Optional().Comment("宽度"),
		field.Int64("group_id").Optional().Comment("宽度"),
	)
}

// Edges of the AsLayer.
func (AsLayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", AsAllGroup.Type).Ref("allLayer").Field("group_id").Unique(),
	}
}
