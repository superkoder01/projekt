package data_collector

import (
	"RDN-application/internal/ports"
	"RDN-application/internal/services/data_collector/http_scrapper"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
)

type collectorFactory struct {
	cfg    config.AppConfig
	logger logger.Logger
}

func NewHttpCollectorFactory(cfg config.AppConfig, logger logger.Logger) *collectorFactory {
	return &collectorFactory{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *collectorFactory) MakeService() ports.DataCollector {
	return http_scrapper.NewDataCollector(s.cfg, s.logger)
}
