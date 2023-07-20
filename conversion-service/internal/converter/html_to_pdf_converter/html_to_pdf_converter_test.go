package html_to_pdf_converter

import (
	"ConversionService/config"
	"ConversionService/internal/mock"
	"ConversionService/pkg/logger"
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var cfg = config.HtmlToPdfConverterConfig{
	Url:        "url",
	Dpi:        100,
	MarginLRTB: []int{1, 2, 3, 4},
}

func TestConvert200WithFooter(t *testing.T) {
	var data, _ = convert("html", "footerPath", 5, func(request *http.Request, data Data, content string) (int, string) {

		assert.Equal(t, "POST", request.Method)
		assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
		assert.Equal(t, "html", content)
		assert.Equal(t, "footerPath", data.Options["footer-html"])
		assert.Equal(t, float64(1), data.Options["margin-left"])
		assert.Equal(t, float64(2), data.Options["margin-right"])
		assert.Equal(t, float64(3), data.Options["margin-top"])
		assert.Equal(t, float64(9), data.Options["margin-bottom"])
		assert.Equal(t, float64(100), data.Options["dpi"])
		assert.Equal(t, "", data.Options["no-stop-slow-scripts"])
		assert.Equal(t, "", data.Options["enable-local-file-access"])

		return 200, "data"
	})

	assert.Equal(t, "data", data)
}

func TestConvert200WithoutFooter(t *testing.T) {
	var data, _ = convert("html", "", 0, func(request *http.Request, data Data, content string) (int, string) {

		assert.Equal(t, "POST", request.Method)
		assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
		assert.Equal(t, "html", content)
		assert.Nil(t, data.Options["footer-html"])
		assert.Equal(t, float64(1), data.Options["margin-left"])
		assert.Equal(t, float64(2), data.Options["margin-right"])
		assert.Equal(t, float64(3), data.Options["margin-top"])
		assert.Equal(t, float64(4), data.Options["margin-bottom"])
		assert.Equal(t, float64(100), data.Options["dpi"])
		assert.Equal(t, "", data.Options["no-stop-slow-scripts"])
		assert.Equal(t, "", data.Options["enable-local-file-access"])

		return 200, "data"
	})

	assert.Equal(t, "data", data)
}

func TestConvert500(t *testing.T) {
	var _, err = convert("html", "", 0, func(request *http.Request, data Data, content string) (int, string) {
		return 500, ""
	})

	assert.NotNil(t, err)
}

func convert(inputData string, footerPath string, footerMarg int, verify func(request *http.Request, data Data, content string) (int, string)) (string, error) {
	appLogger := logger.NewApiLogger(&config.Logger{})
	appLogger.InitLogger()

	var outputData, err = NewHtmlToPdfConverter(cfg, appLogger, &mock.MockHttpClient{
		DoFunc: func(request *http.Request) (*http.Response, error) {
			byteData, _ := ioutil.ReadAll(request.Body)
			var data Data
			json.Unmarshal(byteData, &data)
			content, _ := base64.StdEncoding.DecodeString(data.Contents)
			code, outputData := verify(request, data, string(content))

			return &http.Response{
				StatusCode: code,
				Body:       ioutil.NopCloser(strings.NewReader(outputData)),
			}, nil
		},
	}).Convert(inputData, footerPath, footerMarg)
	return string(outputData), err
}
