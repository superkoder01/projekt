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
