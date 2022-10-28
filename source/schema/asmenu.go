package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMenu holds the schema definition for the AsMenu entity.
type AsMenu struct {
	ent.Schema
}

// Table Name
func (AsMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_menu"},
	}
}

// Fields of the AsMenu.
func (AsMenu) Fields() []ent.Field {
	return ExpandFields(
		field.Int64("parent_id").Optional().StructTag(`json:"parentId,string"`).Comment("父菜单id"),
		field.String("name").Optional().Comment("菜单名称"),
		field.String("alias").Optional().Comment("菜单别名(拼音自动生成)"),
		field.String("path").Optional().Comment("菜单路径"),
		field.String("icon").Optional().Comment("菜单图标"),
		field.Int64("sort").Optional().Comment("排序值"),
		field.Int64("category").StructTag(`json:"category,string"`).Comment("菜单类型"),
		field.Int64("is_open").Optional().Comment("是否打开新页面"),
		field.String("remark").Optional().Comment("描述"),
		field.Int64("reform_status").Optional())
}

// Edges of the AsMenu.
func (AsMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", AsMenu.Type).From("parentx").Field("parent_id").Unique(),
		edge.To("roles", AsRole.Type).StorageKey(
			edge.Columns("menu_id", "role_id"),
			edge.Table("as_menu_role"),
		),
	}
}
