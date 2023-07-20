package server

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/email/smtp"
	"NotificationEmailService/internal/logger"
)

func NewByConfig(config *config.EmailServiceConfig) *Server {
	log := logger.NewApiLogger(&config.Logger)
	return New(config, log, smtp.New(config.Smtp, log))
}
