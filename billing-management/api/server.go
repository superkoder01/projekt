package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/handler"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/controller"
	rabbit_handler "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/rabbitmq/handler"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/configuration"
)

const (
	basePath = "/api/core"
)

var (
	// mysql handlers
	distributionNetworkOperatorHandler handler.Handler
	parameterNameHandler               handler.Handler
	tariffGroupOsdHandler              handler.Handler

	//mongodb controllers
	contractController         controller.Controller
	offerController            controller.Controller
	offerDraftController       controller.Controller
	invoiceController          controller.Controller
	pricingController          controller.Controller
	tariffGroupLabelController controller.Controller

	//email handler
	emailHandler rabbit_handler.EmailHandler
)

type HTTPServer interface {
	Run()
}

type httpServer struct {
	hf     handler.HandlerFactory
	cf     controller.ControllerFactory
	eh     rabbit_handler.EmailHandler
	engine *gin.Engine
}

func NewHttpServer(hf handler.HandlerFactory, cf controller.ControllerFactory, eh rabbit_handler.EmailHandler) *httpServer {
	return &httpServer{hf: hf, cf: cf, eh: eh}
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

	g.GET("/pricing", pricingController.List)
	g.GET("/pricing/:id", pricingController.FindOne)
	g.PUT("/pricing/:id", pricingController.Update)
	g.POST("/pricing", pricingController.Create)

	g.GET("/parameter_name", parameterNameHandler.List)

	g.GET("/tariff/group/osd", tariffGroupOsdHandler.List)
	g.GET("/tariff/group/osd/:id", tariffGroupOsdHandler.GetByID)
	g.POST("/tariff/group/osd", tariffGroupOsdHandler.Create)

	g.GET("/distribution_network_operator/:id", distributionNetworkOperatorHandler.GetByID)
	g.GET("/distribution_network_operator", distributionNetworkOperatorHandler.List)
	g.POST("/distribution_network_operator", distributionNetworkOperatorHandler.Create)
	g.PUT("/distribution_network_operator/:id", distributionNetworkOperatorHandler.UpdateByID)
	g.DELETE("/distribution_network_operator/:id", distributionNetworkOperatorHandler.DeleteByID)

	g.GET("/contract", contractController.List)
	g.POST("/contract", contractController.Create)
	g.PUT("/contract/:id", contractController.Update)
	g.GET("/contract/:id", contractController.FindOne)
	g.GET("/contract/customer/:id", contractController.List)

	g.GET("/offer", offerController.List)
	g.POST("/offer", offerController.Create)
	g.GET("/offer/:id", offerController.FindOne)
	g.PUT("/offer/:id", offerController.Update)
	g.GET("/offer/customer/:id", offerController.List)
	g.PUT("/offer/send/:id", emailHandler.SendEmailOffer)

	g.GET("/draft_offer", offerDraftController.List)
	g.GET("/draft_offer/:id", offerDraftController.FindOne)
	g.PUT("/draft_offer/:id", offerDraftController.Update)
	g.POST("/draft_offer", offerDraftController.Create)

	g.GET("/invoice", invoiceController.List)
	g.GET("/invoice/:id", invoiceController.FindOne)
	g.GET("/invoice/customer/:id", invoiceController.List)

	g.GET("/tariff_group_label", tariffGroupLabelController.List)

}

func (h *httpServer) initializeHandlers() {
	// mysql handlers
	distributionNetworkOperatorHandler = h.hf.New(handler.DISTRIBUTION_NETWORK_OPERATOR)
	parameterNameHandler = h.hf.New(handler.PARAMETER_NAME)
	tariffGroupOsdHandler = h.hf.New(handler.TARIFF_GROUP_OSD)

	// mongodb controller
	contractController = h.cf.New(controller.CONTRACT)
	offerController = h.cf.New(controller.OFFER)
	offerDraftController = h.cf.New(controller.OFFER_DRAFTS)
	invoiceController = h.cf.New(controller.INVOICE)
	pricingController = h.cf.New(controller.PRICING)
	tariffGroupLabelController = h.cf.New(controller.TARIFF_GROUP_LABEL)

	//email handler
	emailHandler = h.eh
}
