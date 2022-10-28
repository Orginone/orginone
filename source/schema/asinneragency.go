package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsInnerAgency holds the schema definition for the AsInnerAgency entity.
type AsInnerAgency struct {
	ent.Schema
}

// Table Name
func (AsInnerAgency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_inner_agency"},
	}
}

// Fields of the AsInnerAgency.
func (AsInnerAgency) Fields() []ent.Field {
	return ExpandFields(
		field.String("agency_name").Comment("机构名称"),
		field.String("agency_code").Optional().Comment("部门编码(系统生成)"),
		field.String("tenant_code").Comment("所属租户ID"),
		field.Int64("parent_id").Optional().StructTag(`json:"parentId,string"`).Comment("父Id"))
}

// Edges of the AsInnerAgency.
func (AsInnerAgency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", AsInnerAgency.Type).From("parent").Field("parent_id").Unique(),
		edge.To("jobs", AsJob.Type).StorageKey(
			edge.Columns("agency_id", "job_id"),
			edge.Table("as_agency_job"),
		),
		edge.To("users", AsUser.Type).StorageKey(
			edge.Columns("agency_id", "user_id"),
			edge.Table("as_agency_user"),
		),
		edge.To("persons", AsPerson.Type).StorageKey(
			edge.Columns("agency_id", "person_id"),
			edge.Table("as_agency_person"),
		),
		edge.To("roleDistribs", AsMarketRoleDistribution.Type),
	}
}
