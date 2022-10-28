package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsProperties holds the schema definition for the AsProperties entity.
type AsProperties struct {
	ent.Schema
}

// Table Name
func (AsProperties) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_properties"},
	}
}

// Fields of the AsPropertiesDistribution.
func (AsProperties) Fields() []ent.Field {
	return ExpandFields(
		field.String("properties_name").Comment("岗位名称"),
		field.Int64("group_id").Comment("租户id"),
	)
}

// Edges of the AsPropertiesDistribution.
func (AsProperties) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("allTenants", AsPropertiesDistribution.Type),
	}
}
