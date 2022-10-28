package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsGroupTenantRelations holds the schema definition for the AsGroupTenantRelations entity.
type AsGroupTenantRelations struct {
	ent.Schema
}

// Table Name
func (AsGroupTenantRelations) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_group_tenant_relations"},
	}
}

// Fields of the AsGroupTenantRelations.
func (AsGroupTenantRelations) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("parent_id").Optional().Comment("父集团或租户id"),
		field.Int64("son_id").Optional().Comment("子集团或租户id"),
		field.Int64("type").Optional().Comment("类型1-集团,2-租户"),
		field.Int64("sort").Default(0).Optional().Comment("排序"),
		field.Int64("serial").Optional().Comment("序列号"),
		field.String("group_code").Optional().Comment("集团编码"),
		field.Time("expires_time").GoType(date.Now()).Optional().Comment("过期时间"),
		field.Int64("is_hide").Default(0).Optional().Comment("是否在树形中隐藏节点"),
	)
}

// Edges of the AsGroupTenantRelations.
func (AsGroupTenantRelations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", AsAllGroup.Type).Ref("allTenants").Field("parent_id").Unique(),
		edge.From("tenant", AsTenant.Type).Ref("allGroups").Field("son_id").Unique(),
	}
}
