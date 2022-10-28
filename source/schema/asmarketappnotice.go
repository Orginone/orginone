package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AsMarketAppNotice holds the schema definition for the AsMarketAppNotice entity.
type AsMarketAppNotice struct {
	ent.Schema
}

// Table Name
func (AsMarketAppNotice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app_notice"},
	}
}

// Fields of the AsAllGroup.
func (AsMarketAppNotice) Fields() []ent.Field {
	return ExpandFields(
		field.String("notice_title").Optional().Comment("标题"),
		field.String("notice_content").Optional().Comment("内容"),
		field.Int64("notice_release_unit_id").Optional().Comment("发布方"),
		field.Int64("notice_release_status").Comment("发布状态"),
		field.Time("notice_release_time").GoType(date.Now()).Optional().Comment("发布时间"),
		field.String("notice_role_ids").Optional().Comment("发布对象角色"),
		field.String("notice_unit_ids").Optional().Comment("发布对象单位"),
		field.Int64("group_or_unit").Optional().Comment("集团或单位"),
		field.String("unit_query_ids").Optional().Comment("单位ids用于查询"))
}

// Edges of the AsAllGroup.
func (AsMarketAppNotice) Edges() []ent.Edge {
	return nil
}
