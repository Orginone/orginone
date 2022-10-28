package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketAppAlert holds the schema definition for the AsMarketAppAlert entity.
type AsMarketAppAlert struct {
	ent.Schema
}

// Table Name
func (AsMarketAppAlert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_alert"},
	}
}

// Fields of the AsMarketAppAlert.
func (AsMarketAppAlert) Fields() []ent.Field {
	return ExpandFields(
		field.String("alert_title").Optional().Comment("标题"),
		field.String("alert_content").Optional().Comment("内容"),
		field.String("alert_business").Optional().Comment("业务"),
		field.Int64("alert_emergency_level").Optional().Comment("紧急程度"),
		field.Int64("alert_release_app_id").Optional().Comment("发布方"),
		field.Time("alert_release_time").GoType(date.Now()).Optional().Comment("发布时间"),
		field.String("alert_role_ids").Optional().Comment("分发对象角色"),
		field.String("alert_job_ids").Optional().Comment("分发对象岗位"),
		field.Int64("alert_status").Optional().Comment("预警消息状态"),
	)
}

// Edges of the AsMarketAppAlert.
func (AsMarketAppAlert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appx", AsMarketApp.Type).Ref("appAlerts").Field("alert_release_app_id").Unique(),
	}
}
