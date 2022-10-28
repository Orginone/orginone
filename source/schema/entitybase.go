package schema

import (
	"orginone/common/tools/date"
	"orginone/common/tools/snowflake"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func ExpandFields(custom ...ent.Field) []ent.Field {
	fields := make([]ent.Field, 1)
	fields[0] = field.Int64("id").Unique().Immutable().
		DefaultFunc(snowflake.SnowGen.NextVal).StructTag(`json:"id,string"`).Comment("主键")
	fields = append(fields, custom...)
	fields = append(fields,
		field.Int64("is_deleted").Default(0).Comment("是否被删除"),
		field.Int64("status").Default(2).Optional().Comment("状态"),
		field.Int64("create_user").Optional().Comment("创建用户"),
		field.Int64("update_user").Optional().Comment("更新用户"),
		field.Time("create_time").GoType(date.Now()).Default(date.Now).Immutable().Optional().Comment("创建时间"),
		field.Time("update_time").GoType(date.Now()).Default(date.Now).UpdateDefault(date.Now).Optional().Comment("更新时间"))
	return fields
}
