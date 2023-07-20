package influx

import (
	"RDN-application/internal/repository/dao"
	"context"
)

type StoreRepository interface {
	BatchInsert(ctx context.Context, collection []dao.HourDataDao, date string) error
	FindAllFromDate(ctx context.Context, date string) (dao.DailyDataDao, error)
	FindAllBetween(ctx context.Context, startDate, endDate string) ([]dao.DailyDataDao, error)
}
