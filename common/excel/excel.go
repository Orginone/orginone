package excel

import (
	"fmt"
	"io"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"
	"orginone/common/tools/date"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"gopkg.in/oauth2.v3/utils/uuid"
)

//将人员信息写入excel文件
func WritePersonsToFile(persons []*schema.AsPerson) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]string{"姓名", "身份证", "性别", "生日", "邮箱", "手机号", "人员编码"})
	for index, person := range persons {
		gender := ""
		switch person.Gender {
		case 1:
			gender = "男"
			break
		case 2:
			gender = "女"
			break
		}
		if excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2),
			&[]interface{}{person.RealName, person.IDCard, gender, person.UserBirthday,
				person.UserEmail, person.PhoneNumber, person.UserCode}) != nil {
			return nil, err
		}
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取人员信息
func ReadPersonsFromFile(file io.Reader) ([]*schema.AsPerson, error) {
	persons := make([]*schema.AsPerson, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return persons, err
	}
	phones := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "", "", "", "", "", "")
		if index != 0 {
			person := new(schema.AsPerson)
			person.RealName = row[0]
			person.IDCard = row[1]
			if !tools.IsNilOrEmpty(person.IDCard) &&
				!tools.IsIDCard(person.IDCard) {
				continue
			}
			switch row[2] {
			case "男":
				person.Gender = 1
				break
			case "女":
				person.Gender = 2
				break
			default:
				person.Gender = 0
				break
			}
			brithday, err := time.Parse("2006-01-02", row[3])
			if err != nil {
				person.UserBirthday = date.DateTime(brithday)
			}
			person.UserEmail = row[4]
			person.PhoneNumber = row[5]
			person.UserCode = row[6]
			if tools.IsPhone(person.PhoneNumber) {
				isExist := false
				for _, phone := range phones {
					if person.PhoneNumber == phone {
						isExist = true
						break
					}
				}
				if !isExist {
					persons = append(persons, person)
					phones = append(phones, row[5])
				}
			}
		}
	}
	return persons, err
}

//将单位信息写入excel文件
func WriteUnitToFile(units []*schema.AsUnit) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"单位名称", "统一社会信用代码"})
	for index, unit := range units {
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{unit.UnitName, unit.SocialCreditCode})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//将单位信息写入excel文件
func WriteMasterUnitToFile(units []*schema.AsUnit) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"单位名称", "姓名", "手机号"})
	for index, unit := range units {
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{unit.UnitName, unit.LinkMan, unit.LinkPhone})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取人员信息
func ReadMasterUnitFromFile(file io.Reader) ([]*schema.AsUnit, error) {
	units := make([]*schema.AsUnit, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return units, err
	}
	phones := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "", "", "", "", "", "")
		if index != 0 {
			unit := new(schema.AsUnit)
			unit.UnitName = row[0]
			unit.LinkMan = row[1]
			unit.LinkPhone = row[2]
			if tools.IsPhone(unit.LinkPhone) {
				isExist := false
				for _, phone := range phones {
					if unit.LinkPhone == phone {
						isExist = true
						break
					}
				}
				if !isExist {
					units = append(units, unit)
					phones = append(phones, row[2])
				}
			}
		}
	}
	return units, err
}

//将内部机构写入excel文件
func WriteAgencyToFile(agencys []*schema.AsInnerAgency) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"部门名称", "上级部门名称"})
	for index, agency := range agencys {
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{agency.AgencyName, agency.Edges.Parent.AgencyName})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取人员信息
func ReadAgencyFromFile(file io.Reader) ([]*model.AgencyData, error) {
	agencys := make([]*model.AgencyData, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return agencys, err
	}
	names := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "")
		if index != 0 {
			isExist := false
			for _, name := range names {
				if row[0] == name {
					isExist = true
					break
				}
			}
			if !isExist {
				names = append(names, row[0])
				agencys = append(agencys, &model.AgencyData{
					AgencyName:       row[0],
					ParentAgencyName: row[1],
				})
			}
		}
	}
	return agencys, err
}

//将集团信息写入excel文件
func WriteGroupRelationToFile(groups []*model.GroupRelationData) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"集团名称", "管理单位统一社会信用代码", "上级集团名称"})
	for index, group := range groups {
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{group.GroupName, group.SocialCreditCode, group.ParentGroupName})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取租户集团关系信息
func ReadGroupRelationFromFile(file io.Reader) ([]*model.GroupRelationData, error) {
	groupRelationDatas := make([]*model.GroupRelationData, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return groupRelationDatas, err
	}
	groupNames := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "", "")
		if index != 0 {
			isExist := false
			if !tools.IsNilOrEmpty(row[0]) &&
				!tools.IsNilOrEmpty(row[2]) {
				for _, groupName := range groupNames {
					if row[0] == groupName {
						isExist = true
						break
					}
				}
				if !isExist {
					groupNames = append(groupNames, row[0])
					groupRelationDatas = append(groupRelationDatas, &model.GroupRelationData{
						GroupName:        row[0],
						SocialCreditCode: row[1],
						ParentGroupName:  row[2],
					})
				}
			}
		}
	}
	return groupRelationDatas, err
}

//将岗位信息写入excel文件
func WriteJobToFile(jobs []*schema.AsJob) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"岗位名称"})
	for index, job := range jobs {
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{job.JobName})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取人员信息
func ReadJobFromFile(file io.Reader) ([]*schema.AsJob, error) {
	jobs := make([]*schema.AsJob, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return jobs, err
	}
	names := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "", "", "", "", "", "")
		if index != 0 {
			isExist := false
			for _, name := range names {
				if row[0] == name {
					isExist = true
					break
				}
			}
			if !isExist {
				job := new(schema.AsJob)
				job.JobName = row[0]
				jobs = append(jobs, job)
				names = append(names, row[0])
			}
		}
	}
	return jobs, err
}

//从excel文件中读取租户集团关系信息
func ReadTenantRelationFromFile(file io.Reader) ([]*system.TenantRelationData, error) {
	tenantRelationDatas := make([]*system.TenantRelationData, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return tenantRelationDatas, err
	}
	socialCreditCodes := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "")
		if index != 0 {
			isExist := false
			if tools.IsSocialCreditCode(row[0]) && !tools.IsNilOrEmpty(row[1]) {
				for _, socialCreditCode := range socialCreditCodes {
					if row[0] == socialCreditCode {
						isExist = true
						break
					}
				}
				if !isExist {
					socialCreditCodes = append(socialCreditCodes, row[0])
					tenantRelationDatas = append(tenantRelationDatas, &system.TenantRelationData{
						SocialCreditCode: row[0],
						ParentGroupName:  row[1],
					})
				}
			}
		}
	}
	return tenantRelationDatas, err
}

//将内部机构写入excel文件
func WriteTenantToFile(tenants []*schema.AsTenant) ([]byte, error) {
	excel := excelize.NewFile()
	sheetName := excel.GetSheetName(0)
	err := excel.SetSheetRow(sheetName, "A1", &[]interface{}{"租户名称", "统一社会信用代码", "管理员姓名", "手机号", "行政区划代码", "行政区划名称", "实/虚单位", "创建时间"})
	for index, tenant := range tenants {
		virtual := "实单位"
		if tenant.IsVirtual == 1 {
			virtual = "虚单位"
		}
		err = excel.SetSheetRow(sheetName, fmt.Sprintf("A%d", index+2), &[]string{tenant.TenantName, tenant.Edges.Unit.SocialCreditCode, tenant.Edges.Unit.LinkMan, tenant.Edges.Unit.LinkPhone, tenant.Edges.Unit.AdministrationDivisionCode, tenant.Edges.Unit.AdministrationDivisionName, virtual, tenant.CreateTime.String()})
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	path := "/opt/" + id.String() + ".xlsx"
	err = excel.SaveAs(path)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)
	return os.ReadFile(path)
}

//从excel文件中读取租户集团关系信息
func ReadTenantFromFile(file io.Reader) ([]*schema.AsTenant, error) {
	tenantDatas := make([]*schema.AsTenant, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return tenantDatas, err
	}
	socialCreditCodes := make([]string, 0)
	tenantNames := make([]string, 0)
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	for index, row := range rows {
		row := append(row, "", "", "", "", "", "", "", "")
		if index != 0 {
			isExist := false
			if (tools.IsNilOrEmpty(row[1]) || tools.IsSocialCreditCode(row[1])) &&
				!tools.IsNilOrEmpty(row[0]) &&
				!tools.IsNilOrEmpty(row[2]) &&
				!tools.IsNilOrEmpty(row[3]) {
				var virtual int64 = 1
				switch row[6] {
				case "实单位":
					virtual = 0
					break
				default:
					break
				}
				for _, tenantName := range tenantNames {
					if row[0] == tenantName {
						isExist = true
						break
					}
				}
				for _, socialCreditCode := range socialCreditCodes {

					if row[1] == socialCreditCode {
						isExist = true
						break
					}
				}
				if !isExist {
					socialCreditCodes = append(socialCreditCodes, row[0])
					tenantNames = append(tenantNames, row[0])
					tenant := &schema.AsTenant{
						TenantName: row[0],
						IsVirtual:  virtual,
						Edges:      schema.AsTenantEdges{},
					}
					tenant.Edges.Unit = &schema.AsUnit{
						UnitName:                   row[0],
						SocialCreditCode:           row[1],
						LinkMan:                    row[2],
						LinkPhone:                  row[3],
						IsVirtual:                  strconv.FormatInt(virtual, 64),
						AdministrationDivisionCode: row[4],
						AdministrationDivisionName: row[5],
					}
					tenantDatas = append(tenantDatas, tenant)
				}
			}
		}
	}
	return tenantDatas, err
}
