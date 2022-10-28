package schema

import (
	"orginone/common/tools/date"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AsUnit holds the schema definition for the AsUnit entity.
type AsUnit struct {
	ent.Schema
}

// Table Name
func (AsUnit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "as_unit"},
	}
}

// Fields of the AsUnit.
func (AsUnit) Fields() []ent.Field {
	return ExpandFields(
		field.String("unit_name").Comment("单位名称"),
		field.String("unit_name_en").Optional().Comment("单位英文名称"),
		field.Int64("organization_form").Optional().Comment("组织形式;10、公司制企业11.国有独资企业 12.其他有限责任公司 13.上市股份有限公司 股票代码 14。非上市股份有限公司)20.公司制企业21.非公司制独资企业 22.其他非公司制企业) 30.企业化管理事业单位 40。其他"),
		field.Int64("unit_type").Optional().Comment("单位分类1、主管部门 2、二级单位 3、基层单位 100、政府财政事务 200、安全管理事务 300、"),
		field.String("parent_node_name").Optional().Comment("上级单位名称"),
		field.Int64("charge_section_id").Optional().Comment("主管部门id"),
		field.String("charge_section_code").Optional().Comment("主管部门code"),
		field.String("charge_section_name").Optional().Comment("主管部门名称"),
		field.String("local_financial_code").Optional().Comment("本级财政代码"),
		field.String("local_financial_name").Optional().Comment("本级财政名称"),
		field.String("local_financial_id").Optional(),
		field.String("office_administration_code").Optional().Comment("本级机关事务管理局代码"),
		field.String("office_administration_name").Optional().Comment("本级机关事务管理局名称"),
		field.String("administration_division_code").Optional().Comment("行政区划代码"),
		field.String("administration_division_name").Optional().Comment("行政区划名称"),
		field.Int64("budget_code").Optional().Comment("财政预算代码;编码方法;采用层次码,用数字表示,代码结构为每3位一层。一级单位3位码,二级单位6位码,……,七级单位21位码。其中,预算单位本级由该预算单位的单位层次码加上001表示。中央一级预算单位统一的3位代码表见表5。 应用范围;中央使用,地方参照。 代码	名称 101	国务院办公厅 102	国家发展和改革委员会 105	教育部 106	科学技术部 107	国家国防科技工业局 108	国家民族事务委员会 109	国家体育总局 110	国家人口和计划生育委员会 111	公安部 112	国家安全部 113	司法部 114	外交部 115	监察部 117	人力资源和社会保障部 118	民政部 119	财政部 120	住房和城乡建设部 121	国土资源部 122	铁道部 123	交通运输部 124	工业和信息化部 125	农业部 126	水利部 127	全国社会保障基金理事会"),
		field.String("college_code").Optional().Comment("高等院校代码院校代号即院校代码或学校代码,为全国各高校录取时为方便考生填报志愿而加注的由数字组成的代号串。院校代码就如同是学校的一个身份证号,方便查询学校信息,教育部为高校编排的代码有5位此代码全国通用),各省教育考试院为高校编排代码有4位此代码一般作填报高考志愿用,同一所高校在不同省份代码也不一样),由于高校办学情况每年都有变动,所以高校代码也有变化。 10001 北京大学 北京医科大学并入"),
		field.String("organization_code").Optional().Comment("组织机构代码组织机构代码按照强制性国家标准GB11714《全国组织机构代码编制规则》编制,由八位数字(或大写拉丁字母)本体代码和一位数字(或大写拉丁字母)校验码组成。组织机构代码证组织机构代码证书包括正本、副本和电子副本(IC卡),代码登记部门在为组织机构赋码发证的同时,还要采集28项基础信息,并按照国家标准对这些信息进行编码,将这些信息存入代码数据库和代码证电子副本(IC卡)中,供代码应用部门使用。代码登记部门所采集的基础信息包括:机构名称、机构地址、机构类型、经济性质、行业分类、规模、法人代表等。"),
		field.Int64("unit_type_code").Optional().Comment("单位类型名称代码	名称	说明 1	行政	指依宪法和有关组织法的规定设置的,行使国家行政职权,负责对国家各项行政事务进行组织、管理、监督和指挥的国家机关。这里包括行政、立法、司法、军队、党务等按公务员法管理的单位。 2	事业	为了社会公益目的,从事教育、文化、卫生、科技等活动并以非盈利性为主的社会服务组织。 3	参公	参照公务员法管理的单位。 4、 企业 企业单位 9	其他	除上述行政、事业外的其他单位。"),
		field.Int64("unit_type_name").Optional().Comment("单位类型名称代码	名称	说明 1	行政	指依宪法和有关组织法的规定设置的,行使国家行政职权,负责对国家各项行政事务进行组织、管理、监督和指挥的国家机关。这里包括行政、立法、司法、军队、党务等按公务员法管理的单位。 2	事业	为了社会公益目的,从事教育、文化、卫生、科技等活动并以非盈利性为主的社会服务组织。 3	参公	参照公务员法管理的单位。 4、 企业 企业单位 9	其他	除上述行政、事业外的其他单位。"),
		field.Int64("unit_basic_property").Optional().Comment("单位基本性质"),
		field.String("social_credit_code").Optional().Comment("统一社会信用代码"),
		field.String("budget_unit_name").Optional().Comment("预算单位级次名称代码	名称 1	一级 2	二级 3	三级 4	四级 5	五级 6	六级 7	七级"),
		field.String("budget_unit_code").Optional().Comment("预算单位级次代码"),
		field.Int64("budget_management_level").Optional().Comment("预算管理级次1、中央级 2、省级 3、地市)级 4、县级 5、乡镇级"),
		field.String("firm_name").Optional().Comment("行业名称采用国家标准GB/T 4754《国民经济行业分类》。 编码方法;采用层次码,第1位英文字母表示门类;第2、3位数字表示大类;第4位数字表示中类;第5位数字表示小类。 经营性国有资产、行政事业性国有资产、金融性国有资产和资源性国有资产的主体分类是否可以按照这个标准出。"),
		field.String("firm_code").Optional().Comment("行业代码代码	名称 A	农、林、牧、渔业 A01	农业 A011	种植业 …	…… A019	其他农业 …	…… B	采掘业 …	…… C	制造业 …	… C13	食品加工业 C131	粮食及饲料加工业 C1311	碾米业"),
		field.String("telephone_number").Optional().Comment("电话号码"),
		field.String("phone_number").Optional().Comment("手机号码"),
		field.String("fax_number").Optional().Comment("传真号码"),
		field.String("email_address").Optional().Comment("电子邮件"),
		field.String("province").Optional().Comment("省"),
		field.String("city").Optional().Comment("市"),
		field.String("county").Optional().Comment("区县"),
		field.String("street_address").Optional().Comment("街道地址"),
		field.String("postal_code").Optional().Comment("邮政编码"),
		field.Int64("fund_supply_mode_name").Optional().Comment("经费供给方式名称;代码	名称	说明 1	全额	指由财政供应全部经费的预算单位。 2	差额	指财政按单位收支差额供应经费的预算单位。 3	自收自支	不靠财政供应经费的预算单位。 9	其他	指财政按其他方式供应经费的预算单位。"),
		field.Int64("fund_supply_mode_code").Optional().Comment("经费供给方式代码;代码	名称	说明 1	全额	指由财政供应全部经费的预算单位。 2	差额	指财政按单位收支差额供应经费的预算单位。 3	自收自支	不靠财政供应经费的预算单位。 9	其他	指财政按其他方式供应经费的预算单位。"),
		field.Int64("administration_level_name").Optional().Comment("单位行政级别名称;代码	名称 1	正部省)级 2	副部省)级 3	正厅地)级 4	副厅地)级 5	正处县)级 6	副处县)级 7	正科级 8	副科级 9	股级"),
		field.Int64("administration_level_type").Optional().Comment("行政单位部门性质分类代码"),
		field.Int64("department_category_name").Optional().Comment("事业单位部门性质分类名称"),
		field.Int64("business_department_classification_code").Optional().Comment("事业单位部门性质分类代码;代码	名称 1	行政单位 11	共产党机关 12	政府机关 13	人大机关 14	政协机关 15	民主党派机关 16	群众团体 2	公检法司 21	公安 22	检察 23	法院 24	司法行政 25	监狱 26	劳教 29	其他 3	驻外机构 代码	名称 1	农、林、牧、渔业 2	交通运输、仓储和邮政业 3	科学研究、技术服务和地质勘查业 4	水利、环境和公共设施管理业 5	教育 6	卫生、社会保障和社会福利业 7	文化、体育和娱乐业 9	其他"),
		field.String("principal").Optional().Comment("负责人"),
		field.String("contact").Optional().Comment("联系人"),
		field.String("corporate_representative").Optional().Comment("法人代表"),
		field.Int64("economic_type").Optional().Comment("经济类型;10、国有及国有控股是否中外合资合作企业 11.是 12,否) 20、厂办大集体21 中央企厂办大集体 22、中央下放企业厂办大集体;23、地方企业厂办大集体) 30.其他城镇计提"),
		field.String("financial_affiliation").Optional().Comment("财务隶属关系行政隶属关系代码-部门标识代码)"),
		field.Int64("financial_unit").Optional().Comment("是否财政单位"),
		field.Int64("vertical_unit").Optional().Comment("是否垂管单位"),
		field.Int64("virtual_unit").Optional().Comment("是否虚拟单位"),
		field.Int64("accounting_system").Optional().Comment("单位执行会计制度:10.行政单位会计制度 20.事业单位会计制度 21.科学事业单位会计制度 22.中小学校会计制度 23.高等学校会计制度 24.医院会计制度 25.基层医疗卫生机构会计制度 26.测绘事业单位会计制度 27.地质勘查单位会计制度 28.彩票机构会计制度 30.民间非营利组织会计制度 50.其他"),
		field.String("corporate_tag").Optional().Comment("企业标志"),
		field.String("maintainer_mark").Optional().Comment("维修商标志"),
		field.String("supplier_mark").Optional().Comment("供应商标志"),
		field.String("manufacture_mark").Optional().Comment("制造商标志"),
		field.String("asset_disposal_agency_tag").Optional().Comment("资产处置代理机构标志"),
		field.String("state_asset_management_company").Optional().Comment("国有资产经营公司标志"),
		field.Int64("institution_number").Optional().Comment("事业编制人数"),
		field.Int64("authorized_number").Optional().Comment("行政编制人数"),
		field.Int64("of_side_number").Optional().Comment("编制外人数"),
		field.String("institution_authority").Optional().Comment("事业编制"),
		field.String("administration_authority").Optional().Comment("行政编制"),
		field.Time("enable_time").GoType(date.Now()).Optional().Comment("启用时间"),
		field.Time("disable_time").GoType(date.Now()).Optional().Comment("停用时间"),
		field.Time("seal_time").GoType(date.Now()).Optional().Comment("封存时间"),
		field.String("unit_remark").Optional().Comment("备注"),
		field.String("asset_download_status").Optional().Comment("资产数据下载状态"),
		field.String("latitude_and_longitude").Optional().Comment("GPS经纬度"),
		field.Int64("sort").Optional().Comment("排序"),
		field.String("organ_code").Optional().Comment("组织单位)编码"),
		field.Int64("had_inner_number").Optional().Comment("实有内设机构数量"),
		field.Int64("approve_inner_number").Optional().Comment("批准内设机构数量"),
		field.String("img").Optional().Comment("单位logo"),
		field.String("link_man").Optional().Comment("单位联系人"),
		field.String("link_phone").Optional().Comment("单位管理员联系方式"),
		field.Int64("tenant_id").StructTag(`json:"tenantId,string"`).Optional().Comment("租户id"),
		field.String("unit_code").Optional().Comment("单位编码"),
		field.Int64("authority_total").Optional().Comment("编制合计"),
		field.Int64("staff_on_active_duty").Optional().Comment("在职人员"),
		field.Int64("retired_staff").Optional().Comment("离退休人员"),
		field.Int64("other_staff").Optional().Comment("其它人员"),
		field.String("is_virtual").Default("0").Comment("实/虚单位"),
		field.Int64("unit_reform").Optional().Comment("事业单位改革分类10.行政类事业单位21.公益一类事业单位22.公益二类事业单位23.生产经营类事业单位90.暂未明确类别单位"),
		field.Int64("person_expenditure").Optional().Comment("基本支出人员经费功能分类"),
		field.String("institution_type").Optional().Comment("行业机构类型83 教育91 中国共产党机关92 国家机构94 社会保障95 群众团体、社会团体和其他成员组织96 基层群众自治组织及其他组织"),
		field.Int64("department_identification").Optional().Comment("部门标识"),
		field.Int64("total_headcount").Optional())
}

// Edges of the AsUnit.
func (AsUnit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenantx", AsTenant.Type).Ref("unit").Field("tenant_id").Unique(),
	}
}
