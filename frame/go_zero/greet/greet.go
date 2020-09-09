package main

import (
	"flag"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"golearn/frame/go_zero/greet/internal/config"
	"golearn/frame/go_zero/greet/internal/handler"
	"golearn/frame/go_zero/greet/internal/svc"
)

var configFile = flag.String("f", "etc/greet-api.json", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	server.Start()
}
