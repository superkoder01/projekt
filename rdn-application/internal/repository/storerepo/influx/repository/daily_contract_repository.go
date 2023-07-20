package repository

import (
	"RDN-application/internal/repository/dao"
	"RDN-application/internal/repository/storerepo/influx/client"
	"RDN-application/internal/repository/storerepo/influx/entity"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"strings"
	"time"
)

const DateFormat = "02-01-2006"

type storeRepo struct {
	influxClient   *client.Client
	collectionName string
	bucketName     string
	orgName        string
	config         config.AppConfig
	logger         logger.Logger
}

func NewStoreRepo(config config.AppConfig, logger logger.Logger) *storeRepo {
	store := storeRepo{
		influxClient:   client.NewInfluxClient(config),
		collectionName: config.GetStoreConfig().CollectionName,
		config:         config,
		logger:         logger,
		bucketName:     config.GetStoreConfig().BucketName,
		orgName:        config.GetStoreConfig().DbName,
	}
	return &store
}

func (repo *storeRepo) BatchInsert(ctx context.Context, collection []dao.HourDataDao, date string) error {
	var points []*write.Point
	saveDate := time.Now()

	for _, i := range collection {
		measurement := influxdb2.NewPointWithMeasurement(repo.orgName)
		points = append(points, measurement.
			AddTag("date", date).
			AddTag("hour", i.GetHour()).
			AddField("value", i.GetValue()).
			SetTime(saveDate))
	}
	writeAPI := repo.influxClient.WriteAPIBlocking(repo.orgName, repo.bucketName)
	defer repo.influxClient.Close()

	err := writeAPI.WritePoint(
		ctx, points...,
	)

	if err != nil {
		repo.logger.Errorf("Failed to save data %v", err)
		return err
	}
	repo.logger.Infof("Data saved")
	return nil
}

func (repo *storeRepo) FindAllFromDate(ctx context.Context, date string) (dao.DailyDataDao, error) {
	queryAPI := repo.influxClient.QueryAPI(repo.orgName)
	var daos []entity.HourData
	query := `from(bucket: "` + repo.bucketName + `")
				|> range(start: -1y)
				|> filter(fn: (r) => r["date"] == "` + date + `")
				|> mean()`
	repo.logger.Debugf("Executing query \"%v\"", query)
	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		repo.logger.Errorf("Failed to execute query ! %v", query)
		return nil, err
	}

	for result.Next() {
		value, ok := result.Record().Value().(float64)
		if !ok {
			return nil, errors.New("invalid data format in database (value: value-float64)")
		}
		hour, ok := result.Record().ValueByKey("hour").(string)
		hour = strings.Split(hour, "-")[0]
		if !ok {
			return nil, errors.New("invalid data format in database (value: hour-string)")
		}

		daos = append(daos, entity.HourData{Value: value, Hour: hour})
	}

	return &entity.DailyData{
		Date:  date,
		Hours: daos,
	}, nil
}

func (repo *storeRepo) FindAllBetween(ctx context.Context, startDate, endDate string) ([]dao.DailyDataDao, error) {
	queryAPI := repo.influxClient.QueryAPI(repo.orgName)
	var daos []dao.DailyDataDao
	var dateQueryBuilder strings.Builder

	start, sErr := Date(startDate)
	end, eErr := Date(endDate)

	if sErr != nil {
		return nil, sErr
	} else if eErr != nil {
		return nil, eErr
	}
	numberOfDaysInBetween := end.Sub(start) / (24 * time.Hour)
	currentDate := start

	for i := 0; i < int(numberOfDaysInBetween); i++ {
		dateQueryBuilder.WriteString("(" + currentDate.Format(DateFormat) + ")|")
		currentDate = currentDate.Add(time.Hour * 24)
	}
	//For last day without |
	dateQueryBuilder.WriteString("(" + currentDate.Format(DateFormat) + ")")

	query := `from(bucket: "` + repo.bucketName + `")
				|> range(start: -1y)
				|> filter(fn: (r) => r["date"] =~ /(` + dateQueryBuilder.String() + `)$/)
				|> group(columns: ["date"], mode:"by")
				|> yield(name: "mean")`
	repo.logger.Debugf("Executing query \"%v\"", query)
	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		repo.logger.Errorf("Failed to execute query ! %v", query)
		return nil, err
	}

	for result.Next() {
		date, ok := result.Record().ValueByKey("date").(string)
		if !ok {
			return nil, errors.New("invalid data format in database (value: date-string)")
		}
		array := findDailyDataInArray(daos, date)

		value, ok := result.Record().Value().(float64)
		if !ok {
			return nil, errors.New("invalid data format in database (value: value-float64)")
		}
		hour, ok := result.Record().ValueByKey("hour").(string)
		hour = strings.Split(hour, "-")[0]
		if !ok {
			return nil, errors.New("invalid data format in database (value: hour-string)")
		}

		if array == nil {
			daos = append(daos, &entity.DailyData{
				Date: date, Hours: []entity.HourData{{Value: value, Hour: hour}},
			})
		} else {
			ar := *array
			ar.AppendHourData(entity.HourData{Hour: hour, Value: value})
		}
	}
	return daos, nil
}

func findDailyDataInArray(array []dao.DailyDataDao, date string) *dao.DailyDataDao {
	for _, v := range array {
		if v.GetDate() == date {
			return &v
		}
	}
	return nil
}

func Date(date string) (time.Time, error) {
	return time.Parse(DateFormat, date)
}
