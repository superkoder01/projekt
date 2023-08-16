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