package converter

import (
	"ConversionService/config"
	"ConversionService/internal/domain/model"
	"ConversionService/pkg/logger"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetTitle0(t *testing.T) {
	html := "<html>\n<head>\n    <title>test</title>\n    <style type=\"text/css\"></style>\n</head>\n<body>\n</body>\n</html>"
	result := newConverter().getTitle(html)
	assert.Equal(t, "test.pdf", result)
}

func TestGetTitle1(t *testing.T) {
	html := "<html>\n<head>\n    <style type=\"text/css\"></style>\n</head>\n<body>\n</body>\n</html>"
	result := newConverter().getTitle(html)
	assert.Equal(t, "Document.pdf", result)
}

func TestJoinHtmlAndCss0(t *testing.T) {
	html := "<html>\n<head>\n    <link rel=\"stylesheet\" href=\"styles.css\"/>\n</head>\n<body>\n</body>\n</html>"
	css := ".p0 {\n    text-align: center;\n}"
	result, _ := newConverter().joinHtmlAndCss(html, "template/content.html", func(filename string) ([]byte, error) {
		assert.Equal(t, "template/styles.css", filename)
		return []byte(css), nil
	})
	assert.Equal(t, "<html>\n<head>\n<style type=\"text/css\">\n.p0 {\n    text-align: center;\n}\n</style>\n</head>\n<body>\n</body>\n</html>", result)
}

func TestJoinHtmlAndCss1(t *testing.T) {
	html := "<html>\n<head>\n    <link rel=\"stylesheet\" href=\"styles0.css\"/>\n    <link rel=\"stylesheet\" href=\"styles1.css\"/>\n</head>\n<body>\n</body>\n</html>"
	css0 := ".p0 {\n    text-align: center;\n}"
	css1 := ".p1 {\n    text-align: left;\n}"
	result, _ := newConverter().joinHtmlAndCss(html, "template/content.html", func(filename string) ([]byte, error) {
		if filename == "template/styles0.css" {
			return []byte(css0), nil
		} else if filename == "template/styles1.css" {
			return []byte(css1), nil
		}
		return nil, fmt.Errorf("Cannot find file %v", filename)
	})
	assert.Equal(t, "<html>\n<head>\n<style type=\"text/css\">\n.p0 {\n    text-align: center;\n}\n.p1 {\n    text-align: left;\n}\n</style>\n</head>\n<body>\n</body>\n</html>", result)
}

func TestJoinHtmlAndJs(t *testing.T) {
	html := "<html>\n<head>\n    <script type=\"text/javascript\" src=\"common.js\">\n        function test0() {\n        }\n</script>\n</head>\n<body>\n</body>\n</html>"
	js := "function test1() {\n    }\n"
	result, _ := newConverter().joinHtmlAndJs(html, "template/content.html", func(filename string) ([]byte, error) {
		assert.Equal(t, "template/common.js", filename)
		return []byte(js), nil
	})
	assert.Equal(t, "<html>\n<head>\n<script type=\"text/javascript\">\nfunction test1() {\n    }\n        function test0() {\n        }\n</script>\n</head>\n<body>\n</body>\n</html>", result)
}

func TestGetTemplatePath0(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{
		Template: []config.TemplateEngineConfig{
			newTemplateEngineConfig("path0", "provider-0", "type0", map[string]interface{}{"key": "value"}),
			newTemplateEngineConfig("path1", "provider-0", "type0", nil),
		},
	}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "type0", map[string]interface{}{"test": "test"}))
	assert.Equal(t, "path1", result)
}

func TestGetTemplatePath1(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{
		Template: []config.TemplateEngineConfig{
			newTemplateEngineConfig("path0", "provider-0", "type0", map[string]interface{}{"key": "value"}),
			newTemplateEngineConfig("path1", "provider-0", "type0", nil),
		},
	}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "type0", map[string]interface{}{"key": "value"}))
	assert.Equal(t, "path0", result)
}

func TestGetTemplatePath2(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{
		Template: []config.TemplateEngineConfig{
			newTemplateEngineConfig("path0", "provider-0", "type0", map[string]interface{}{"key": "value"}),
			newTemplateEngineConfig("path1", "provider-0", "type0", nil),
		},
	}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "type0", map[string]interface{}{"key1": "value"}))
	assert.Equal(t, "path1", result)
}

func TestGetTemplatePath3(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{
		Template: []config.TemplateEngineConfig{
			newTemplateEngineConfig("path0", "provider-0", "type0", map[string]interface{}{"key": map[string]interface{}{"key1": "value"}}),
			newTemplateEngineConfig("path1", "provider-0", "type0", nil),
		},
	}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "type0", map[string]interface{}{"key": "value"}))
	assert.Equal(t, "path1", result)
}

func TestGetTemplatePathContractB2BWithRepurchase(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{Template: []config.TemplateEngineConfig{newTemplateEngineConfig("path0", "provider-0", "Contract", map[string]interface{}{
		"contractDetails.agreementType": "B2B",
		"repurchase.id":                 ".+",
	})}}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "Contract", map[string]interface{}{
		"contractDetails": map[string]interface{}{
			"agreementType": "B2B",
		},
		"repurchase": map[string]interface{}{
			"id": "test",
		},
	}))
	assert.Equal(t, "path0", result)
}

func TestGetTemplatePathContractB2BWithoutRepurchase(t *testing.T) {
	converter := newConverter()
	converter.config = config.Config{Template: []config.TemplateEngineConfig{newTemplateEngineConfig("path0", "provider-0", "Contract", map[string]interface{}{
		"contractDetails.agreementType": "B2B",
		"repurchase.id":                 "^$",
	})}}
	result, _, _ := converter.getTemplatePath(newData("provider", 0, "Contract", map[string]interface{}{
		"contractDetails": map[string]interface{}{
			"agreementType": "B2B",
		},
		"repurchase": map[string]interface{}{},
	}))
	assert.Equal(t, "path0", result)
}

func newConverter() *converterStruct {
	appLogger := logger.NewApiLogger(&config.Logger{})
	appLogger.InitLogger()
	return &converterStruct{logger: appLogger}
}

func newData(provider string, version int, contentType string, payoad map[string]interface{}) model.Data {
	return model.Data{
		Header: model.Header{
			Provider: provider,
			Version:  strconv.Itoa(version),
			Content: model.Content{
				Type: contentType,
			},
		},
		Payload: payoad,
	}
}

func newTemplateEngineConfig(path string, version string, contentType string, conditions map[string]interface{}) config.TemplateEngineConfig {
	return config.TemplateEngineConfig{
		Path:       path,
		Versions:   []string{version},
		Type:       contentType,
		Conditions: conditions,
	}
}
