package ports

import (
	"RDN-application/internal/model"
	"context"
	"time"
)

type DataService interface {
	SaveDataToRepo(context.Context, model.Collection, time.Time) error
	GetDataFromDate(ctx context.Context, time time.Time) (model.Model, error)
	GetDataInBetween(ctx context.Context, startDate, endDate time.Time) ([]model.Model, error)
}

type DataServiceFactory interface {
	MakeService() DataService
}
