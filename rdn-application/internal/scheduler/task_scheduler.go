package scheduler

import (
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"github.com/jasonlvhit/gocron"
	"time"
)

const DateFormat = "02-01-2006 15:04"

type Scheduler struct {
	config config.AppConfig
	cron   *gocron.Scheduler
	logger logger.Logger
	job    *jobScheduler
}

func NewScheduler(config config.AppConfig, logger logger.Logger) *Scheduler {
	return &Scheduler{
		cron:   gocron.NewScheduler(),
		config: config,
		logger: logger,
	}
}

type jobScheduler struct {
	task    func(ctx context.Context, errChan chan error, time []time.Time)
	errChan chan error
	ctx     context.Context
}

func (scheduler *Scheduler) AppendJob(scheduledTask func(ctx context.Context, errChan chan error, time []time.Time), ctx context.Context, errChan chan error) {
	scheduler.job = &jobScheduler{
		task:    scheduledTask,
		errChan: errChan,
		ctx:     ctx,
	}
}

func (scheduler *Scheduler) Reschedule(newStartDate string) {
	oldJob, _ := scheduler.cron.NextRun()
	scheduler.cron.RemoveByRef(oldJob)

	job := scheduler.cron.Every(1).Day().At(newStartDate)
	err := job.Do(scheduler.job.task, scheduler.job.ctx, scheduler.job.errChan, []time.Time{})

	if err != nil {
		scheduler.logger.Fatalf("Failed to schedule job !! %v", err)
	} else {
		_, time := scheduler.cron.NextRun()
		scheduler.logger.Infof("Job scheduled: %v", time.Format(DateFormat))
		scheduler.cron.Start()
	}
}

func (scheduler *Scheduler) ScheduleJob() {
	job := scheduler.cron.Every(1).Day().At("17:35")
	err := job.Do(scheduler.job.task, scheduler.job.ctx, scheduler.job.errChan, []time.Time{})

	if err != nil {
		scheduler.logger.Fatalf("Failed to schedule job !! %v", err)
	} else {
		_, time := scheduler.cron.NextRun()
		scheduler.logger.Infof("Job scheduled: %v", time.Format(DateFormat))
		scheduler.cron.Start()
	}
}

func (scheduler *Scheduler) GetStatus() (string, int) {
	response := ""
	run, _ := scheduler.cron.NextRun()
	response += run.NextScheduledTime().Local().Format(DateFormat)

	return response, scheduler.cron.Len()
}

func (scheduler *Scheduler) ExecuteJob(ctx context.Context, errChan chan error, jobTime time.Time) bool {
	scheduler.logger.Infof("---Starting job---")
	scheduler.job.task(ctx, errChan, []time.Time{jobTime})
	scheduler.logger.Infof("---Job done---")
	return true
}
