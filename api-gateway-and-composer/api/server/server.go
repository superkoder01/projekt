/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
