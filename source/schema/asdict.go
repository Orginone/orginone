package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsDict holds the schema definition for the AsDict entity.
type AsDict struct {
	ent.Schema
}

// Table Name
func (AsDict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_dict"},
	}
}

// Fields of the AsDict.
func (AsDict) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("parent_id").Optional().StructTag(`json:"parentId,string"`).Comment("父主键"),
		field.String("code").Optional().Comment("字典码"),
		field.Int64("dict_key").Optional().Comment("字典值"),
		field.String("dict_value").Optional().Comment("字典名称"),
		field.Int64("sort").Optional().Comment("排序"),
		field.String("remark").Optional().Comment("字典备注"),
		field.Int64("currversion").Optional().Comment("当前版本号"),
		field.Int64("version").Default(0).Optional().Comment("版本号"),
		field.Int64("dictparent_id").Optional().Comment("枚举父节点"))
}

// Edges of the AsDict.
func (AsDict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", AsDict.Type).From("parentx").Field("parent_id").Unique(),
	}
}
