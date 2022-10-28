package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppComponent holds the schema definition for the AsMarketAppComponent entity.
type AsMarketAppComponent struct {
	ent.Schema
}

// Table Name
func (AsMarketAppComponent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_component"},
	}
}

// Fields of the AsMarketAppComponent.
func (AsMarketAppComponent) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("app_id").Optional().Comment("app主键"),
		field.String("code").Optional().Comment("编码"),
		field.String("name").Optional().Comment("名称"),
		field.String("url").Optional().Comment("url"),
		field.Int64("type").Optional().Comment("类型"),
		field.String("preview_pic").Optional().Comment("预览图片"),
		field.String("layout_type").Optional().Comment("布局类型"),
		field.String("layout_config").Optional().Comment("布局配置"),
		field.String("tenant_code").Optional().Comment("租户编码"),
		field.String("source").Optional().Comment("平台还剩应用"))
}

// Edges of the AsMarketAppComponent.
func (AsMarketAppComponent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appComponents").Field("app_id").Unique(),
	}
}
