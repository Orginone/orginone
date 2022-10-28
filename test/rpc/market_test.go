package test_rpc

import (
	"context"
	"fmt"
	"orginone/common/rpcs/market"
	"testing"

	"github.com/zeromicro/go-zero/zrpc"
)

func TestFindMarketApp(t *testing.T) {
	c, _ := zrpc.NewClientWithTarget("127.0.0.1:1022")
	client := market.NewMarket(c)
	unitRes, err := client.FindMarketApp(context.Background(), &market.MarketAppReq{
		AppName:    "资产",
		SaleStatus: []int64{1},
		Page: &market.PageRequest{
			Offset: 0,
			Limit:  20,
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(unitRes.Json)
	}
}
