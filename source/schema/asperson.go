package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsPerson holds the schema definition for the AsPerson entity.
type AsPerson struct {
	ent.Schema
}

// Table Name
func (AsPerson) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_person"},
	}
}

// Fields of the AsPerson.
func (AsPerson) Fields() []ent.Field {
	return ExpandFields(
		field.String("tenant_code").Optional().Comment("租户名称"),
		field.String("real_name").Comment("姓名"),
		field.String("id_card").Optional().Comment("身份证号"),
		field.Int64("gender").Optional().Comment("性别;1、男2、女"),
		field.Time("user_birthday").GoType(date.Now()).Optional().Comment("出生日期"),
		field.String("user_email").Optional().Comment("邮箱号码"),
		field.String("user_photo").Optional().Comment("照片"),
		field.Int64("user_id").StructTag(`json:"userId,string"`).Comment("对应用户id"),
		field.String("phone_number").Comment("手机号码"),
		field.String("province").Optional().Comment("省"),
		field.String("city").Optional().Comment("市"),
		field.String("street_address").Optional().Comment("地区"),
		field.String("user_code").Optional().Comment("用户code"),
		field.Int64("is_master").Optional().Comment("是否是主单位(0-否,1-是)"))
}

// Edges of the AsPerson.
func (AsPerson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userx", AsUser.Type).Ref("person").Field("user_id").Unique().Required(),
		edge.From("agencys", AsInnerAgency.Type).Ref("persons"),
		edge.From("jobs", AsJob.Type).Ref("persons"),
	}
}
