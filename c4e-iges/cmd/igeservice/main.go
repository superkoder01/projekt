package main

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/services/alarmservice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/factory"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/handlers/http"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/handlers/igeservice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/contractrepo"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/invoicerepo"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/publisher"
	"os"
	"os/signal"
	"syscall"
)

var (
	invoiceRepoFactory      ports.InvoiceRepoFactory
	contractRepoFactory     ports.ContractReposFactory
	invoicePublisherFactory ports.InvoicePublisherFactory
	alarmServiceFactory     ports.AlarmServiceFactory
	appLogger               logger.Logger
	appConfig               *config.AppConfig
)

func init() {
	// Initialize configuration
	cfg, err := config.LoadConfig()
	appConfig = cfg
	if err != nil {
		panic(fmt.Errorf("failed to load configuration: %s", err))
	}

	// Initialize logger
	appLogger = logger.NewLogger(appConfig)
	appLogger.InitLogger()

	initializeFactories()
}

func initializeFactories() {
	ctx := context.Background()

	// Contract repo factory
	contractRepoFactory = contractrepo.NewContractRepoFactory(ctx,
		contractrepo.ContractRepoType(appConfig.Ports.ContractRepoPort), appLogger, appConfig)

	// invoice repo factory
	invoiceRepoFactory = invoicerepo.NewInvoiceRepoFactory(ctx,
		invoicerepo.InvoiceRepoType(appConfig.Ports.InvoiceRepoPort), appLogger, appConfig)

	// Invoice publisher factory
	invoicePublisherFactory = publisher.NewPublisherFactory(ctx,
		publisher.PublisherType(appConfig.Ports.InvoicePublisherPort), appLogger, appConfig)

	// Alarm service factory
	alarmServiceFactory = alarmservice.NewAlarmServiceFactory(ctx,
		alarmservice.AlarmServiceType(appConfig.Ports.AlarmServicePort), appLogger, appConfig)
}

func main() {
	faktory := factory.NewFactoryBuilder().
		WithContractRepo(contractRepoFactory.MakeRepo()).
		WithInvoiceRepo(invoiceRepoFactory.MakeRepo()).
		WithInvoicePublisher(invoicePublisherFactory.MakePublisher()).
		WithAlarmService(alarmServiceFactory.MakeService()).Create()

	service := igeservice.New(context.Background(), faktory, appLogger, appConfig)
	registerShutdownHook(service)
	err := service.Start(false) // DryRun
	if err != nil {
		appLogger.Fatalf("%s-%s failed to start, reason: %v", appConfig.ServiceName, appConfig.ServiceVersion, err)
	}
}

func registerShutdownHook(service igeservice.Service) {
	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-osSignalChan
		service.Shutdown()
		os.Exit(0)
	}()
}

func StartStatusService(errChan chan<- error) {
	service := http.NewStatusHttpHandler(appLogger, appConfig)
	err := service.Start()
	if err != nil {
		appLogger.Fatalf("status service failed to start, reason: %v", err)
		errChan <- err
	}
	appLogger.Info("status service started")
}
