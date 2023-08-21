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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/api/controller"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
)

const (
	basePath = "/api/signer"
)

var (
	//signer-service controllers
	signerController controller.Controller
)

type HTTPServer interface {
	Run()
}

type httpServer struct {
	cf     controller.ControllerFactory
	engine *gin.Engine
}

func NewHttpServer(cf controller.ControllerFactory) *httpServer {
	return &httpServer{cf: cf}
}

func (h *httpServer) Run() {
	h.engine = gin.New()
	h.engine.RouterGroup.Group(basePath)

	h.engine.Use(logRequest)
	h.engine.Use(attachHeaders)
	h.engine.Use(handleOptions)
	h.engine.Use(verifyRBAC)

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))
	h.initializeHandlers()
	h.setUpRoutes()
	h.engine.Run(":" + conf.GetListenPortConfig().Port)
}

func (h *httpServer) setUpRoutes() {
	g := h.engine.Group(basePath)

	g.POST("/initSign", signerController.InitSign)
	g.POST("/signingCompletedNotification", signerController.SigningCompletedNotification)

}

func (h *httpServer) initializeHandlers() {
	// signer-service controller
	signerController = h.cf.New(controller.SIGNER)
}
