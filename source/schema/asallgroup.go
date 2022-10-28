package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsAllGroup holds the schema definition for the AsAllGroup entity.
type AsAllGroup struct {
	ent.Schema
}

// Table Name
func (AsAllGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_all_group"},
	}
}

// Fields of the AsAllGroup.
func (AsAllGroup) Fields() []ent.Field {
	return ExpandFields(
		field.String("tenant_code").Optional().Comment("租户编码(管理租户,默认为创建租户)"),
		field.String("group_name").Optional().Comment("集团名称"),
		field.String("group_description").Optional().Comment("集团描述"),
		field.String("group_code").Optional().Comment("集团编码"),
		field.Int64("depth").Default(0).Optional().Comment("深度"),
		field.Int64("type").Default(2).Optional().Comment("类型1-集团,2-虚节点)"),
		field.String("social_credit_code").Optional().Comment("统一社会信用代码"))
}

// Edges of the AsAllGroup.
func (AsAllGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("allLayer", AsLayer.Type),
		edge.To("appPurchases", AsMarketAppPurchase.Type),
		edge.To("allTenants", AsGroupTenantRelations.Type),
		edge.To("appGroupDistribs", AsMarketAppGroupDistribution.Type),
		edge.To("appGroupDistribConfigs", AsAppGroupDistributionData.Type),
		edge.To("appGroupDistribsRelation", AsMarketAppGroupDistributionRelation.Type),
	}
}
