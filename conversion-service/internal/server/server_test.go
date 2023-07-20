package server

import (
	"ConversionService/config"
	"ConversionService/internal/converter/html_to_pdf_converter"
	"ConversionService/internal/domain/model"
	"ConversionService/internal/mock"
	"ConversionService/pkg/logger"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	cfg = config.Config{
		Service: config.ServiceConfig{
			ReleaseMode: true,
		},
		HtmlToPdfConverter: config.HtmlToPdfConverterConfig{
			Url:        "http://127.0.0.1:80",
			Dpi:        300,
			MarginLRTB: []int{15, 15, 15, 15},
		},
		Template: []config.TemplateEngineConfig{
			{
				Type:               "Contract",
				Versions:           []string{"proider-1.0"},
				Path:               "../../template/contract-B2C.html",
				FooterTemplatePath: "/myfiles/template/footer.html",
				FooterMargin:       10,
			},
		},
		Logger: config.Logger{
			Encoding: "console",
		},
	}
)

func TestConvertContractToPdf(t *testing.T) {
	var contract = model.Data{
		Header: model.Header{
			Version:  "1.0",
			Provider: "proider",
			Content:  model.Content{Type: "contract"},
		},
		Payload: map[string]interface{}{
			"contractDetails": map[string]interface{}{
				"number": "ContractDetailsNumber",
			},
		},
	}

	recorder := execute(contract, func(request *http.Request, content string, options map[string]interface{}) {
		assert.Equal(t, "127.0.0.1:80", request.URL.Host)
		assert.Equal(t, "POST", request.Method)
		assert.Contains(t, content, "ContractDetailsNumber")
		assert.Equal(t, "/myfiles/template/footer.html", options["footer-html"])
		assert.Equal(t, float64(15), options["margin-left"])
		assert.Equal(t, float64(15), options["margin-right"])
		assert.Equal(t, float64(15), options["margin-top"])
		assert.Equal(t, float64(25), options["margin-bottom"])
		assert.Equal(t, float64(300), options["dpi"])
		assert.Equal(t, "", options["no-stop-slow-scripts"])
		assert.Equal(t, "", options["enable-local-file-access"])
	})

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func execute(content any, verify func(request *http.Request, content string, options map[string]interface{})) *httptest.ResponseRecorder {
	appLogger := logger.NewApiLogger(&cfg.Logger)
	appLogger.InitLogger()

	engine := NewServer(&cfg, appLogger).newEngine(&mock.MockHttpClient{
		DoFunc: func(request *http.Request) (*http.Response, error) {
			bytes, _ := ioutil.ReadAll(request.Body)
			var body html_to_pdf_converter.Data
			json.Unmarshal(bytes, &body)
			byteHtml, _ := base64.StdEncoding.DecodeString(body.Contents)

			verify(request, string(byteHtml), body.Options)

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			}, nil
		},
	})

	contentBytes, _ := json.Marshal(content)
	request, _ := http.NewRequest("POST", "/api/convert-pdf", bytes.NewReader(contentBytes))
	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, request)

	return recorder
}
