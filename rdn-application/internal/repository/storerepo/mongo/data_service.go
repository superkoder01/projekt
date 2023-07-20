package mongo

import (
	"RDN-application/internal/model"
	"RDN-application/internal/repository/dao"
	"RDN-application/internal/repository/storerepo/mongo/repository"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"time"
)

type dataService struct {
	config config.AppConfig
	logger logger.Logger
	store  StoreRepository
}

func NewDataService(config config.AppConfig, logger logger.Logger) *dataService {
	return &dataService{
		store:  repository.NewStoreRepo(config, logger),
		config: config,
		logger: logger,
	}
}

func (service *dataService) SaveDataToRepo(ctx context.Context, model model.Collection, time time.Time) error {
	service.logger.Infof("Save data to repo {%v}", model)
	service.logger.Infof("Collected data date {%v}", time)
	//err := service.store.BatchInsert(ctx, model.GetCollectionData())
	//if err != nil {
	//	service.logger.Errorf("Save error ! %v", err)
	//	return err
	//} else {
	return nil
	//}
}

func (service *dataService) GetDataFromDate(ctx context.Context, time time.Time) (dao.DailyDataDao, error) {
	return nil, nil
}

func (service *dataService) GetDataInBetween(ctx context.Context, startDate, endDate time.Time) error {
	return nil
}
