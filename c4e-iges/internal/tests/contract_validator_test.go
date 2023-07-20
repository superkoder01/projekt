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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"os"
	"testing"
)

type ContractValidatorTestSuite struct {
	suite.Suite
	Cfg *config.AppConfig
}

func (suite *ContractValidatorTestSuite) SetupSuite() {
	os.Setenv("CONFIGPATH", "../../config/")
	suite.Cfg, _ = config.LoadConfig()
}

func (suite *ContractValidatorTestSuite) TearDownSuite() {
}

func (suite *ContractValidatorTestSuite) SetupTest() {
}

func (suite *ContractValidatorTestSuite) TearDownTest() {
}

func (suite *ContractValidatorTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *ContractValidatorTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *ContractValidatorTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println("=== Suite Stats:")
	fmt.Printf("=> suite duration: %v [s]\n", stats.End.Sub(stats.Start).Seconds())
	for _, v := range stats.TestStats {
		fmt.Printf("=> %s : %v [s] : %v\n", v.TestName, v.End.Sub(v.Start).Seconds(), passed(v.Passed))
	}
}

func passed(status bool) string {
	if status {
		return fmt.Sprint("PASSED")
	}
	return fmt.Sprint("FAILED")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestContractValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ContractValidatorTestSuite))
}

///////////////////////////////////////////////////////////////////////
func (suite *ContractValidatorTestSuite) TestContractOK() {
	require := require2.New(suite.T())

	ctx := context.Background()
	contractSap := []billing.ServiceAccessPoints{{
		ObjectName:           "",
		Address:              "",
		SapCode:              "PL_ZEWD_1405012128_06",
		MeterNumber:          "",
		Osd:                  billing.Osd{},
		TariffGroup:          "",
		EstimatedEnergyUsage: billing.Energy{},
		DeclaredEnergyUsage:  billing.Energy{},
		ConnectionPower:      billing.Energy{},
		ContractedPower:      billing.Energy{},
		CurrentSeller:        billing.CurrentSeller{},
		Phase:                "",
		SourceType:           "",
		SourcePower:          billing.Energy{},
	}}
	contract := &billing.Contract{
		Id:     "123",
		Header: billing.Header{},
		Payload: billing.ContractPayload{
			ContractDetails: billing.ContractDetails{
				Title:         "",
				TypeName:      "",
				ClientType:    "",
				TpaParameter:  "",
				Number:        "GSv0.1/4/05/07/2022",
				CreationDate:  "",
				State:         "",
				CustomerId:    "",
				TransactionId: "",
				ReferenceId:   "",
				TariffGroup:   "",
				AgreementType: "",
			},
			SellerDetails: billing.SellerDetails{},
			CustomerDetails: billing.CustomerDetails{
				CustomerId:  "12312312",
				FirstName:   "",
				LastName:    "",
				DisplayName: "",
				Pesel:       "",
				Nip:         "",
				Regon:       "",
				Address:     billing.Address{},
				Contact:     billing.Contact{},
			},
			ContractConditions:  billing.ContractConditions{},
			ServiceAccessPoints: contractSap,
			PriceList:           nil,
			Repurchase:          billing.Repurchase{},
		},
	}

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

	contractValidator := validators.NewContractValidator(loggerMock, &config.AppConfig{})

	require.NoError(contractValidator.ValidateContract(ctx, invoiceEvent, contract))
}

func (suite *ContractValidatorTestSuite) TestContractNoSapInContract() {
	require := require2.New(suite.T())

	ctx := context.Background()
	contract := &billing.Contract{
		Id:     "123",
		Header: billing.Header{},
		Payload: billing.ContractPayload{
			ContractDetails: billing.ContractDetails{
				Title:         "",
				TypeName:      "",
				ClientType:    "",
				TpaParameter:  "",
				Number:        "GSv0.1/4/05/07/2022",
				CreationDate:  "",
				State:         "",
				CustomerId:    "",
				TransactionId: "",
				ReferenceId:   "",
				TariffGroup:   "",
				AgreementType: "",
			},
			SellerDetails: billing.SellerDetails{},
			CustomerDetails: billing.CustomerDetails{
				CustomerId:  "12312312",
				FirstName:   "",
				LastName:    "",
				DisplayName: "",
				Pesel:       "",
				Nip:         "",
				Regon:       "",
				Address:     billing.Address{},
				Contact:     billing.Contact{},
			},
			ContractConditions:  billing.ContractConditions{},
			ServiceAccessPoints: nil,
			PriceList:           nil,
			Repurchase:          billing.Repurchase{},
		},
	}

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
	loggerMock.On("Warnf", mock.Anything, mock.Anything, mock.Anything).Maybe()

	contractValidator := validators.NewContractValidator(loggerMock, suite.Cfg)

	require.Error(contractValidator.ValidateContract(ctx, invoiceEvent, contract))
}

func (suite *ContractValidatorTestSuite) TestContractNoSapInInvoiceEvent() {
	require := require2.New(suite.T())

	ctx := context.Background()
	contractSap := []billing.ServiceAccessPoints{{
		ObjectName:           "",
		Address:              "",
		SapCode:              "PL_ZEWD_1405012128_06",
		MeterNumber:          "",
		Osd:                  billing.Osd{},
		TariffGroup:          "",
		EstimatedEnergyUsage: billing.Energy{},
		DeclaredEnergyUsage:  billing.Energy{},
		ConnectionPower:      billing.Energy{},
		ContractedPower:      billing.Energy{},
		CurrentSeller:        billing.CurrentSeller{},
		Phase:                "",
		SourceType:           "",
		SourcePower:          billing.Energy{},
	}}
	contract := &billing.Contract{
		Id:     "123",
		Header: billing.Header{},
		Payload: billing.ContractPayload{
			ContractDetails: billing.ContractDetails{
				Title:         "",
				TypeName:      "",
				ClientType:    "",
				TpaParameter:  "",
				Number:        "GSv0.1/4/05/07/2022",
				CreationDate:  "",
				State:         "",
				CustomerId:    "",
				TransactionId: "",
				ReferenceId:   "",
				TariffGroup:   "",
				AgreementType: "",
			},
			SellerDetails: billing.SellerDetails{},
			CustomerDetails: billing.CustomerDetails{
				CustomerId:  "12312312",
				FirstName:   "",
				LastName:    "",
				DisplayName: "",
				Pesel:       "",
				Nip:         "",
				Regon:       "",
				Address:     billing.Address{},
				Contact:     billing.Contact{},
			},
			ContractConditions:  billing.ContractConditions{},
			ServiceAccessPoints: contractSap,
			PriceList:           nil,
			Repurchase:          billing.Repurchase{},
		},
	}

	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: nil,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	contractValidator := validators.NewContractValidator(loggerMock, suite.Cfg)

	require.Error(contractValidator.ValidateContract(ctx, invoiceEvent, contract))
}

func (suite *ContractValidatorTestSuite) TestContractSapCodeMismatch() {
	require := require2.New(suite.T())

	ctx := context.Background()
	contractSap := []billing.ServiceAccessPoints{{
		ObjectName:           "",
		Address:              "",
		SapCode:              "PL_ZEWD_1405012128_06",
		MeterNumber:          "",
		Osd:                  billing.Osd{},
		TariffGroup:          "",
		EstimatedEnergyUsage: billing.Energy{},
		DeclaredEnergyUsage:  billing.Energy{},
		ConnectionPower:      billing.Energy{},
		ContractedPower:      billing.Energy{},
		CurrentSeller:        billing.CurrentSeller{},
		Phase:                "",
		SourceType:           "",
		SourcePower:          billing.Energy{},
	}}
	contract := &billing.Contract{
		Id:     "123",
		Header: billing.Header{},
		Payload: billing.ContractPayload{
			ContractDetails: billing.ContractDetails{
				Title:         "",
				TypeName:      "",
				ClientType:    "",
				TpaParameter:  "",
				Number:        "GSv0.1/4/05/07/2022",
				CreationDate:  "",
				State:         "",
				CustomerId:    "",
				TransactionId: "",
				ReferenceId:   "",
				TariffGroup:   "",
				AgreementType: "",
			},
			SellerDetails: billing.SellerDetails{},
			CustomerDetails: billing.CustomerDetails{
				CustomerId:  "12312312",
				FirstName:   "",
				LastName:    "",
				DisplayName: "",
				Pesel:       "",
				Nip:         "",
				Regon:       "",
				Address:     billing.Address{},
				Contact:     billing.Contact{},
			},
			ContractConditions:  billing.ContractConditions{},
			ServiceAccessPoints: contractSap,
			PriceList:           nil,
			Repurchase:          billing.Repurchase{},
		},
	}

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_99"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Warnf", mock.Anything, mock.Anything, mock.Anything).Maybe()

	contractValidator := validators.NewContractValidator(loggerMock, suite.Cfg)

	require.Error(contractValidator.ValidateContract(ctx, invoiceEvent, contract))
}

func (suite *ContractValidatorTestSuite) TestContractNilInvoiceEvent() {
	require := require2.New(suite.T())

	ctx := context.Background()
	contractSap := []billing.ServiceAccessPoints{{
		ObjectName:           "",
		Address:              "",
		SapCode:              "PL_ZEWD_1405012128_06",
		MeterNumber:          "",
		Osd:                  billing.Osd{},
		TariffGroup:          "",
		EstimatedEnergyUsage: billing.Energy{},
		DeclaredEnergyUsage:  billing.Energy{},
		ConnectionPower:      billing.Energy{},
		ContractedPower:      billing.Energy{},
		CurrentSeller:        billing.CurrentSeller{},
		Phase:                "",
		SourceType:           "",
		SourcePower:          billing.Energy{},
	}}
	contract := &billing.Contract{
		Id:     "123",
		Header: billing.Header{},
		Payload: billing.ContractPayload{
			ContractDetails: billing.ContractDetails{
				Title:         "",
				TypeName:      "",
				ClientType:    "",
				TpaParameter:  "",
				Number:        "GSv0.1/4/05/07/2022",
				CreationDate:  "",
				State:         "",
				CustomerId:    "",
				TransactionId: "",
				ReferenceId:   "",
				TariffGroup:   "",
				AgreementType: "",
			},
			SellerDetails: billing.SellerDetails{},
			CustomerDetails: billing.CustomerDetails{
				CustomerId:  "12312312",
				FirstName:   "",
				LastName:    "",
				DisplayName: "",
				Pesel:       "",
				Nip:         "",
				Regon:       "",
				Address:     billing.Address{},
				Contact:     billing.Contact{},
			},
			ContractConditions:  billing.ContractConditions{},
			ServiceAccessPoints: contractSap,
			PriceList:           nil,
			Repurchase:          billing.Repurchase{},
		},
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	contractValidator := validators.NewContractValidator(loggerMock, suite.Cfg)

	require.Error(contractValidator.ValidateContract(ctx, nil, contract))
}

func (suite *ContractValidatorTestSuite) TestContractNilContract() {
	require := require2.New(suite.T())

	ctx := context.Background()

	invoiceEventSap := make(map[string]invoice.BillingData)
	invoiceEventSap["PL_ZEWD_1405012128_99"] = invoice.BillingData{}
	invoiceEvent := &invoice.InvoiceEvent{
		Contract:            "test-contract",
		ServiceAccessPoints: invoiceEventSap,
		StartDate:           "01/06/2022",
		EndDate:             "30/06/2022",
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything).Maybe()

	contractValidator := validators.NewContractValidator(loggerMock, suite.Cfg)

	require.Error(contractValidator.ValidateContract(ctx, invoiceEvent, nil))
}
