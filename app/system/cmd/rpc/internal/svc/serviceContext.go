package svc

import (
	"orginone/app/system/cmd/rpc/internal/config"
	"orginone/app/system/cmd/rpc/internal/store"
)

type ServiceContext struct {
	Config      config.Config
	SystemStore store.SystemStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		SystemStore: *store.NewSystemStore(c.Store),
	}
}
