package ports

import (
	"RDN-application/internal/model"
	"context"
	"time"
)

type DataCollector interface {
	CollectDataFromDate(ctx context.Context, date time.Time, channel chan model.Collection)
}

type CollectorFactory interface {
	MakeService() DataCollector
}
