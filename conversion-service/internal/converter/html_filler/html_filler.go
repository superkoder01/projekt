package html_filler

import (
	"ConversionService/pkg/logger"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var variablePattern = regexp.Regexp{}
var convertDatePattern = regexp.Regexp{}
var convertFloatPattern = regexp.Regexp{}

func init() {
	variablePattern = *regexp.MustCompile("\\$\\{([a-zA-Z0-9, /:%\"\\(\\)\\!\\.\\-]*)\\}")
	convertDatePattern = *regexp.MustCompile("convertDate\\((.*), *\"(.*)\", *\"(.*)\"\\)")
	convertFloatPattern = *regexp.MustCompile("convertFloat\\((.*), *\"(.*)\", *(.*)\\)")
}

type HtmlFiller interface {
	Fill(html string, json map[string]interface{}) (string, error)
}

type htmlFillerStruct struct {
	isDebugMode bool
	logger      logger.Logger
}

func NewHtmlFiller(isDebugMode bool, logger logger.Logger) HtmlFiller {
	return &htmlFillerStruct{
		isDebugMode: isDebugMode,
		logger:      logger,
	}
}

func (receiver *htmlFillerStruct) Fill(html string, data map[string]interface{}) (string, error) {
	receiver.logger.Debugf("Processing template...")

	for _, variableSubmatch := range variablePattern.FindAllStringSubmatch(html, -1) {

		if len(variableSubmatch) == 2 {
			variable := variableSubmatch[1]
			var value string

			if variable == "." {

				bytes, err := json.Marshal(data)
				if err != nil {
					return "", err
				}

				value = string(bytes)

			} else {

				optional := strings.Index(variableSubmatch[0], "!") >= 0

				if optional == true {
					variable = strings.Replace(variable, "!", "", 1)
				}

				convertDate := convertDatePattern.FindAllStringSubmatch(variable, -1)
				convertFloat := convertFloatPattern.FindAllStringSubmatch(variable, -1)

				if convertDate != nil && len(convertDate) == 1 && len(convertDate[0]) == 4 {
					value = fmt.Sprintf("%v", receiver.findElementInJson(data, convertDate[0][1], optional))
					value = receiver.convertDate(value, convertDate[0][2], convertDate[0][3])
				} else if convertFloat != nil && len(convertFloat) == 1 && len(convertFloat[0]) == 4 {
					value = fmt.Sprintf("%v", receiver.findElementInJson(data, convertFloat[0][1], optional))
					value = receiver.convertFloat(value, convertFloat[0][2], convertFloat[0][3])
				} else {
					value = fmt.Sprintf("%v", receiver.findElementInJson(data, variable, optional))
				}
			}

			html = strings.Replace(html, variableSubmatch[0], value, 1)
		} else {
			html = strings.Replace(html, variableSubmatch[0], "", 1)
		}
	}

	receiver.logger.Debugf("Processing completed !")
	return html, nil
}

func (receiver *htmlFillerStruct) findElementInJson(json map[string]interface{}, key string, isOptional bool) interface{} {
	newKey := strings.Split(key, ".")[0]

	/*todo move to some wrapper*/
	if receiver.isDebugMode == true {
		receiver.logger.Debugf("key: %v", key)
	}

	data := json[newKey]
	switch data.(type) {
	case string:
		return data
	case int:
		return data
	case int32:
		return data
	case int64:
		return data
	case float32:
		return data
	case float64:
		return data
	case []interface{}:
		var element interface{}
		indexString := strings.Split(key, ".")[1]
		_, after, _ := strings.Cut(key, newKey+"."+indexString+".")
		index, _ := strconv.Atoi(indexString)
		var length = len(data.([]interface{}))

		if index >= length && isOptional == true {
			return ""
		}

		if length == 0 || index >= length {
			receiver.logger.Warnf("Key not found %v", key)
			return ""
		}

		//engine.logger.Infof("element : %v, key: %v, indexString : %v, index : %v, data : %v", element, key, indexString, index, data.([]interface{})[index])
		switch data.([]interface{})[index].(type) {
		case string:
			return data.([]interface{})[index]
		case int:
			return data.([]interface{})[index]
		case int32:
			return data.([]interface{})[index]
		case int64:
			return data.([]interface{})[index]
		case float32:
			return data.([]interface{})[index]
		case float64:
			return data.([]interface{})[index]
		default:
			element = data.([]interface{})[index]
			return receiver.findElementInJson(element.(map[string]interface{}), after, isOptional)
		}

	case map[string]interface{}:
		_, after, found := strings.Cut(key, newKey+".")
		if found == false {
			receiver.logger.Errorf("Key not found %v", key)
			return nil
		}
		return receiver.findElementInJson(data.(map[string]interface{}), after, isOptional)
	default:
		if isOptional == true {
			return ""
		} else if receiver.isDebugMode == true {
			receiver.logger.Errorf("unknown field !!")
		}

		receiver.logger.Warnf("Key not found %v", key)
		return ""
	}
}

func (receiver *htmlFillerStruct) convertDate(value string, currentFormat string, newFormat string) string {
	if value == "" {
		return value
	}

	var date, err = time.Parse(currentFormat, value)
	if err != nil {
		receiver.logger.Warnf("Cannot parse %v %v. %v", value, currentFormat, err)
		return value
	}

	return date.Format(newFormat)
}

func (receiver *htmlFillerStruct) convertFloat(value string, format string, removeZeros string) string {
	if value == "" {
		return value
	}

	float, err := strconv.ParseFloat(value, 32)
	if err != nil {
		receiver.logger.Warnf("Cannot parse %v. %v", value, err)
		return value
	}

	value = fmt.Sprintf(format, float)
	if removeZeros == "true" {
		return strings.TrimRight(strings.TrimRight(value, "0"), ".")
	}
	return value
}
