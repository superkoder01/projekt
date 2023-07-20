package main

import (
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/api/server"
	initConfig "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/init"
)

var (
	logger = logging.MustGetLogger("main")
)

func main() {
	// Start HTTP server
	initConfig.InitConfig()
	server.NewHttpServer().Run()
}
