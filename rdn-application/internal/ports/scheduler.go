package ports

import (
	"context"
	"time"
)

type Scheduler interface {
	ExecuteJob(context.Context, chan error, time.Time) bool
	ScheduleJob()
	Reschedule(string)
	GetStatus() (string, int)
}
