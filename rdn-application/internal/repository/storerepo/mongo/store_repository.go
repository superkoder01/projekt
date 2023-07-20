package mongo

import (
	"RDN-application/internal/repository/dao"
	"context"
)

type StoreRepository interface {
	BatchInsert(ctx context.Context, collection []interface{}) error
	FindAllFromDate(ctx context.Context, date string) ([]dao.HourDataDao, error)
	FindAllBetween(ctx context.Context, startDate, endDate string) ([]dao.HourDataDao, error)
}
