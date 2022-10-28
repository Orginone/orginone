package svc

import (
	"orginone/app/system/cmd/api/internal/config"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SystemRpc system.System
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SystemRpc: system.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
	}
}
