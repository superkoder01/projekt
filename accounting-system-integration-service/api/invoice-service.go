package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/mongodb"
	model "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger = logging.MustGetLogger("api")
)

type InvoiceService struct {
	session mongodb.MongoSession
}

func NewInvoiceService(mongoSession mongodb.MongoSession) *InvoiceService {
	return &InvoiceService{mongoSession}
}

func (is *InvoiceService) GetSummaryInvoicesByDateRange(ctx echo.Context, params GetSummaryInvoicesByDateRangeParams) error {
	logger.Info("GetSummaryInvoicesByDateRange: ", params)

	return is.commonGetCustomerSummaryInvoicesByDateRange(ctx, params, "")
}

func (is *InvoiceService) GetCustomerSummaryInvoicesByDateRange(ctx echo.Context, customerId string, params GetSummaryInvoicesByDateRangeParams) error {
	logger.Infof("GetCustomerSummaryInvoicesByDateRange for customer %s and params %s ", customerId, params)

	return is.commonGetCustomerSummaryInvoicesByDateRange(ctx, params, customerId)
}

func (is *InvoiceService) GetSummaryInvoiceByInvoiceNumber(ctx echo.Context, invoiceNumber string) error {
	logger.Info("GetSummaryInvoiceByInvoiceNumber: ", invoiceNumber)

	invoiceNumberReplaced := strings.ReplaceAll(invoiceNumber, "_", "/")
	singleRes := is.session.FindOne(bson.M{mongodb.INVOICE_NUMBER_KEY: invoiceNumberReplaced, mongodb.HEADER_CONTENT_TYPE: mongodb.INVOICE}, mongodb.INVOICE_COLLECTION_NAME)

	if configuration.UseCommonModel() {
		var elem model.InvoiceSummary
		err := singleRes.Decode(&elem)
		if err != nil {
			logger.Error(err)
			if err == mongo.ErrNoDocuments {
				return echo.NewHTTPError(http.StatusOK, "No documents found")
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError, "Temporary server error")
			}
		}

		logger.Debug("Summary invoice: ", elem)
		return ctx.JSON(http.StatusOK, elem)
	} else {
		var elem model.InvoiceProsument
		err := singleRes.Decode(&elem)
		if err != nil {
			logger.Error(err)
			if err == mongo.ErrNoDocuments {
				return echo.NewHTTPError(http.StatusOK, "No documents found")
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError, "Temporary server error")
			}
		}
		logger.Debug("Invoice from database: ", elem)
		result := createSummaryInvoice(elem)
		logger.Debug("Summary invoice: ", result)
		return ctx.JSON(http.StatusOK, result)
	}
}

func (is *InvoiceService) commonGetCustomerSummaryInvoicesByDateRange(ctx echo.Context, params GetSummaryInvoicesByDateRangeParams, customerId string) error {
	logger.Debugf("commonGetCustomerSummaryInvoicesByDateRange for customer %s and params %s ", customerId, params)

	startTime, e := time.Parse(DATE_LAYOUT, params.StartAt)
	if e != nil {
		logger.Errorf("parse date error: %s", e)
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect date format provided for StartAt. Expected format: YYYY-MM-DD")
	}
	logger.Debug("StartTime: ", startTime)

	var filterDate primitive.M
	if params.EndAt == nil {
		filterDate = bson.M{"$gte": startTime}
	} else {
		endTimeProvided, e := time.Parse(DATE_LAYOUT, *params.EndAt)
		if e != nil {
			logger.Errorf("parse date error: %s", e)
			return echo.NewHTTPError(http.StatusBadRequest, "Incorrect date format provided for EndAt. Expected format: YYYY-MM-DD")
		}
		// adds one day to include invoices from end day
		endTime := endTimeProvided.AddDate(0, 0, 1)
		logger.Debug("EndTime: ", endTime)

		filterDate = bson.M{"$gte": startTime, "$lte": endTime}
	}

	var filter primitive.M
	if customerId == "" {
		filter = bson.M{mongodb.INVOICE_ISSUE_DATE: filterDate, mongodb.HEADER_CONTENT_TYPE: mongodb.INVOICE}
	} else {
		filter = bson.M{mongodb.INVOICE_ISSUE_DATE: filterDate, mongodb.HEADER_CONTENT_TYPE: mongodb.INVOICE, mongodb.INVOICE_CUSTOMER_ID: customerId}
	}

	logger.Debug("Filter: ", filter)
	cursor, errDb := is.session.Find(filter, mongodb.INVOICE_COLLECTION_NAME)
	if errDb != nil {
		logger.Errorf("can not get collection from database: %s", errDb)
		return echo.NewHTTPError(http.StatusInternalServerError, "Temporary server error")
	}
	if configuration.UseCommonModel() {
		var result []model.InvoiceSummary
		i := 0
		for cursor.Next(context.TODO()) {
			var elem model.InvoiceSummary
			err := cursor.Decode(&elem)
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Temporary server error")
			}
			result = append(result, elem)
			i++
		}
		logger.Debug("Number of invoices: ", i)

		if i == 0 {
			return echo.NewHTTPError(http.StatusOK, "No documents found")
		} else {
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		var result []SummaryInvoice
		i := 0
		for cursor.Next(context.TODO()) {
			var elem model.InvoiceProsument
			err := cursor.Decode(&elem)
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Temporary server error")
			}
			r := createSummaryInvoice(elem)
			result = append(result, r)
			i++
		}
		logger.Debug("Number of invoices found: ", i)

		if i == 0 {
			return echo.NewHTTPError(http.StatusOK, "No documents found")
		} else {
			return ctx.JSON(http.StatusOK, result)
		}
	}
}

func createSummaryInvoice(elem model.InvoiceProsument) SummaryInvoice {
	var result SummaryInvoice = SummaryInvoice{}

	result.Header.Content.Type = &elem.Header.Content.Type
	result.Header.Provider = elem.Header.Provider
	result.Header.Version = elem.Header.Version

	// copy InvoiceDetails
	result.Payload.InvoiceDetails.CustomerId = elem.Payload.InvoiceDetails.RegistrationNumber
	result.Payload.InvoiceDetails.IssueDate = elem.Payload.InvoiceDetails.IssueDt.Format(DATE_LAYOUT)
	result.Payload.InvoiceDetails.InvoiceNumber = elem.Payload.InvoiceDetails.Number
	result.Payload.InvoiceDetails.Status = elem.Payload.InvoiceDetails.Status
	result.Payload.InvoiceDetails.ServiceDate = convertDateFormat(elem.Payload.InvoiceDetails.ServiceDt, INTERNAL_DATE_LAYOUT)
	result.Payload.InvoiceDetails.BillingStartDate = convertDateFormat(elem.Payload.InvoiceDetails.BillingStartDt, INTERNAL_DATE_LAYOUT)
	result.Payload.InvoiceDetails.BillingEndDate = convertDateFormat(elem.Payload.InvoiceDetails.BillingEndDt, INTERNAL_DATE_LAYOUT)
	result.Payload.InvoiceDetails.PaymentDueDate = elem.Payload.PaymentDetails.PaymentDueDt.Format(DATE_LAYOUT)
	result.Payload.InvoiceDetails.BankAccountNumber = elem.Payload.PaymentDetails.BankDetails.Account

	// copy CustomerDetails
	var ct CustomerDetails = CustomerDetails{}
	ct.CustomerId = elem.Payload.CustomerDetails.RegistrationNumber
	ct.DisplayName = &elem.Payload.CustomerDetails.DisplayName
	ct.FirstName = &elem.Payload.CustomerDetails.FirstName
	ct.LastName = &elem.Payload.CustomerDetails.LastName
	ct.Nip = &elem.Payload.CustomerDetails.Nip
	ct.Regon = &elem.Payload.CustomerDetails.Regon
	result.Payload.CustomerDetails = &ct

	// copy SellSummary
	result.Payload.SellSummary = convertSellSummary(elem.Payload.SellSummary)

	return result
}

func convertSellSummary(inputSS model.SellSummary) *SellSummary {
	var ss SellSummary = SellSummary{}
	var items []SellSummaryItem

	for _, i := range inputSS.Items {
		items = append(items, convertSellSummaryItem(i))
	}
	ss.GeneralInformation = &items

	// copy Total
	var t Total = Total{}
	t.GrossValue = inputSS.Total.GrossVal
	t.NetValue = inputSS.Total.NetVal
	t.TaxValue = inputSS.Total.TaxVal
	ss.Total = &t
	return &ss
}

func convertSellSummaryItem(item model.SellSummaryItem) SellSummaryItem {
	var result SellSummaryItem
	result.Name = item.ItemName
	result.Code = item.ItemCode
	result.GrossValue = item.GrossVal
	result.NetValue = item.NetVal
	result.TaxValue = item.TaxVal
	result.VatRate = float64(item.VatRate)

	return result
}

func convertDateFormat(inputDate string, layout string) string {
	inputDateObj, e := time.Parse(layout, inputDate)
	if e != nil {
		logger.Errorf("parse date error: %s", e)
	}

	result := inputDateObj.Format(DATE_LAYOUT)
	logger.Debugf("convert date from %s to %s", inputDate, result)
	return result
}
