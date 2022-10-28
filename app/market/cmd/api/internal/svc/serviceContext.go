package svc

import (
	"orginone/app/market/cmd/api/internal/config"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	MarketRpc market.Market
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MarketRpc: market.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
	}
}
