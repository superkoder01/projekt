package factory

import (
	"RDN-application/internal/ports"
	"RDN-application/internal/repository/storerepo/influx"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
)

type influxStoreFactory struct {
	cfg    config.AppConfig
	logger logger.Logger
}

func NewInfluxDataStoreFactory(cfg config.AppConfig, logger logger.Logger) *influxStoreFactory {
	return &influxStoreFactory{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *influxStoreFactory) MakeService() ports.DataService {
	return influx.NewDataService(s.cfg, s.logger)
}
