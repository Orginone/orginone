package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsUser holds the schema definition for the AsUser entity.
type AsUser struct {
	ent.Schema
}

// Table Name
func (AsUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_user"},
	}
}

// Fields of the AsUser.
func (AsUser) Fields() []ent.Field {
	return ExpandFields(
		field.String("tenant_code").Optional().Comment("租户编码"),
		field.String("pwd").Sensitive().Comment("密码"),
		field.String("phone_number").Optional().Comment("账号"),
		field.Int64("is_admin").Default(0).Comment("用户是否为对应租户的超级管理员"),
		field.Int64("tenant_applying_state").Default(0).Comment("与租户的隶属关系; 0-注册,新增来的,1-申请的并在审核中的,2-审核通过的已加入的)3 审核拒绝的,4-全部的 5-0和2的集合"),
		field.Int64("is_master").Optional().Default(1).Comment("主从关系;1->主;0->从;"),
		field.Int64("is_created").Default(3).Comment("0-租户创建者,1-租户单位管理员,2-既是租户创建者又是单位管理员,3-都不是"),
		field.String("open_id").Optional().Comment("移动端openId"),
		field.String("user_name").Optional().Comment("用户名"))
}

// Edges of the AsUser.
func (AsUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("person", AsPerson.Type).Unique(),
		edge.To("roles", AsRole.Type).StorageKey(
			edge.Table("as_user_role"),
			edge.Columns("user_id", "role_id"),
		),
		edge.To("jobs", AsJob.Type).StorageKey(
			edge.Table("as_user_job"),
			edge.Columns("user_id", "job_id"),
		),
		edge.To("usedapps", AsMarketUsedApp.Type),
		edge.To("workingDatas", AsWorkingData.Type),
		edge.From("agencys", AsInnerAgency.Type).Ref("users"),
		edge.To("roleDistribs", AsMarketRoleDistribution.Type),
		edge.To("appMenusUserSorts", AsMarketMenuUserSort.Type),
		edge.To("appUserTemplates", AsMarketAppUserTemplate.Type),
	}
}
