package svc

import (
	"orginone/app/market/cmd/rpc/internal/config"
	"orginone/app/market/cmd/rpc/internal/store"
)

type ServiceContext struct {
	Config      config.Config
	MarketStore store.MarketStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MarketStore: *store.NewMarketStore(c.Store),
	}
}
