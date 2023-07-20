package main

import (
	"RDN-application/internal/ports"
	"RDN-application/internal/repository/storerepo/factory"
	"RDN-application/internal/server"
	"RDN-application/internal/services/data_collector"
	"RDN-application/internal/services/msgbroker"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"log"
)

var (
	dataService             ports.DataService
	dataServiceFactory      ports.DataServiceFactory
	notificationRepoFactory ports.NotificationRepoFactory
	collectorFactory        ports.CollectorFactory
)

func initFactory(cfg config.AppConfig, logger logger.Logger) {
	switch cfg.GetPortsConfig().StoreRepoType {
	case "mongo":
		dataServiceFactory = factory.NewMongoDataServiceFactory(cfg, logger)
	case "influx":
		dataServiceFactory = factory.NewInfluxDataStoreFactory(cfg, logger)
	default:
		panic("unknown store repo type: " + cfg.GetPortsConfig().StoreRepoType)
	}

	switch cfg.GetPortsConfig().NotificationServiceType {
	case "rabbitmq":
		notificationRepoFactory = msgbroker.NewRabbitmqFactory(cfg, logger)
	default:
		panic("unknown notification repo type: " + cfg.GetPortsConfig().NotificationServiceType)
	}

	switch cfg.GetPortsConfig().CollectorType {
	case "http":
		collectorFactory = data_collector.NewHttpCollectorFactory(cfg, logger)
	default:
		panic("unknown collector type: " + cfg.GetPortsConfig().CollectorType)
	}
	dataService = dataServiceFactory.MakeService()
}

func main() {
	log.Println("Starting rdn service")

	configPath := config.GetConfigPath()
	cfg, err := config.GetConfigurationLocal(configPath)

	if err != nil {
		log.Fatalf("Error while loading config file: %s error: %v", configPath, err)
	}

	appLogger := logger.NewApiLogger(cfg.GetLoggerConfig())
	appLogger.InitLogger()
	initFactory(cfg, appLogger)
	appLogger.Infof("%v is running ", cfg.GetServiceConfig().ServiceName)

	fetchServer := server.NewServer(
		cfg,
		dataService,
		notificationRepoFactory.MakeService(),
		collectorFactory.MakeService(),
		appLogger,
	)
	err = fetchServer.Run()

	if err != nil {
		appLogger.Fatalf("failed to start rdn-fetcher service %v", err)
	}
}
