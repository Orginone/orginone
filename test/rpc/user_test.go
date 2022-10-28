package test_rpc

import (
	"context"
	"fmt"
	"orginone/common/rpcs/user"
	"testing"

	"github.com/zeromicro/go-zero/zrpc"
)

func TestUserReReg(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1020")
	client := user.NewUser(c)
	_, err := client.UserReReg(context.TODO(), &user.UserReRegReq{PhoneNumber: "18072943461", RealName: "realName", Pwd: "123456"})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("success")
	}
}
func TestJoinTeantByCode(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1020")
	client := user.NewUser(c)
	_, err := client.JoinTeantByCode(context.TODO(), &user.TenantCodeReq{
		UserId:     1,
		TenantCode: "00000000000000000000",
		Account:    "18888888888",
	})
	if err != nil {
		t.Error(err)
	}
}
