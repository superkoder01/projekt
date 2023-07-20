package converter

import (
	"ConversionService/config"
	"ConversionService/internal/converter/html_filler"
	"ConversionService/internal/converter/html_to_pdf_converter"
	"ConversionService/internal/domain/model"
	"ConversionService/pkg/logger"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var titlePattern = regexp.Regexp{}
var cssPattern = regexp.Regexp{}
var jsPattern = regexp.Regexp{}

type Converter interface {
	Convert(data model.Data) ([]byte, string, error)
}

type converterStruct struct {
	htmlToPdfConverter html_to_pdf_converter.HtmlToPdfConverter
	config             config.Config
	logger             logger.Logger
	htmlFiller         html_filler.HtmlFiller
}

func init() {
	titlePattern = *regexp.MustCompile("<title>(.*)</title>")
	cssPattern = *regexp.MustCompile(" *<link rel=\"stylesheet\" href=\"(.*)\"/>\n")
	jsPattern = *regexp.MustCompile(" *<script type=\"text/javascript\" src=\"(.*)\">\n")
}

func NewConverter(htmlToPdfConverter html_to_pdf_converter.HtmlToPdfConverter, htmlFiller html_filler.HtmlFiller, config config.Config, logger logger.Logger) Converter {
	return &converterStruct{
		htmlToPdfConverter: htmlToPdfConverter,
		htmlFiller:         htmlFiller,
		config:             config,
		logger:             logger,
	}
}

func (receiver *converterStruct) Convert(data model.Data) ([]byte, string, error) {
	receiver.logger.Debugf("Processing template type %v", data.Header.Content.Type)

	templatePath, footerTemplatePath, footerMargin := receiver.getTemplatePath(data)
	if len(templatePath) < 1 {
		receiver.logger.Errorf("Template not found for type: %v provider: %v version: %v ",
			data.Header.Content.Type,
			data.Header.Provider,
			data.Header.Version,
		)
		return nil, "", fmt.Errorf("Cannot find template")
	}

	htmlTemplate, err := ioutil.ReadFile(templatePath)
	if err != nil {
		receiver.logger.Errorf("Cannot open template file %v ! %v", templatePath, err)
		return nil, "", err
	}

	html := string(htmlTemplate[:])

	html, err = receiver.joinHtmlAndCss(html, templatePath, func(filename string) ([]byte, error) {
		return ioutil.ReadFile(filename)
	})
	if err != nil {
		receiver.logger.Errorf("Cannot join HTML and CSS! %v", err)
		return nil, "", err
	}

	html, err = receiver.joinHtmlAndJs(html, templatePath, func(filename string) ([]byte, error) {
		return ioutil.ReadFile(filename)
	})
	if err != nil {
		receiver.logger.Errorf("Cannot join HTML and JS! %v", err)
		return nil, "", err
	}

	html, err = receiver.htmlFiller.Fill(html, data.Payload)
	if err != nil {
		receiver.logger.Errorf("Cannot fill HTML body ! %v", err)
		return nil, "", err
	}

	if receiver.config.Service.DebugMode == true {
		err = ioutil.WriteFile("doc/"+data.Header.Content.Type+".html", []byte(html), 0644)
		if err != nil {
			receiver.logger.Errorf("Unable to save HTML file !")
		}
	} else {
		receiver.logger.Debugf("Html file: " + data.Header.Content.Type + " .html created")
	}

	pdf, err := receiver.htmlToPdfConverter.Convert(html, footerTemplatePath, footerMargin)
	if err != nil {
		receiver.logger.Errorf("Cannot generate pdf file ! %v", err)
		return nil, "", err
	}

	fileName := receiver.getTitle(html)
	receiver.logger.Debugf("%v generated !", fileName)

	return pdf, fileName, nil
}

func (receiver *converterStruct) getTitle(html string) (fileName string) {
	titles := titlePattern.FindAllStringSubmatch(html, -1)
	if len(titles) != 1 || len(titles[0]) != 2 {
		return "Document.pdf"
	}

	return titles[0][1] + ".pdf"
}

func (receiver *converterStruct) joinHtmlAndCss(html string, htmlPath string, fileProvider func(filename string) ([]byte, error)) (string, error) {
	stringSubmatch := cssPattern.FindAllStringSubmatch(html, -1)
	content := ""
	for i, css := range stringSubmatch {
		if len(css) == 2 {

			dir, _ := filepath.Split(htmlPath)

			cssContent, err := fileProvider(dir + css[1])
			if err != nil {
				receiver.logger.Errorf("Cannot open file %v! %v", css[1], err)
				return "", err
			}

			if len(content) > 0 {
				content += "\n"
			}

			content += string(cssContent)
			if i < len(stringSubmatch)-1 {
				html = strings.Replace(html, css[0], "", 1)
			} else {
				html = strings.Replace(html, css[0], "<style type=\"text/css\">\n"+content+"\n</style>\n", 1)
			}
		}
	}
	return html, nil
}

func (receiver *converterStruct) joinHtmlAndJs(html string, htmlPath string, fileProvider func(filename string) ([]byte, error)) (string, error) {
	js := jsPattern.FindAllStringSubmatch(html, -1)
	if len(js) == 1 {
		if len(js[0]) == 2 {

			dir, _ := filepath.Split(htmlPath)

			jsContent, err := fileProvider(dir + js[0][1])
			if err != nil {
				receiver.logger.Errorf("Cannot open file %v! %v", js[0][1], err)
				return "", err
			}

			html = strings.Replace(html, js[0][0], "<script type=\"text/javascript\">\n"+string(jsContent), 1)
		}
	}
	return html, nil
}

func (receiver *converterStruct) getTemplatePath(data model.Data) (string, string, int) {
	templatePath, footerTemplatePath, footerMargin := "", "", 0
	for _, templateConfig := range receiver.config.Template {
		if strings.ToLower(templateConfig.Type) == strings.ToLower(data.Header.Content.Type) {
			for _, version := range templateConfig.Versions {
				if strings.ToLower(version) == strings.ToLower(data.Header.Provider+"-"+data.Header.Version) && receiver.isValid(templateConfig.Conditions, data.Payload) {
					templatePath = templateConfig.Path
					footerTemplatePath = templateConfig.FooterTemplatePath
					footerMargin = templateConfig.FooterMargin
					receiver.logger.Debugf("templatePath %v", templatePath)
					return templatePath, footerTemplatePath, footerMargin
				}
			}
		}
	}
	return templatePath, footerTemplatePath, footerMargin
}

func (receiver *converterStruct) isValid(conditions map[string]interface{}, data map[string]interface{}) bool {
	if conditions == nil {
		return true
	}

	if data == nil {
		return false
	}

	result := false
	for key, value := range conditions {
		path := strings.Split(key, ".")

		pattern, err := regexp.Compile(fmt.Sprintf("%v", value))
		if err == nil {

			foundedValue := receiver.findValue(path, data)
			foundedValueString := ""
			if foundedValue != nil {
				foundedValueString = fmt.Sprintf("%v", foundedValue)
			}
			if pattern.FindAllString(foundedValueString, -1) != nil {
				result = true
			} else {
				return false
			}

		} else {
			receiver.logger.Debugf("Cannot compile regexp %v", value)
		}
	}

	return result
}

func (receiver *converterStruct) findValue(path []string, data map[string]interface{}) any {
	if len(path) == 0 {
		return nil
	} else if len(path) == 1 {
		return data[path[0]]
	} else {
		var subData = data[path[0]].(map[string]interface{})
		return receiver.findValue(path[1:len(path)], subData)
	}
}
