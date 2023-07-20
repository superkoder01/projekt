package main

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/server"
	"log"
	"os"
)

func main() {

	configPath := os.Getenv("config")
	cfg, err := config.GetEmailConfiguration(configPath)
	if err != nil {
		log.Fatalf("Error while loading config file: %s error: %v", configPath, err)
	}

	if err := server.NewByConfig(cfg).RunEmailServer(); err != nil {
		panic(err)
	}
}
