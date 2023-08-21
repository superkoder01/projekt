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
package main

import (
	server "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/api"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/api/controller"
	i "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/init"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service"
)

func main() {
	// Initialize config etc, panic on error
	i.InitApp()

	// Initialize Service factory

	apiLogger := logger.NewApiLogger(configuration.GetLoggerConfig())

	msf := service.NewServiceFactory(apiLogger)

	// Initialize Handler factory
	mcf := controller.NewControllerFactory(msf)

	apiLogger.Debug("Application initialization completed")

	// Start HTTP server
	server.NewHttpServer(mcf).Run()
}
