package html_to_pdf_converter

import (
	"ConversionService/config"
	"ConversionService/pkg/logger"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HtmlToPdfConverter interface {
	Convert(html string, footerPath string, footerMargin int) ([]byte, error)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Data struct {
	Contents string                 `json:"contents"`
	Options  map[string]interface{} `json:"options"`
}

type htmlToPdfConverterStruct struct {
	config     config.HtmlToPdfConverterConfig
	logger     logger.Logger
	Options    map[string]interface{} `json:"options"`
	httpClient HttpClient
}

func NewHtmlToPdfConverter(config config.HtmlToPdfConverterConfig, logger logger.Logger, httpClient HttpClient) HtmlToPdfConverter {
	options := map[string]interface{}{
		"margin-left":              config.MarginLRTB[0],
		"margin-right":             config.MarginLRTB[1],
		"margin-top":               config.MarginLRTB[2],
		"margin-bottom":            config.MarginLRTB[3],
		"dpi":                      config.Dpi,
		"no-stop-slow-scripts":     "",
		"enable-local-file-access": "",
	}
	if config.DebugJavascript == true {
		options["debug-javascript"] = ""
	}
	return &htmlToPdfConverterStruct{
		config:     config,
		logger:     logger,
		httpClient: httpClient,
		Options:    options,
	}
}

func (receiver *htmlToPdfConverterStruct) Convert(html string, footerPath string, footerMargin int) ([]byte, error) {
	if len(footerPath) > 1 {
		receiver.logger.Debugf("Using footer \"%v\"", footerPath)
		receiver.Options["footer-html"] = footerPath
	} else {
		delete(receiver.Options, "footer-html")
	}

	receiver.Options["margin-bottom"] = receiver.config.MarginLRTB[3] + footerMargin

	data, err := json.Marshal(Data{
		Contents: base64.StdEncoding.EncodeToString([]byte(html)),
		Options:  receiver.Options,
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", receiver.config.Url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := receiver.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		receiver.logger.Warnf("Invalid response ! %v", response)
		return nil, fmt.Errorf("response %v", response.StatusCode)
	}

	return data, nil
}
