package server

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/labstack/echo/v4"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/api"
)

var (
	logger = logging.MustGetLogger("http_server")
)

type HTTPServer interface {
	Run()
}

type httpServer struct {
	is *api.InvoiceService
}

func NewHttpServer(is *api.InvoiceService) *httpServer {
	return &httpServer{is}
}

func (h *httpServer) Run() {
	port, baseUri := configuration.GetHttpConfigData()

	logger.Debug("initializing HTTP server ...")
	// This is how you set up a basic Echo router
	e := echo.New()

	e.Use(logRequest)
	e.Use(attachHeaders)
	e.Use(verifyRBAC)

	logger.Debug("register endpoints: ", baseUri)
	api.RegisterHandlersWithBaseURL(e, h.is, baseUri)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", port)))

	logger.Debug("HTTP server started")
}