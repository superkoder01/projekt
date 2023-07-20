package ports

import "github.com/gin-gonic/gin"

type DataHandler interface {
	GetDay(ctx *gin.Context)
	GetDayList(ctx *gin.Context)
}

type ScheduleHandler interface {
	GetSchedulerStatus(ctx *gin.Context)
	ScheduleNow(ctx *gin.Context)
	Reschedule(ctx *gin.Context)
	RunWithDate(ctx *gin.Context)
}
