package http

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/services"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"net/http"
	"strconv"
)

type statusHTTPHandler struct {
	log     logger.Logger
	cfg     *config.AppConfig
	status  ports.Status
	started bool
}

func NewStatusHttpHandler(log logger.Logger, cfg *config.AppConfig) *statusHTTPHandler {
	return &statusHTTPHandler{
		log:     log,
		cfg:     cfg,
		status:  services.NewStatusService(log, cfg),
		started: false,
	}
}

func (h *statusHTTPHandler) rootHandler(e *echo.Echo, c echo.Context) error {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, data)
}

func (h *statusHTTPHandler) checkStatus(c echo.Context) error {
	ok, err := h.status.IsAlive()
	if err != nil || ok == false {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "NOK"})
	}

	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func (h *statusHTTPHandler) Start() error {
	if h.started {
		return nil
	}

	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return h.rootHandler(e, c)
	}).Name = "hello"
	e.GET("/ping", h.checkStatus).Name = "ping"

	httpPort := strconv.Itoa(h.cfg.StatusService.HttpPort)
	if httpPort == "" {
		httpPort = strconv.Itoa(8080)
	}

	return e.Start(":" + httpPort)
}
