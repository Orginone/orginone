package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsTenantAttrRole holds the schema definition for the AsTenantAttrRole entity.
type AsTenantAttrRole struct {
	ent.Schema
}

// Table Name
func (AsTenantAttrRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_tenant_attr_role"},
	}
}

// Fields of the AsTenantAttrRole.
func (AsTenantAttrRole) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("attr_id").Optional().Comment("租户特性ID"),
		field.Int64("role_id").Optional().Comment("角色ID"),
		field.Int64("is_default").Default(0).Comment("是否为默认角色;1-为默认,"))
}

// Edges of the AsTenantAttrRole.
func (AsTenantAttrRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenantAttrx", AsTenantAttr.Type).Ref("attrRoles").Field("attr_id").Unique(),
		edge.From("rolex", AsRole.Type).Ref("attrRoles").Field("role_id").Unique(),
	}
}
