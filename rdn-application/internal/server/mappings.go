package server

import (
	"RDN-application/internal/controllers"
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/api/"
)

func (s *Server) MapHandlers(engine *gin.Engine) {
	route := engine.Group(basePath)

	dataHandler := controllers.NewDataHandler(s.dataService)
	schedulerHandler := controllers.NewSchedulerHandler(s.scheduler, s.logger)

	/*Data*/
	route.GET("tge/:date", dataHandler.GetDay)
	route.GET("tge/from/:startDate/to/:endDate", dataHandler.GetDayList)

	/*Scheduler*/
	route.GET("scheduler/run/:date", schedulerHandler.RunWithDate)
	route.GET("scheduler/status", schedulerHandler.GetSchedulerStatus)
	route.GET("scheduler/reschedule/:hour", schedulerHandler.Reschedule)
	route.GET("scheduler/run/now", schedulerHandler.ScheduleNow)
}
