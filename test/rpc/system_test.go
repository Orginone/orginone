package test_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
)

func TestFindUserByUserId(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021", zrpc.WithTimeout(time.Minute))
	client := system.NewSystem(c)
	rpcres, err := client.FindUserByUserId(context.Background(), &system.PrimaryKeyReq{Id: 1})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(rpcres.Json)
	}
}

func TestUserInfoByUserId(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021", zrpc.WithTimeout(time.Minute))
	client := system.NewSystem(c)
	rpcres, err := client.FindPersonByUserId(context.Background(), &system.PersonByUserReq{UserId: 1})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(rpcres.Json)
	}
}

func TestFindTenantByAccount(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.FindTenantByAccount(context.TODO(), &system.ByAccountReq{Account: "18888888888"})
	if err != nil {
		t.Error(err)
	} else {
		var tenantEntitys []*schema.AsTenant
		err = json.Unmarshal([]byte(rpcres.Json), &tenantEntitys)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(rpcres.Json)
		}
	}
}

func TestSearchTeantByName(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.SearchTeantByName(context.TODO(), &system.SearchTeantReq{TenantName: "开放"})
	if err != nil {
		t.Error(err)
	} else {
		var tenantEntitys []*schema.AsTenant
		err = json.Unmarshal([]byte(rpcres.Json), &tenantEntitys)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(rpcres.Json)
		}
	}
}

func TestJoinTeantByCodeSystem(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	_, err := client.JoinTeantByCode(context.TODO(), &system.TenantCodeReq{
		UserId:     1,
		TenantCode: "00000000000000000000",
		Account:    "18888888888",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestScenesCreateTenant(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.CreateTenant(context.TODO(), &system.CreateTenantReq{UserId: 1, SocialCreCode: "91330106MA7BX78XX00", TenantName: "杭州快字科技有限公司00"})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(rpcres.Json)
	}
}

func TestFindAllGroupByCodeAndType(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.FindAllGroupByCodeAndType(context.TODO(), &system.SearchAllGroupReq{TenantCode: "ZH200903021124xS76I2", GroupType: 1})
	if err != nil {
		t.Error(err)
	} else {
		var tenantEntitys []*schema.AsAllGroup
		err = json.Unmarshal([]byte(rpcres.Json), &tenantEntitys)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(rpcres.Json)
		}
	}
}

func TestFindRoleMenusByUser(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.FindRoleMenusByUser(context.TODO(), &system.RoleMenusByUserReq{UserId: 1, PlatformId: 1, RoleIds: nil, TenantCode: "00000"})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(rpcres.Json)
	}
}

func TestFindMarketAppRoleMenusByUser(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.FindMarketAppRoleMenusByUser(context.TODO(), &system.RoleMenusByUserReq{UserId: 1, PlatformId: 1, RoleIds: nil, TenantCode: "000000"})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(rpcres.Json)
	}
}

func TestUpdateUnitInfo(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	rpcres, err := client.FindTenantById(context.Background(), &system.TenantByIdReq{Id: 1301341017987850241})
	if err != nil {
		t.Error(err)
	}
	var tenantEntity *schema.AsTenant
	err = json.Unmarshal([]byte(rpcres.Json), &tenantEntity)
	if err != nil {
		t.Error(err)
	}
	tenantEntity.Edges.Unit.UnitNameEn += "123456"
	buf, err := json.Marshal(tenantEntity.Edges.Unit)
	unitRes, err := client.UpdateUnitInfo(context.Background(), &system.CommonRpcReq{Json: string(buf)})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}
func TestFindInnerAgencyTreeByTenantCode(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	unitRes, err := client.FindInnerAgencyTreeByTenantCode(context.Background(), &system.TenantCodeReq{
		TenantCode: "ZH200903024648VE2nU8",
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}

func TestFindAgencyDeptCodeByTenantCode(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	unitRes, err := client.FindAgencyDeptCodeByTenantCode(context.Background(), &system.TenantCodeReq{
		TenantCode: "ZH200903020731BMijol",
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}
func TestSearchPersonByTenantCode(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	unitRes, err := client.FindTenantPersons(context.Background(), &system.SearchPersonByTenantCodeReq{
		TenantCode: "ZH200903020731BMijol",
		IsActivate: 3,
		Flag:       1,
		Page: &system.PageRequest{
			Offset:      0,
			Limit:       10,
			Filter:      "",
			SearchCount: true,
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}

func TestFindAllGroupByTentantCode(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	unitRes, err := client.FindAllGroupByTentantCode(context.Background(), &system.FindAllGroupByTenantCodeReq{
		TenantCode: "000000",
		Type:       1,
		Page: &system.PageRequest{
			Offset:      10,
			Limit:       10,
			Filter:      "",
			SearchCount: true,
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}

func TestFindGroupLayerById(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1021")
	client := system.NewSystem(c)
	unitRes, err := client.FindGroupLayerById(context.Background(), &system.GroupLayerByIdReq{
		GroupId:       1405434611979104257,
		SourceGroupId: 1405434611979104257,
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}
