package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsRole holds the schema definition for the AsRole entity.
type AsRole struct {
	ent.Schema
}

// Table Name
func (AsRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_role"},
	}
}

// Fields of the AsRole.
func (AsRole) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("sort").Comment("排序"),
		field.String("role_alias").Optional().Comment("角色别名角色名的首字母拼音),可用于接口鉴权,自动生成"),
		field.String("role_name").Optional().Comment("角色名"),
		field.String("role_description").Optional().Comment("角色描述"))
}

// Edges of the AsRole.
func (AsRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", AsUser.Type).Ref("roles"),
		edge.To("jobs", AsJob.Type).StorageKey(
			edge.Table("as_role_job"),
			edge.Columns("role_id", "job_id"),
		),
		edge.From("menus", AsMenu.Type).Ref("roles"),
		edge.To("attrRoles", AsTenantAttrRole.Type),
	}
}
