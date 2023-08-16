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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/generators"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/mocks"
	"os"
	"testing"
)

type NumberGeneratorTestSuite struct {
	suite.Suite
	Cfg *config.AppConfig
}

func (suite *NumberGeneratorTestSuite) SetupSuite() {
	os.Setenv("CONFIGPATH", "../../config/")
	suite.Cfg, _ = config.LoadConfig()
}

func (suite *NumberGeneratorTestSuite) TearDownSuite() {
}

func (suite *NumberGeneratorTestSuite) SetupTest() {
}

func (suite *NumberGeneratorTestSuite) TearDownTest() {
}

func (suite *NumberGeneratorTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *NumberGeneratorTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *NumberGeneratorTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
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
func TestNumberGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(NumberGeneratorTestSuite))
}

func mockResultFn() (func(ctx context.Context, customerId string, from string, to string) int64, func(ctx context.Context, customerId string, from string, to string) error) {
	return func(ctx context.Context, customerId string, from string, to string) int64 {
			if customerId == "123123120" && from == "01/06/2022" && to == "30/06/2022" {
				return 0
			} else if customerId == "123123122" && from == "01/06/2022" && to == "30/06/2022" {
				return 2
			} else {
				return 0
			}
		},
		func(ctx context.Context, customerId string, from string, to string) error {
			if customerId == "shouldFail" {
				return fmt.Errorf("unable to connect database")
			}

			return nil
		}
}

func mockRepurchaseResultFn() (func(ctx context.Context, customerId string, from string, to string) int64, func(ctx context.Context, customerId string, from string, to string) error) {
	return func(ctx context.Context, customerId string, from string, to string) int64 {
			if customerId == "123123120" && from == "01/06/2022" && to == "30/06/2022" {
				return 0
			} else if customerId == "123123122" && from == "01/06/2022" && to == "30/06/2022" {
				return 2
			} else {
				return 0
			}
		},
		func(ctx context.Context, customerId string, from string, to string) error {
			if customerId == "shouldFail" {
				return fmt.Errorf("unable to connect database")
			}

			return nil
		}
}

///////////////////////////////////////////////////////////////////////
func (suite *NumberGeneratorTestSuite) TestGenerateInvoiceNumberOK() {
	require := require2.New(suite.T())

	ctx := context.Background()
	event := invoice.InvoiceEvent{
		Contract:            "GSv0.1/4/05/07/2022",
		ServiceAccessPoints: nil,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	tests := []struct {
		customerId            string
		expectedInvoiceNumber string
		event                 invoice.InvoiceEvent
	}{
		{"123123120", "2022/06/123123120/SP/1", event},
		{"123123121", "2022/06/123123121/SP/1", event},
		{"123123122", "2022/06/123123122/SP/3", event},
		{"123123123", "2022/06/123123123/SP/1", event},
		{"shouldFail", "2022/06/shouldFail/SP/1", event},
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Warnf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceRepoMock := mocks.NewInvoiceRepo(suite.T())
	invoiceRepoMock.On("CountInvoices", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResultFn())

	numberGenerator := generators.NewInvoiceNumberGenerator(invoiceRepoMock, loggerMock, suite.Cfg)

	for _, test := range tests {
		actualNumber, err := numberGenerator.GetNumber(ctx, test.customerId, test.event)
		require.Equal(test.expectedInvoiceNumber, actualNumber)
		require.NoError(err, "should not happen")
	}
}

func (suite *NumberGeneratorTestSuite) TestGenerateRepurchaseInvoiceNumberOK() {
	require := require2.New(suite.T())

	ctx := context.Background()
	event := invoice.InvoiceEvent{
		Contract:            "GSv0.1/4/05/07/2022",
		ServiceAccessPoints: nil,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	tests := []struct {
		customerId            string
		expectedInvoiceNumber string
		event                 invoice.InvoiceEvent
	}{
		{"123123120", "2022/06/123123120/OD/1", event},
		{"123123121", "2022/06/123123121/OD/1", event},
		{"123123122", "2022/06/123123122/OD/3", event},
		{"123123123", "2022/06/123123123/OD/1", event},
		{"shouldFail", "2022/06/shouldFail/OD/1", event},
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Warnf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	invoiceRepoMock := mocks.NewInvoiceRepo(suite.T())
	invoiceRepoMock.On("CountRepurchaseInvoices", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRepurchaseResultFn())

	numberGenerator := generators.NewRepurchaseInvoiceNumberGenerator(invoiceRepoMock, loggerMock, suite.Cfg)

	for _, test := range tests {
		actualNumber, err := numberGenerator.GetNumber(ctx, test.customerId, test.event)
		require.Equal(test.expectedInvoiceNumber, actualNumber)
		require.NoError(err, "should not happen")
	}
}
