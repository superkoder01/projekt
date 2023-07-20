package main

import (
	"ConversionService/config"
	"ConversionService/internal/server"
	"ConversionService/pkg/logger"
	"log"
	"os"
)

func main() {
	configPath := os.Getenv("config")
	cfg, err := config.GetConfiguration(configPath)
	if err != nil {
		log.Fatalf("Loading config error: %v", err)
	}
	appLogger := logger.NewApiLogger(&cfg.Logger)
	appLogger.InitLogger()

	appLogger.Infof("%v is running ", cfg.Service.ServiceName)
	appLogger.Infof("Loaded configuration: %v", cfg.Service.ServiceName)

	s := server.NewServer(cfg, appLogger)

	serverErr := s.Run()
	if serverErr != nil {
		appLogger.Fatalf("Server failed to start")
	}
}
