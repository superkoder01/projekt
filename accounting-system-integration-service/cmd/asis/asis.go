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
	"github.com/op/go-logging"

	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/api"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/server"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
	i "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/init"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/mongodb"
)

var (
	logger = logging.MustGetLogger("main")
)

func main() {
	// Initialize config etc, panic on error
	i.InitApp()

	//Initialize MongoDB database connection
	mongoConfig := configuration.GetMongoDatabaseConfig()
	mongoSession, err := mongodb.NewMongoSession(*mongoConfig)
	if err != nil {
		logger.Errorf("mongo database connection error: %s", err)
	}

	logger.Debug("Database connection initialization completed")

	// Create an instance of our handler which satisfies the generated interface
	invoiceService := api.NewInvoiceService(mongoSession)

	server.NewHttpServer(invoiceService).Run()
}
