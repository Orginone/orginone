package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AsTenantIcon holds the schema definition for the AsTenantIcon entity.
type AsTenantIcon struct {
	ent.Schema
}

// Table Name
func (AsTenantIcon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_tenant_icon"},
	}
}

// Fields of the AsTenantIcon.
func (AsTenantIcon) Fields() []ent.Field {
	return ExpandFields(
		field.String("tenant_code").Optional().Comment("租户编码"),
		field.String("icon").Optional().Comment("图标"))
}

// Edges of the AsTenantIcon.
func (AsTenantIcon) Edges() []ent.Edge {
	return nil
}
