package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AsPersonSingle holds the schema definition for the AsPersonSingle entity.
type AsPersonSingle struct {
	ent.Schema
}

// Table Name
func (AsPersonSingle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_person_single"},
	}
}

// Fields of the AsPersonSingle.
func (AsPersonSingle) Fields() []ent.Field {
	return ExpandFields(
		field.String("real_name").Comment("姓名"),
		field.String("id_card").Optional().Immutable().Comment("身份证号"),
		field.Int64("gender").Optional().Comment("性别;1、男2、女"),
		field.Time("user_birthday").GoType(date.Now()).Optional().Comment("出生日期"),
		field.String("user_email").Optional().Comment("邮箱号码"),
		field.String("user_photo").Optional().Comment("照片"),
		field.String("phone_number").Comment("手机号码"),
		field.String("province").Optional().Comment("省"),
		field.String("city").Optional().Comment("市"),
		field.String("street_address").Optional().Comment("地区"))
}

// Edges of the AsPersonSingle.
func (AsPersonSingle) Edges() []ent.Edge {
	return []ent.Edge{}
}
