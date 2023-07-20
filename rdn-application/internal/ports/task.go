package ports

import (
	"context"
	"time"
)

type Task interface {
	ProcessCollectionTask() func(context.Context, chan error, []time.Time)
}
