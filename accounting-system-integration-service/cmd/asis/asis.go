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
