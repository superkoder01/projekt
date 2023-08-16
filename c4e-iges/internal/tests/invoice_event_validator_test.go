/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	require2 "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/validators"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/mocks"
	"os"
	"testing"
)

type InvoiceEventValidatorTestSuite struct {
	suite.Suite
	Cfg *config.AppConfig
}

func (suite *InvoiceEventValidatorTestSuite) SetupSuite() {
	os.Setenv("CONFIGPATH", "../../config/")
	suite.Cfg, _ = config.LoadConfig()
}

func (suite *InvoiceEventValidatorTestSuite) TearDownSuite() {
}

func (suite *InvoiceEventValidatorTestSuite) SetupTest() {
}

func (suite *InvoiceEventValidatorTestSuite) TearDownTest() {
}

func (suite *InvoiceEventValidatorTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *InvoiceEventValidatorTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *InvoiceEventValidatorTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println("=== Suite Stats:")
	fmt.Printf("=> suite duration: %v [s]\n", stats.End.Sub(stats.Start).Seconds())
	for _, v := range stats.TestStats {
		fmt.Printf("=> %s : %v [s] : %v\n", v.TestName, v.End.Sub(v.Start).Seconds(), passed(v.Passed))
	}
}

//func passed(status bool) string {
//	if status {
//		return fmt.Sprint("PASSED")
//	}
//	return fmt.Sprint("FAILED")
//}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestInvoiceEventValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(InvoiceEventValidatorTestSuite))
}

///////////////////////////////////////////////////////////////////////
func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventOK() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_06"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.NoError(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}

func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventNoContract() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_06"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.Error(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}

func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventNoBillingStartDate() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_06"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.Error(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}

func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventNoBillingEndDate() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_06"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.Error(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}

func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventNoSapData() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: nil,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.Error(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}

func (suite *InvoiceEventValidatorTestSuite) TestInvoiceEventEmptyDapData() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceEventValidator := validators.NewInvoiceEventValidator(loggerMock, suite.Cfg)

	require.Error(invoiceEventValidator.ValidateInvoiceEvent(ctx, invoiceEvent))
}
