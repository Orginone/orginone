package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsWorkingData holds the schema definition for the AsWorkingData entity.
type AsWorkingData struct {
	ent.Schema
}

// Table Name
func (AsWorkingData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_working_data"},
	}
}

// Fields of the AsWorkingData.
func (AsWorkingData) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Comment("应用ID"),
		field.Int64("user_id").Comment("用户ID"),
		field.Int64("type").Comment("在途业务类型 1-用户 2-部门 3-岗位"))
}

// Edges of the AsUser.
func (AsWorkingData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", AsUser.Type).Ref("workingDatas").Field("user_id").Unique().Required()}
}
