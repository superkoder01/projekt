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
