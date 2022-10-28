package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AsInputData holds the schema definition for the AsInputData entity.
type AsInputData struct {
	ent.Schema
}

// Table Name
func (AsInputData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_input_data"},
	}
}

// Fields of the AsInputData.
func (AsInputData) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("file_id").Comment("导入文件ID"),
		field.String("file_name").Comment("导入的文件名称"),
		field.String("table_name").Comment("表名称"),
		field.Int64("type").Optional().Comment("导入类型 0-文件 1-数据"),
		field.Int64("t_count").Optional().Comment("导入成功条数"),
		field.Int64("f_count").Optional().Comment("导入失败条数"),
		field.String("context").Optional().Comment("导入失败原因"),
		field.Time("end_time").GoType(date.Now()).Optional().Comment("结束时间"),
		field.Int64("total_time").Optional().Optional().Comment("总共耗时"),
		field.String("tenant_code").Optional().Comment("所属租户ID"),
		field.Int64("import_type").Optional().Optional().Comment(""))
}

// Edges of the AsInputData.
func (AsInputData) Edges() []ent.Edge {
	return nil
}
