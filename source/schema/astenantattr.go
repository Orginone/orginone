package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsTenantAttr holds the schema definition for the AsTenantAttr entity.
type AsTenantAttr struct {
	ent.Schema
}

// Table Name
func (AsTenantAttr) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_tenant_attr"},
	}
}

// Fields of the AsTenantAttr.
func (AsTenantAttr) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("parent_id").Optional().StructTag(`json:"parentId,string"`).Comment("父ID"),
		field.String("attr_name").Comment("tenantType对应的名字,也就是租户类型的名字"),
		field.String("attr_remark").Optional().Comment("类型备注"))
}

// Edges of the AsTenantAttr.
func (AsTenantAttr) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", AsTenantAttr.Type).From("parentx").Field("parent_id").Unique(),
		edge.To("attrRoles", AsTenantAttrRole.Type),
	}
}
