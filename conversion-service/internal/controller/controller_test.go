package controller

import (
	"ConversionService/config"
	"ConversionService/internal/domain/model"
	"ConversionService/internal/mock"
	"ConversionService/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertToPdf200(t *testing.T) {
	recorder := convertPdf(model.Data{}, []byte("testdata"), "testfile", nil)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "testdata", string(recorder.Body.Bytes()))
	assert.Equal(t, "attachment; filename=testfile", recorder.Header().Get("Content-Disposition"))
}

func TestConvertToPdf400(t *testing.T) {
	recorder := convertPdf("test", nil, "", nil)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestConvertToPdf500(t *testing.T) {
	recorder := convertPdf(model.Data{}, nil, "", fmt.Errorf("Error"))
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func convertPdf(jsonData any, outputData []byte, fileName string, error error) *httptest.ResponseRecorder {
	inputData, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "/api/convert-pdf", bytes.NewReader(inputData))
	recorder := httptest.NewRecorder()
	var appLogger = logger.NewApiLogger(&config.Logger{})
	appLogger.InitLogger()
	var controller = NewController(mock.NewMockConverter(outputData, fileName, error), appLogger)
	var engine = gin.New()
	engine.POST("/api/convert-pdf", controller.ConvertPdf)
	engine.ServeHTTP(recorder, request)
	return recorder
}
