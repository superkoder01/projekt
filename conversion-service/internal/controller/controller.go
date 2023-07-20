package controller

import (
	"ConversionService/internal/converter"
	"ConversionService/internal/domain/model"
	"ConversionService/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	converter converter.Converter
	logger    logger.Logger
}

func NewController(converter converter.Converter, logger logger.Logger) *Controller {
	return &Controller{
		converter: converter,
		logger:    logger,
	}
}

func (receiver *Controller) ConvertPdf(context *gin.Context) {
	receiver.logger.Debugf("Convert PDF. Request: %v", context.Request)

	inputData, _ := ioutil.ReadAll(context.Request.Body)
	receiver.logger.Debugf("Body: %v", string(inputData))

	var jsonData model.Data
	err := json.Unmarshal(inputData, &jsonData)
	if err != nil {
		receiver.logger.Errorf("Bad request", err)
		context.Status(http.StatusBadRequest)
		_, _ = context.Writer.Write([]byte(err.Error()))
		return
	}

	outputData, fileName, err := receiver.converter.Convert(jsonData)
	if err != nil {
		receiver.logger.Errorf("Internal server error", err)
		context.Status(http.StatusInternalServerError)
		_, _ = context.Writer.Write([]byte(err.Error()))
		return
	}

	receiver.logger.Debugf("Request processed successful")
	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Transfer-Encoding", "binary")
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	context.Data(http.StatusCreated, "application/octet-stream", outputData)
}
