package common

import (
	"fmt"
	"math"
	"orginone/common"
	"orginone/common/excel"
	"orginone/common/schema"
	"orginone/common/tools"
	"orginone/common/tools/date"
	"os"
	"strings"
	"testing"
)

func TestGenTenantCode(t *testing.T) {
	fmt.Println(tools.GenTenantCode())
}

func TestGenPassword(t *testing.T) {
	pwd, err := tools.BcryptEncryptPwd("123456")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(pwd)
	}
}

func TestSnake(t *testing.T) {
	fmt.Println(tools.Snake("Username"))
	fmt.Println(tools.Snake("userName"))
	fmt.Println(tools.Snake("ID"))
}

func TestWritePersonsToFile(t *testing.T) {
	excel.WritePersonsToFile(make([]*schema.AsPerson, 0))
}

func TestReadPeoplesFromFile(t *testing.T) {
	file, _ := os.Open("/opt/人员导出.xlsx")
	excel.ReadPersonsFromFile(file)
}

func TestIsPhone(t *testing.T) {
	ok := tools.IsPhone("18868878263")
	fmt.Println(ok)
}

func TestDateTime(t *testing.T) {
	t.Log(date.Now())
}

func TestGenKey(t *testing.T) {
	fmt.Println(tools.GenAppKeyOrSecret(50))
	fmt.Println(strings.ToUpper(tools.GenAppKeyOrSecret(128)))
}

func TestGenToken(t *testing.T) {
	fmt.Println(strings.Split("", ","))
	fmt.Println(common.GenJwtToken("1234qwer", int64(math.MaxInt32), map[string]interface{}{
		"userId":     1,
		"account":    "202005201314",
		"tenantCode": "000000",
	}))
}
