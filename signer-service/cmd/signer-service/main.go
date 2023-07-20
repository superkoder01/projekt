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
