package controllers

import (
	"RDN-application/internal/ports"
	"RDN-application/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const HourFormat = "15:04"
const DateFormat = "02-01-2006"

type schedulerController struct {
	scheduler ports.Scheduler
	logger    logger.Logger
}

func NewSchedulerHandler(scheduler ports.Scheduler, logger logger.Logger) *schedulerController {
	return &schedulerController{scheduler: scheduler, logger: logger}
}

func (controller *schedulerController) GetSchedulerStatus(ctx *gin.Context) {
	nextRun, jobsLen := controller.scheduler.GetStatus()

	/*todo add scheduler history*/
	ctx.JSON(200, "Next scheduled run at "+nextRun+". Number of scheduled jobs: "+fmt.Sprintf("%v", jobsLen))
}
func (controller *schedulerController) Reschedule(ctx *gin.Context) {
	hour := ctx.Param("hour")
	time, err := time.Parse(HourFormat, hour)
	if err != nil {
		HandleError(NewError("Invalid hour format", 400), ctx)
		return
	}
	controller.scheduler.Reschedule(time.Format(HourFormat))
	nextRun, jobsLen := controller.scheduler.GetStatus()

	ctx.JSON(200, "Sucess ! Next scheduled run at "+nextRun+". Number of scheduled jobs: "+fmt.Sprintf("%v", jobsLen))
}

func (controller *schedulerController) ScheduleNow(ctx *gin.Context) {
	errChan, dataChan := make(chan error), make(chan bool)
	go func() {
		dataChan <- controller.scheduler.ExecuteJob(ctx, errChan, time.Now().Local())
	}()

	select {
	case <-ctx.Request.Context().Done():
		controller.logger.Infof("Http context canceled")
	case err := <-errChan:
		HandleError(NewError("Failed to execute task! Exception: "+err.Error(), 400), ctx)
		return
	case <-dataChan:
		/*navigate to today's data*/
		ctx.JSON(200, "")
	}
}

func (controller *schedulerController) RunWithDate(ctx *gin.Context) {
	date := ctx.Param("date")
	jobDate, err := time.Parse(DateFormat, date)
	if err != nil {
		HandleError(NewError("Invalid date format", 400), ctx)
		return
	}

	errChan, dataChan := make(chan error), make(chan bool)
	go func() {
		jobDate = jobDate.AddDate(0, 0, -1)
		dataChan <- controller.scheduler.ExecuteJob(ctx, errChan, jobDate)
	}()

	select {
	case err := <-errChan:
		HandleError(NewError("Failed to execute task! Exception: "+err.Error(), 400), ctx)
		return
	case <-dataChan:
		/*navigate to today's data*/
		ctx.JSON(200, "")
	}
}
