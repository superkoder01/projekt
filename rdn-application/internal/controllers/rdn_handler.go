package controllers

import (
	"RDN-application/internal/ports"
	"github.com/gin-gonic/gin"
	"time"
)

const DateLayout = "2006-01-02"

type dataHandler struct {
	service ports.DataService
}

func NewDataHandler(service ports.DataService) *dataHandler {
	return &dataHandler{service: service}
}

func (c *dataHandler) GetDay(ctx *gin.Context) {
	date := ctx.Param("date")
	time, err := time.Parse(DateLayout, date)

	if err != nil {
		HandleError(NewError("Invalid date format", 400), ctx)
		return
	}

	data, err := c.service.GetDataFromDate(ctx, time)

	if err != nil {
		HandleError(err, ctx)
		return
	}
	ctx.JSON(200, data)
}

func (c *dataHandler) GetDayList(ctx *gin.Context) {
	startDate, endDate := ctx.Param("startDate"), ctx.Param("endDate")
	startTime, errSt := time.Parse(DateLayout, startDate)
	endTime, errEn := time.Parse(DateLayout, endDate)

	if errSt != nil || errEn != nil {
		HandleError(NewError("Invalid date format", 400), ctx)
		return
	}

	data, err := c.service.GetDataInBetween(ctx, startTime, endTime)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	ctx.JSON(200, data)
}
