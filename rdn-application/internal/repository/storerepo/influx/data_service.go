package influx

import (
	"RDN-application/internal/model"
	"RDN-application/internal/repository/dao"
	"RDN-application/internal/repository/storerepo/influx/entity"
	"RDN-application/internal/repository/storerepo/influx/repository"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"sort"
	"strconv"
	"time"
)

const DateFormat = "02-01-2006"

type dataService struct {
	config config.AppConfig
	logger logger.Logger
	store  StoreRepository
}

func NewDataService(config config.AppConfig, logger logger.Logger) *dataService {
	return &dataService{
		config: config,
		logger: logger,
		store:  repository.NewStoreRepo(config, logger),
	}
}

func (service *dataService) SaveDataToRepo(ctx context.Context, model model.Collection, time time.Time) error {
	date := time.Format(DateFormat)
	data, err := service.store.FindAllFromDate(ctx, date)
	if err != nil {
		service.logger.Errorf("Failed to connect to influx ! error: %v", err)
		return err
	}

	if data != nil && len(data.GetHourData()) < 1 {
		service.logger.Infof("Save data to repo {%v}", model)
		service.logger.Infof("Collect data date {%v}", time)

		err := service.store.BatchInsert(ctx, entity.MapDayCollectionDtoToDao(model), date)
		if err != nil {
			service.logger.Errorf("Failed to save data to influx ! error: %v", err)
			return err
		}
	} else {
		service.logger.Warnf("Data for date {%v} already exist in database !", time)
	}

	return nil
}

func (service *dataService) GetDataFromDate(ctx context.Context, time time.Time) (model.Model, error) {
	service.logger.Infof("GetDataFromDate: %v", time)
	daos, err := service.store.FindAllFromDate(ctx, time.Format(DateFormat))
	service.logger.Debugf("Collected data: %v", daos)

	if err != nil {
		service.logger.Errorf("Failed fetch data: %v", err)
		return nil, err
	} else {
		return service.createFixingResponse(model.FIXING1, daos), nil
	}
}

func (service *dataService) GetDataInBetween(ctx context.Context, startDate, endDate time.Time) ([]model.Model, error) {
	service.logger.Infof("GetDataInBetween %v and %v", startDate, endDate)
	daos, err := service.store.FindAllBetween(ctx, startDate.Format(DateFormat), endDate.Format(DateFormat))
	service.logger.Debugf("Collected data: %v", daos)

	if err != nil {
		service.logger.Errorf("Failed fetch data: %v", err)
		return nil, err
	} else {
		var response []model.Model
		for _, dataDao := range daos {
			response = append(response, service.createFixingResponse(model.FIXING1, dataDao))
		}
		return response, nil
	}
}

func (service *dataService) createFixingResponse(fixing model.FixingType, data dao.DailyDataDao) *model.Fixing {
	var hourArray []model.Model
	array := data.GetHourData()
	sort.SliceStable(array, func(i, j int) bool {
		intA, errA := strconv.Atoi(array[i].GetHour())
		intB, errB := strconv.Atoi(array[j].GetHour())
		if errA != nil || errB != nil {
			return false
		}
		return intA < intB
	})

	for _, hourData := range array {
		hourArray = append(hourArray, hourData)
	}

	return model.NewFixing(fixing, data.GetDate(), hourArray)
}
