package factory

import (
	"RDN-application/internal/ports"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
)

type mongoStoreFactory struct {
	cfg    config.AppConfig
	logger logger.Logger
}

func NewMongoDataServiceFactory(cfg config.AppConfig, logger logger.Logger) *mongoStoreFactory {
	return &mongoStoreFactory{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *mongoStoreFactory) MakeService() ports.DataService {
	//return mongo.NewDataService(s.cfg, s.logger)
	panic("service not implemented")
}
