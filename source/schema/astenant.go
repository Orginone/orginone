package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsTenant holds the schema definition for the AsTenant entity.
type AsTenant struct {
	ent.Schema
}

// Table Name
func (AsTenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_tenant"},
	}
}

// Fields of the AsTenant.
func (AsTenant) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("tenant_type").Comment("租户类型"),
		field.String("tenant_name").Comment("租户名称"),
		field.String("tenant_code").Comment("租户编码系统生成)"),
		field.String("theme").Default("default").Optional().Comment("租户界面主题"),
		field.Int64("is_virtual").Default(1).Comment("是否虚拟单位1-虚,0-实)"),
	)
}

// Edges of the AsTenant.
func (AsTenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("unit", AsUnit.Type).Unique(),
		edge.To("allGroups", AsGroupTenantRelations.Type),
		edge.To("allproperties", AsPropertiesDistribution.Type),
	}
}
