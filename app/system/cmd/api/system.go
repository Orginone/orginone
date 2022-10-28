package main

import (
	"flag"
	"fmt"
	"os"

	"orginone/app/system/cmd/api/internal/config"
	"orginone/app/system/cmd/api/internal/handler"
	"orginone/app/system/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/system-api.yaml", "the config file")

func main() {
	flag.Parse()
	os.Mkdir("./Image/", 002)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
