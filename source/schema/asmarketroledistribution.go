package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketRoleDistribution holds the schema definition for the AsMarketRoleDistribution entity.
type AsMarketRoleDistribution struct {
	ent.Schema
}

// Table Name
func (AsMarketRoleDistribution) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_role_distribution"},
	}
}

// Fields of the AsMarketRoleDistribution.
func (AsMarketRoleDistribution) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("role_id").Optional().Comment("角色id"),
		field.Int64("user_id").Optional().Comment("用户id"),
		field.Int64("agency_id").Optional().Comment("部门id"),
		field.Int64("job_id").Optional().Comment("岗位id"),
		field.String("tenant_code").Optional())
}

// Edges of the AsMarketRoleDistribution.
func (AsMarketRoleDistribution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userx", AsUser.Type).Ref("roleDistribs").Field("user_id").Unique(),
		edge.From("agencyx", AsInnerAgency.Type).Ref("roleDistribs").Field("agency_id").Unique(),
		edge.From("jobx", AsJob.Type).Ref("roleDistribs").Field("job_id").Unique(),
		edge.From("rolex", AsMarketAppRole.Type).Ref("roleDistribs").Field("role_id").Unique(),
	}
}
