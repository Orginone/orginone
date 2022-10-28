package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppComponentTemplate holds the schema definition for the AsMarketAppComponentTemplate entity.
type AsMarketAppComponentTemplate struct {
	ent.Schema
}

// Table Name
func (AsMarketAppComponentTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_component_template"},
	}
}

// Fields of the AsMarketAppComponentTemplate.
func (AsMarketAppComponentTemplate) Fields() []ent.Field {
	return ExpandFields(
		field.String("name").Optional().Comment("模板名称"),
		field.String("config").Optional().Comment("配置"),
		field.Int64("is_default").Optional().Comment("是否是默认模板"))
}

// Edges of the AsMarketAppComponentTemplate.
func (AsMarketAppComponentTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appUserTemplates", AsMarketAppUserTemplate.Type),
	}
}
