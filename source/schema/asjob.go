package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsJob holds the schema definition for the AsJob entity.
type AsJob struct {
	ent.Schema
}

// Table Name
func (AsJob) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_job"},
	}
}

// Fields of the AsJob.
func (AsJob) Fields() []ent.Field {
	return ExpandFields(
		field.String("job_name").Comment("岗位名称"),
		field.String("tenant_code").Comment("租户编码"),
		field.Int64("type").Default(1).Optional().Comment("岗位类型;1-系统默认岗位,2-用户自定义岗位,3-应用岗位"),
		field.Int64("sort").Optional().Comment("排序"))
}

// Edges of the AsJob.
func (AsJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("persons", AsPerson.Type).StorageKey(
			edge.Table("as_job_person"),
			edge.Columns("job_id", "person_id"),
		),
		edge.From("roles", AsRole.Type).Ref("jobs"),
		edge.From("users", AsUser.Type).Ref("jobs"),
		edge.From("agencys", AsInnerAgency.Type).Ref("jobs"),
		edge.To("roleDistribs", AsMarketRoleDistribution.Type),
	}
}
