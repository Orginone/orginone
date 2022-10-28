package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsPropertiesDistribution holds the schema definition for the AsPropertiesDistribution entity.
type AsPropertiesDistribution struct {
	ent.Schema
}

// Table Name
func (AsPropertiesDistribution) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_properties_distribution"},
	}
}

// Fields of the AsPropertiesDistribution.
func (AsPropertiesDistribution) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("properties_id").Optional().Comment("父集团或租户id"),
		field.Int64("tenant_id").Optional().Comment("子集团或租户id"),
	)
}

// Edges of the AsPropertiesDistribution.
func (AsPropertiesDistribution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("properties", AsProperties.Type).Ref("allTenants").Field("properties_id").Unique(),
		edge.From("tenant", AsTenant.Type).Ref("allproperties").Field("tenant_id").Unique(),
	}
}
