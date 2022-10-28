package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsMarketApp holds the schema definition for the AsMarketApp entity.
type AsMarketApp struct {
	ent.Schema
}

// Table Name
func (AsMarketApp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_market_app"},
	}
}

// Fields of the AsMarketApp.
func (AsMarketApp) Fields() []ent.Field {
	return ExpandFields(
		field.String("app_name").Optional().Comment("应用名称"),
		field.String("icon").Optional().Comment("应用图标"),
		field.String("version").Optional().Comment("应用版本"),
		field.String("contact").Optional().Comment("负责人联系方式"),
		field.String("contact_name").Optional().Comment("负责人姓名"),
		field.String("description").Optional().Comment("应用描述"),
		field.String("file").Optional().Comment("部署文件"),
		field.Int64("sale_status").Default(1).Optional().Comment("上下架状态;0-下架,1-上架"),
		field.String("tenant_id").Optional().Comment("上传应用租户ID"),
		field.Int64("platform").Default(1).Optional().Comment("适配平台;1-pc,2-移动端,3-自适应"),
		field.Int64("target_user").Optional().Comment("目标用户;1-个人,2-单位,3-集团"),
		field.Int64("deploy_status").Optional().Comment("部署状态;0-未部署,1-部署中,2-已部署"),
		field.String("deploy_address").Optional().Comment("部署地址"),
		field.Int64("deploy_type").Default(1).Optional().Comment("部署方式;0-集成部署,1-独立部署"),
		field.Time("publish_time").GoType(date.Now()).Optional().Comment("发布时间"),
		field.Int64("app_type").Optional().Comment("应用类型"),
		field.Time("apply_time").GoType(date.Now()).Optional().Comment("申请时间"),
		field.String("app_address").Optional().Comment("应用地址"),
		field.String("app_mail").Optional().Comment("应用邮箱"),
		field.String("app_photo").Optional().Comment("应用照片"),
		field.Int64("app_field").Optional().Comment("应用领域;13-党政智治、145-数字政府、37-数字社会、57-数字经济、58-数字法制"),
		field.Int64("app_category").Optional().Comment("应用分类;1-IaaS应用、2-PaaS应用、3-DaaS应用、5-SaaS应用"),
		field.Int64("app_project_source").Optional().Comment("项目来源;1-正常立项、2-揭榜挂帅、3-自主建设"),
		field.Int64("app_star").Comment("观星台;1-是、0否"),
		field.Int64("app_founds_source").Optional().Comment("资金来源;1-预算内、2-预算暂存、3-财政专户、4-其他资金、5-财政暂存、6-社保资金、7-专项资金"),
		field.String("inner_url").Optional().Comment("内网url"),
		field.String("out_url").Optional().Comment("外网url"),
		field.Int64("reform_status").Comment("整改状态;0-已认证,1-整改中"),
		field.String("out_ip_url").Optional().Comment("外网IPurl"))
}

// Edges of the AsMarketApp.
func (AsMarketApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appMenus", AsMarketMenu.Type),
		edge.To("appRoles", AsMarketAppRole.Type),
		edge.To("useds", AsMarketUsedApp.Type),
		edge.To("appAlerts", AsMarketAppAlert.Type),
		edge.To("appRedeploys", AsRedeployData.Type),
		edge.To("appKeys", AsMarketAppKeySecret.Type),
		edge.To("appPurchases", AsMarketAppPurchase.Type),
		edge.To("appComponents", AsMarketAppComponent.Type),
		edge.To("appGroupDistribs", AsMarketAppGroupDistribution.Type),
		edge.To("appGroupDistribConfigs", AsAppGroupDistributionData.Type),
		edge.To("appGroupDistribsRelation", AsMarketAppGroupDistributionRelation.Type),
	}
}
