package svc

import (
	"orginone/app/user/cmd/rpc/internal/config"
	"orginone/app/user/cmd/rpc/internal/store"
)

type ServiceContext struct {
	Config    config.Config
	UserStore store.UserStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserStore: *store.NewUserStore(c.Store),
	}
}
