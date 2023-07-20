package html_filler

import (
	"ConversionService/config"
	"ConversionService/pkg/logger"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindElementInJson(t *testing.T) {
	result := findElementInJson("{\"parent\":[{\"child\":\"value\"}]}", "parent.0.child", false)
	assert.Equal(t, "value", result)
}

func TestFindElementInJsonIndexOutOfRange(t *testing.T) {
	result := findElementInJson("{\"parent\":[]}", "parent.0.child", false)
	assert.Equal(t, "", result)
}

func TestConvertDate0(t *testing.T) {
	result := NewFiller().convertDate("2022-07-16T10:25:43.155030028+02:00", "2006-01-02T15:04:05Z07:00", "2/01/2006")
	assert.Equal(t, "16/07/2022", result)
}

func TestConvertDate1(t *testing.T) {
	result := NewFiller().convertDate("2022-07-06T10:25:43.155030028+02:00", "2006-01-02T15:04:05Z07:00", "2/01/2006")
	assert.Equal(t, "6/07/2022", result)
}

func TestConvertDate2(t *testing.T) {
	result := NewFiller().convertDate("2022-07-20T10:25:43.155031504+02:00", "2006-01-02T15:04:05Z07:00", "2/01/2006")
	assert.Equal(t, "20/07/2022", result)
}

func TestConvertDate3(t *testing.T) {
	result := NewFiller().convertDate("2022/07/19 07:25:22", "2006/01/02 15:04:05", "2.01.2006")
	assert.Equal(t, "19.07.2022", result)
}

func TestFillRequiredText0(t *testing.T) {
	html := "<p>${offerDetails.creationDate}</p>"
	json := "{\"offerDetails\":{\"creationDate\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>2022/07/19 07:25:22</p>", result)
}

func TestFillRequiredText1(t *testing.T) {
	html := "${customerDtls.displayName}, zam. w ${customerDtls.address.city} (${customerDtls.address.postCode}) ul. ${customerDtls.address.street}, Pesel: ${!customerDtls.pesel}"
	json := "{\"customerDtls\":{\"displayName\":\"displayName\", \"pesel\":\"pesel\", \"address\":{\"city\":\"city\", \"postCode\":\"postCode\", \"street\":\"street\"}}}"
	result := fill(html, json)
	assert.Equal(t, "displayName, zam. w city (postCode) ul. street, Pesel: pesel", result)
}

func TestFillNotRequiredText(t *testing.T) {
	html := "<p>${!offerDetails.creationDate}</p>"
	json := "{\"offerDetails\":{\"date\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p></p>", result)
}

func TestFillRequiredDate0(t *testing.T) {
	html := "<p>${convertDate(offerDetails.creationDate, \"2006/01/02 15:04:05\", \"2.01.2006\")}</p>"
	json := "{\"offerDetails\":{\"creationDate\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>19.07.2022</p>", result)
}

func TestFillRequiredDate1(t *testing.T) {
	html := "<p>${convertDate(paymentDetails.paymentDueDt, \"2006-01-02T15:04:05Z07:00\", \"2/01/2006\")}</p>"
	json := "{\"paymentDetails\":{\"paymentDueDt\":\"2022-07-20T15:15:15+02:00\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>20/07/2022</p>", result)
}

func TestFillRequiredDate2(t *testing.T) {
	html := "<p>${convertDate(invoiceDetails.issueDt, \"2006-01-02T15:04:05Z07:00\", \"2/01/2006\")}</p>"
	json := "{\"invoiceDetails\":{\"issueDt\":\"2022-07-06T15:15:15+02:00\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>6/07/2022</p>", result)
}

func TestFillNotDate(t *testing.T) {
	html := "<p>${convertDate(offerDetails.creationDate, \"2006/01/02 15:04:05\", \"2.01.2006\")}</p>"
	json := "{\"offerDetails\":{\"creationDate\":\"test\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>test</p>", result)
}

func TestFillNotRequiredDate0(t *testing.T) {
	html := "<p>${!convertDate(offerDetails.creationDate, \"2006/01/02 15:04:05\", \"2.01.2006\")}</p>"
	json := "{\"offerDetails\":{\"creationDate\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>19.07.2022</p>", result)
}

func TestFillNotRequiredDate1(t *testing.T) {
	html := "<p>${!convertDate(offerDetails.creationDate, \"2006/01/02 15:04:05\", \"2.01.2006\")}</p>"
	json := "{\"offerDetails\":{\"date\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p></p>", result)
}

func TestConvertFloat0(t *testing.T) {
	result := NewFiller().convertFloat("12.3456", "%.2f", "false")
	assert.Equal(t, "12.35", result)
}

func TestConvertFloat1(t *testing.T) {
	result := NewFiller().convertFloat("0.1234", "%.2f", "false")
	assert.Equal(t, "0.12", result)
}

func TestConvertFloat2(t *testing.T) {
	result := NewFiller().convertFloat("0.0001234", "%.5f", "false")
	assert.Equal(t, "0.00012", result)
}

func TestConvertFloat4(t *testing.T) {
	result := NewFiller().convertFloat("12.00", "%.2f", "true")
	assert.Equal(t, "12", result)
}

func TestConvertFloat5(t *testing.T) {
	result := NewFiller().convertFloat("12.34", "%.2f", "true")
	assert.Equal(t, "12.34", result)
}

func TestFillRequiredFloat(t *testing.T) {
	html := "<p>${convertFloat(offerDetails.price, \"%.2f\", false)}</p>"
	json := "{\"offerDetails\":{\"price\":\"12.345\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>12.35</p>", result)
}

func TestFillNotFloat(t *testing.T) {
	html := "<p>${convertFloat(offerDetails.price, \"%.2f\", false)}</p>"
	json := "{\"offerDetails\":{\"price\":\"test\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>test</p>", result)
}

func TestFillNotRequiredFloat(t *testing.T) {
	html := "<p>${!convertFloat(offerDetails.price, \"%.5f\", false)}</p>"
	json := "{\"offerDetails\":{\"price\":\"0.0001234\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>0.00012</p>", result)
}

func TestFillWithoutVariable(t *testing.T) {
	html := "<p>test</p>"
	json := "{\"offerDetails\":{\"price\":\"12.345\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>test</p>", result)
}

func TestFillJson(t *testing.T) {
	html := "<p>${.}</p>"
	json := "{\"offerDetails\":{\"date\":\"2022/07/19 07:25:22\"}}"
	result := fill(html, json)
	assert.Equal(t, "<p>{\"offerDetails\":{\"date\":\"2022/07/19 07:25:22\"}}</p>", result)
}

func findElementInJson(jsonString string, key string, optional bool) interface{} {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	return NewFiller().findElementInJson(jsonMap, key, optional)
}

func fill(html string, jsonString string) interface{} {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	html, _ = NewFiller().Fill(html, jsonMap)
	return html
}

func NewFiller() *htmlFillerStruct {
	appLogger := logger.NewApiLogger(&config.Logger{})
	appLogger.InitLogger()

	return &htmlFillerStruct{
		isDebugMode: false,
		logger:      appLogger,
	}
}
