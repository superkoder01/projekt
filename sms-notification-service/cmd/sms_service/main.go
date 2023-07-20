package main

import (
	"NotificationSmsService/internal/server"
	"NotificationSmsService/pkg/config"
	"NotificationSmsService/pkg/logger"
	"log"
	"os"
)

func main() {
	log.Println("Starting sms service")

	configPath := os.Getenv("config")
	cfg, err := config.GetSmsConfiguration(configPath)
	if err != nil {
		log.Fatalf("Error while loading config file: %s error: %v", configPath, err)
	}
	appLogger := logger.NewApiLogger(&cfg.Logger)
	appLogger.InitLogger()

	appLogger.Infof("%v is running ", cfg.Service.ServiceName)
	appLogger.Infof("Loaded configuration: %v", cfg.Service.ServiceName)

	s := server.NewSmsServer(cfg, appLogger)

	serverErr := s.RunSmsServer()
	if serverErr != nil {
		appLogger.Fatalf("Server failed to start")
	}
}
