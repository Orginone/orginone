package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppUserTemplate holds the schema definition for the AsMarketAppUserTemplate entity.
type AsMarketAppUserTemplate struct {
	ent.Schema
}

// Table Name
func (AsMarketAppUserTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_user_template"},
	}
}

// Fields of the AsMarketAppUserTemplate.
func (AsMarketAppUserTemplate) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("user_id").Optional().Comment("用户id"),
		field.Int64("template_id").Optional().Comment("模板ID"),
		field.Int64("use_status").Default(1).Optional().Comment("使用状态(-1-默认模板,1-用户自定义模板)"))
}

// Edges of the AsMarketAppUserTemplate.
func (AsMarketAppUserTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userx", AsUser.Type).Ref("appUserTemplates").Field("user_id").Unique(),
		edge.From("templatex", AsMarketAppComponentTemplate.Type).Ref("appUserTemplates").Field("template_id").Unique(),
	}
}
