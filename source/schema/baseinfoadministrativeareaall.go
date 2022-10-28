package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Baseinfoadministrativeareaall holds the schema definition for the Baseinfoadministrativeareaall entity.
type Baseinfoadministrativeareaall struct {
	ent.Schema
}

// Table Name
func (Baseinfoadministrativeareaall) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "baseinfo_administrative_area_all"},
	}
}

// Fields of the Baseinfoadministrativeareaall.
func (Baseinfoadministrativeareaall) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("pid").Optional().Comment("上级行政区域"),
		field.String("code").Comment("行政区域编码"),
		field.String("name").Comment("行政区域名称"),
		field.String("province").Comment("省/直辖市"),
		field.String("city").Comment("市"),
		field.String("area").Comment("区"),
		field.String("town").Comment("城镇地区"),
		field.String("all_name").Comment("全名"),
		field.Int32("type").Default(1).Optional().Comment("行政区域级别"),
		field.Int32("ts_version").Default(1).Optional().Comment("乐观锁字段"))
}

// Edges of the Baseinfoadministrativeareaall.
func (Baseinfoadministrativeareaall) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", Baseinfoadministrativeareaall.Type).From("parentx").Field("pid").Unique(),
	}
}
