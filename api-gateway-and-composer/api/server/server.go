package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/proxy"
	"os"
)

type HTTPServer interface {
	Run()
}

type httpServer struct {
	engine *gin.Engine
}

func NewHttpServer() *httpServer {
	return &httpServer{}
}

func (h *httpServer) Run() {
	h.engine = gin.New()

	h.engine.Use(logRequest)
	h.engine.Use(attachHeaders)
	h.engine.Use(handleOptions)
	h.engine.Use(verifySignature)
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))
	h.engine.Use(gin.Recovery())
	//register proxies from the TOML file
	//configPath := "configs/proxy.yaml"
	err := proxy.RegisterProxies(h.engine)
	if err != nil {
		fmt.Printf("Mice -[ERR]::Error registering the proxies from toml: %v\n", err)
		os.Exit(1)
	}
	h.engine.Run(":" + conf.GetListenPortConfig().Port)
}
