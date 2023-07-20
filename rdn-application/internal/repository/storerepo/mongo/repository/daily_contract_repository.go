package repository

import (
	"RDN-application/internal/repository/dao"
	"RDN-application/internal/repository/storerepo/mongo/client"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type StoreRepo struct {
	mongoClient    *client.Client
	collectionName string
	config         config.AppConfig
	logger         logger.Logger
}

func NewStoreRepo(config config.AppConfig, logger logger.Logger) *StoreRepo {
	contract := StoreRepo{
		mongoClient:    client.NewMongoClient(config),
		collectionName: config.GetStoreConfig().CollectionName,
		config:         config,
		logger:         logger,
	}

	contract.createCollection()
	return &contract
}

func (repo *StoreRepo) createCollection() {
	ctx, timeout := context.WithTimeout(context.Background(), time.Duration(repo.config.GetStoreConfig().Timeout)*time.Second)
	defer timeout()

	names, err := repo.mongoClient.GetConnection().ListCollectionNames(ctx, bson.D{})
	if err != nil {
		repo.logger.Fatalf("Cannot list collections in database: %v", err)
	}
	for _, name := range names {
		if name == repo.collectionName {
			return
		}
	}

	err = repo.mongoClient.GetConnection().
		CreateCollection(context.Background(), repo.collectionName,
			options.CreateCollection().SetTimeSeriesOptions(
				options.TimeSeries().
					SetTimeField("timestamp").
					SetGranularity("hours").
					SetMetaField("metadata")),
		)

	if err != nil {
		repo.logger.Fatalf("Failed to create database time-series collection !!")
	} else {
		repo.logger.Infof("Collection %v created !", repo.collectionName)
	}
}

func (repo *StoreRepo) BatchInsert(ctx context.Context, collection []interface{}) error {
	var dailyData []interface{}

	for _, element := range collection {
		dailyData = append(dailyData, element)
	}
	_, err := repo.mongoClient.GetConnection().Collection(repo.collectionName).InsertMany(ctx, dailyData)
	return err
}

func (repo *StoreRepo) FindAllFromDate(ctx context.Context, date string) ([]dao.HourDataDao, error) {
	return nil, nil
}

func (repo *StoreRepo) FindAllBetween(ctx context.Context, startDate, endDate string) ([]dao.HourDataDao, error) {
	return nil, nil
}
