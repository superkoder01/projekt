package msgbroker

import (
	"RDN-application/internal/ports"
	"RDN-application/internal/services/msgbroker/producer"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
)

type producerFactory struct {
	cfg    config.AppConfig
	logger logger.Logger
}

func NewRabbitmqFactory(cfg config.AppConfig, logger logger.Logger) *producerFactory {
	return &producerFactory{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *producerFactory) MakeService() ports.NotificationRepo {
	return producer.NewNotificationService(s.cfg, s.logger)
}
