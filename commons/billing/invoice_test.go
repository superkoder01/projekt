package billing

import (
	"encoding/json"
	"fmt"
	require2 "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"time"
)

type InvoiceTestSuite struct {
	suite.Suite
	ExpectedInvoiceNonVAT         InvoiceProsument
	ExpectedInvoiceNonVATAsString string
	ExpectedInvoiceVAT            InvoiceProsument
	ExpectedInvoiceVATAsString    string
}

func (suite *InvoiceTestSuite) SetupSuite() {
	suite.ExpectedInvoiceNonVAT = InvoiceProsument{
		Header: Header{
			Version:  "1.0.0",
			Provider: "keno",
			Content: Content{
				Type:     "invoice",
				Category: "prosument",
			},
		},
		Payload: ProsumentPayload{
			InvoiceDetails: InvoiceDetails{
				Number:         "2022/06/1728333/SP/1",
				IssueDt:        time.Date(2022, 7, 6, 15, 15, 15, 0, time.Local),
				ServiceDt:      "30/06/2022",
				Type:           "VAT",
				CustomerId:     "1728333",
				BillingStartDt: "01/06/2022",
				BillingEndDt:   "30/06/2022",
				Catg:           "prosument",
				Status:         "issued",
			},
			SellerDetails: PartyDetails{
				LegalName:   "Keno Energia Sp. z o.o.",
				DisplayName: "Keno Energia Sp. z o.o.",
				Krs:         "0000935806",
				Nip:         "6312701187",
				Regon:       "520600217",
				Address: Address{
					Street:   "O. Jana Siemińskiego 22",
					PostCode: "44-100",
					City:     "Gliwice",
				},
				Contact: Contact{
					Address: Address{
						Street:   "O. Jana Siemińskiego 22",
						PostCode: "44-100",
						City:     "Gliwice",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "",
						},
						{
							Type:   "fix",
							Number: "32 230 25 71",
						},
					},
					Email: "biuro@keno-energia.com",
					WWW:   "www.keno-energia.com",
				},
			},
			CustomerDetails: PartyDetails{
				CustomerId:  "1728333",
				FirstName:   "Grzegorz",
				LastName:    "Sikora",
				DisplayName: "Grzegorz Sikora",
				Nip:         "",
				Regon:       "",
				Address: Address{
					Street:   "Marylskiego 210",
					PostCode: "05-825",
					City:     "Grodzisk Mazowiecki",
				},
				Contact: Contact{
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "+48987654321",
						},
					},
					Email: "grzegorz.sikora@wp.pl",
					WWW:   "www.greg.com",
				},
			},
			PaymentDetails: PaymentDetails{
				BankDetails: BankDetails{
					Account: "31 2290 0005 0000 6000 4316 5540",
				},
				PaymentTitle: "Nr. Ew: 1728333",
				PaymentDueDt: time.Date(2022, 7, 20, 15, 15, 15, 0, time.Local),
			},
			PayerDetails: PartyDetails{
				CustomerId:  "1728333",
				FirstName:   "Grzegorz",
				LastName:    "Sikora",
				DisplayName: "Grzegorz Sikora",
				Nip:         "",
				Regon:       "",
				Address: Address{
					Street:   "Marylskiego 210",
					PostCode: "05-825",
					City:     "Grodzisk Mazowiecki",
				},
				Contact: Contact{
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "+48987654321",
						},
					},
					Email: "grzegorz.sikora@wp.pl",
					WWW:   "www.greg.com",
				},
			},
			PpeDetails: []PpeItem{
				{
					PpeCode:   "PL_ZEWD_1405012128_06",
					PpeName:   "Grzegorz Sikora",
					PpeObName: "Sikora Grzegorz",
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					TariffGroup: "G11",
					ContractedPower: ContractedPower{
						Value: 12,
						Unit:  "kW",
					},
				},
			},
			ActiveEnergyConsumed: ActiveEnergyConsumed{
				EnergySell: EnergyReading{
					Meters: []Meter{
						{
							MeterNumber: "72421726",
							Items: []Item{
								{
									ItemName: "Sprzedaż energii elektrycznej (A)",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  100,
									NetUnitPrice: 0.40500,
									NetVal:       40.50,
									VatRate:      5,
								},
								{
									ItemName: "Opłata handlowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    0,
										ReadType: "",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    0,
										ReadType: "",
									},
									Factor:       1,
									Consumption:  0,
									NetUnitPrice: 20.00000,
									NetVal:       20.00,
									VatRate:      23,
								},
							},
						},
					},
					ExciseTax: 0.5,
					Subtotal: ConsumptionSubtotal{
						Amount:   100,
						NetValue: 60.50,
					},
				},
				EnergyDistribution: EnergyReading{
					Meters: []Meter{
						{
							MeterNumber: "72421726",
							Items: []Item{
								{
									ItemName: "Opłata dystr. zm. całodobowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  100,
									NetUnitPrice: 0.22230,
									NetVal:       22.23,
									VatRate:      5,
								},
								{
									ItemName: "Opłata dystrybucyjna stała",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  0,
									NetUnitPrice: 6.56000,
									NetVal:       6.56,
									VatRate:      5,
								},
								{
									ItemName: "Opłata abonamentowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  0,
									NetUnitPrice: 0.75000,
									NetVal:       0.75,
									VatRate:      5,
								},
								{
									ItemName: "Opłata OZE",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  100,
									NetUnitPrice: 0.00090,
									NetVal:       0.09,
									VatRate:      5,
								},
								{
									ItemName: "Opłata kogeneracyjna",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  100,
									NetUnitPrice: 0.00470,
									NetVal:       0.47,
									VatRate:      5,
								},
								{
									ItemName: "Opłata przejściowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  0,
									NetUnitPrice: 0.33000,
									NetVal:       0.33,
									VatRate:      5,
								},
								{
									ItemName: "Opłata jakościowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  100,
									NetUnitPrice: 0.00950,
									NetVal:       0.95,
									VatRate:      5,
								},
								{
									ItemName: "Opłata mocowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:       1,
									Consumption:  0,
									NetUnitPrice: 13.25000,
									NetVal:       13.25,
									VatRate:      5,
								},
							},
						},
					},
					ExciseTax: 0,
					Subtotal: ConsumptionSubtotal{
						Amount:   100,
						NetValue: 44.63,
					},
				},
			},
			ActiveEnergyProduced: ActiveEnergyProduced{
				Meters: []MeterProduction{
					{
						MeterNumber: "72421726",
						Items: []MeterProductionItem{
							{
								ItemName:   "Odkup energii",
								ItemCode:   "",
								DateFrom:   "01/06/2022",
								DateTo:     "30/06/2022",
								Production: 90,
								VatRate:    0,
								NetVal:     0,
								TaxVal:     0,
								GrossVal:   30.35,
							},
						},
					},
				},
				Summary: MeterProductionSummary{
					Production: 90,
					NetVal:     0,
					TaxVal:     0,
					GrossVal:   30.35,
				},
			},
			DepositSummary: PpeDeposit{
				Deposit: []PpeDepositItem{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						PpeDepositSummary: PpeDepositSummary{
							DepositCurrent:  15.20,
							DepositConsumed: 15.20,
							DepositNext:     30.35,
						},
					},
				},
				PpeDepositTotal: PpeDepositSummary{
					DepositCurrent:  15.20,
					DepositConsumed: 15.20,
					DepositNext:     30.35,
				},
			},
			ExcessSalesBalance: ExcessSalesBalance{
				Items: []ExcessSalesBalanceItem{
					{
						ItemName: "Wartość sprzedaży energii w bieżącym okresie rozliczeniowym",
						ItemCode: "",
						GrossVal: 42.53,
					},
					{
						ItemName: "Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu",
						ItemCode: "",
						GrossVal: -15.20,
					},
				},
				Summary: Total{
					GrossVal: 27.33,
				},
			},
			SellSummary: SellSummary{
				Items: []SellSummaryItem{
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  5,
						NetVal:   40.50,
						TaxVal:   2.03,
						GrossVal: 42.53,
					},
					{
						ItemName: "Dystrybucja energii",
						ItemCode: "",
						VatRate:  5,
						NetVal:   44.63,
						TaxVal:   2.23,
						GrossVal: 46.86,
					},
					{
						ItemName: "Opłata handlowa",
						ItemCode: "",
						VatRate:  23,
						NetVal:   20.00,
						TaxVal:   4.60,
						GrossVal: 24.60,
					},
				},
				Total: Total{
					NetVal:   105.13,
					TaxVal:   8.86,
					GrossVal: 113.99,
				},
			},
			PaymentSummary: PaymentSummary{
				Items: []SellSummaryItem{
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  0,
						NetVal:   0,
						TaxVal:   0,
						GrossVal: 42.53,
					},
					{
						ItemName: "Uwzględniona nadwyżka",
						ItemCode: "",
						VatRate:  0,
						NetVal:   0,
						TaxVal:   0,
						GrossVal: -15.20,
					},
					{
						ItemName: "Dystrybucja energii",
						ItemCode: "",
						VatRate:  0,
						NetVal:   0,
						TaxVal:   0,
						GrossVal: 46.86,
					},
					{
						ItemName: "Opłata handlowa",
						ItemCode: "",
						VatRate:  0,
						NetVal:   0,
						TaxVal:   0,
						GrossVal: 24.60,
					},
				},
				Total: Total{
					NetVal:   0,
					TaxVal:   0,
					GrossVal: 98.79,
				},
			},
			EnergyValueAnnualBalance: EnergyAnnualBalance{
				History: []DepositHistory{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						Items: []EnergyAnnualBalanceItem{
							{
								ItemName: "Wartość energii pobranej z sieci",
								ItemCode: "",
								Periods:  []float64{200, 40, 100, 90, 40, 42.53, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Wartość energii wprowadzonej do sieci",
								ItemCode: "",
								Periods:  []float64{10, 50, 40, 45, 10.20, 30.35, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Środki na depozycie Prosumenta",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 45, 15.20, 30.35, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Środki wykorzystane z depozytu Prosumenta",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 40, 15.20, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu",
								ItemCode: "",
								Periods:  []float64{0, 0, 0, 0, 0, 30.35, 0, 0, 0, 0, 0, 0, 0},
							},
						},
					},
				},
			},
			EnergyAmountAnnualBalance: EnergyAnnualBalance{
				History: []DepositHistory{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						Items: []EnergyAnnualBalanceItem{
							{
								ItemName: "Ilość energii pobranej z sieci",
								ItemCode: "",
								Periods:  []float64{200, 70, 100, 90, 30, 100, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość energii wprowadzonej do sieci",
								ItemCode: "",
								Periods:  []float64{10, 50, 40, 45, 30, 90, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 45, 45, 90, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość wprowadzonej energii do sieci wykorzystanej",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 30, 45, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu",
								ItemCode: "",
								Periods:  []float64{0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0},
							},
						},
					},
				},
			},
			PpeSummary: PpeSummary{
				Items: []PpeSummaryItem{
					{
						PpeCode:        "PL_ZEWD_1405012128_06",
						Value:          98.79,
						EnergyConsumed: 100,
						EnergyProduced: 90,
					},
				},
				Total: PpeSummaryTotal{
					Value:          98.79,
					EnergyConsumed: 100,
					EnergyProduced: 90,
				},
			},
		},
	}
	//suite.ExpectedInvoiceNonVATAsString = "{\n\t\"header\": {\n\t\t\"version\": \"1.0.0\",\n\t\t\"provider\": \"keno\",\n\t\t\"content\": {\n\t\t\t\"type\": \"invoice\",\n\t\t\t\"catg\": \"prosument\"\n\t\t}\n\t},\n\t\"payload\": {\n\t\t\"invoiceDetails\": {\n\t\t\t\"number\": \"2022/06/1728333/SP/1\",\n\t\t\t\"issueDt\": \"06/07/2022\",\n\t\t\t\"serviceDt\": \"30/06/2022\",\n\t\t\t\"type\": \"VAT\",\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"billingStartDt\": \"01/06/2022\",\n\t\t\t\"billingEndDt\": \"30/06/2022\",\n\t\t\t\"catg\": \"prosument\",\n\t\t\t\"status\": \"issued\"\n\t\t},\n\t\t\"sellerDetails\": {\n\t\t\t\"legalName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"displayName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"krs\": \"0000935806\",\n\t\t\t\"nip\": \"6312701187\",\n\t\t\t\"regon\": \"520600217\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"\"\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"fix\",\n\t\t\t\t\t\t\"number\": \"32 230 25 71\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"biuro@keno-energia.com\",\n\t\t\t\t\"www\": \"www.keno-energia.com\"\n\t\t\t}\n\t\t},\n\t\t\"customerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"paymentDetails\": {\n\t\t\t\"bankDetails\": {\n\t\t\t\t\"account\": \"31 2290 0005 0000 6000 4316 5540\"\n\t\t\t},\n\t\t\t\"paymentTitle\": \"Nr. Ew: 1728333\",\n\t\t\t\"paymentDueDt\": \"20/07/2022\"\n\t\t},\n\t\t\"payerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"ppeDetails\": [\n\t\t\t{\n\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\"ppeName\": \"Grzegorz Sikora\",\n\t\t\t\t\"ppeObName\": \"Sikora Grzegorz\",\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"tariffGroup\": \"G11\",\n\t\t\t\t\"contractedPower\": {\n\t\t\t\t\t\"value\": 12,\n\t\t\t\t\t\"unit\": \"kW\"\n\t\t\t\t}\n\t\t\t}\n\t\t],\n\t\t\"activeEnergyConsumed\": {\n\t\t\t\"energySell\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.405,\n\t\t\t\t\t\t\t\t\"netVal\": 40.5,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 20,\n\t\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"exciseTax\": 0.5,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 60.5\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"energyDistribution\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata dystr. zm. całodobowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.2223,\n\t\t\t\t\t\t\t\t\"netVal\": 22.23,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata dystrybucyjna stała\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 6.56,\n\t\t\t\t\t\t\t\t\"netVal\": 6.56,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata abonamentowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.75,\n\t\t\t\t\t\t\t\t\"netVal\": 0.75,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata OZE\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0009,\n\t\t\t\t\t\t\t\t\"netVal\": 0.09,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata kogeneracyjna\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0047,\n\t\t\t\t\t\t\t\t\"netVal\": 0.47,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata przejściowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.33,\n\t\t\t\t\t\t\t\t\"netVal\": 0.33,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata jakościowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0095,\n\t\t\t\t\t\t\t\t\"netVal\": 0.95,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata mocowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 13.25,\n\t\t\t\t\t\t\t\t\"netVal\": 13.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 44.63\n\t\t\t\t}\n\t\t\t}\n\t\t},\n\t\t\"activeEnergyProduced\": {\n\t\t\t\"meters\": [\n\t\t\t\t{\n\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"01/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"30/06/2022\",\n\t\t\t\t\t\t\t\"production\": 90,\n\t\t\t\t\t\t\t\"grossVal\": 30.35\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"production\": 90,\n\t\t\t\t\"grossVal\": 30.35\n\t\t\t}\n\t\t},\n\t\t\"depositSummary\": {\n\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\"depositNext\": 30.35\n\t\t},\n\t\t\"excessSalesBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"grossVal\": 27.33\n\t\t\t}\n\t\t},\n\t\t\"sellSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 40.5,\n\t\t\t\t\t\"taxVal\": 2.03,\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Dystrybucja energii\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 44.63,\n\t\t\t\t\t\"taxVal\": 2.23,\n\t\t\t\t\t\"grossVal\": 46.86\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 105.13,\n\t\t\t\t\"taxVal\": 8.86,\n\t\t\t\t\"grossVal\": 113.99\n\t\t\t}\n\t\t},\n\t\t\"paymentSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Uwzględniona nadwyżka\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Dystrybucja energii\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 46.86\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"grossVal\": 98.79\n\t\t\t}\n\t\t},\n\t\t\"energyValueAnnualBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii pobranej z sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t200,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t42.53,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t10.2,\n\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki na depozycie Prosumenta\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"energyAmountAnnualBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii pobranej z sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t200,\n\t\t\t\t\t\t70,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"ppeSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"value\": 98.79,\n\t\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\t\"energyProduced\": 90\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"value\": 98.79,\n\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\"energyProduced\": 90\n\t\t\t}\n\t\t}\n\t}\n}"
	suite.ExpectedInvoiceNonVATAsString = "{\n\t\"header\": {\n\t\t\"version\": \"1.0.0\",\n\t\t\"provider\": \"keno\",\n\t\t\"content\": {\n\t\t\t\"type\": \"invoice\",\n\t\t\t\"catg\": \"prosument\"\n\t\t}\n\t},\n\t\"payload\": {\n\t\t\"invoiceDetails\": {\n\t\t\t\"number\": \"2022/06/1728333/SP/1\",\n\t\t\t\"issueDt\": \"2022-07-06T15:15:15+02:00\",\n\t\t\t\"serviceDt\": \"30/06/2022\",\n\t\t\t\"type\": \"VAT\",\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"billingStartDt\": \"01/06/2022\",\n\t\t\t\"billingEndDt\": \"30/06/2022\",\n\t\t\t\"catg\": \"prosument\",\n\t\t\t\"status\": \"issued\"\n\t\t},\n\t\t\"sellerDetails\": {\n\t\t\t\"legalName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"displayName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"krs\": \"0000935806\",\n\t\t\t\"nip\": \"6312701187\",\n\t\t\t\"regon\": \"520600217\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"\"\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"fix\",\n\t\t\t\t\t\t\"number\": \"32 230 25 71\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"biuro@keno-energia.com\",\n\t\t\t\t\"www\": \"www.keno-energia.com\"\n\t\t\t}\n\t\t},\n\t\t\"customerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"paymentDetails\": {\n\t\t\t\"bankDetails\": {\n\t\t\t\t\"account\": \"31 2290 0005 0000 6000 4316 5540\"\n\t\t\t},\n\t\t\t\"paymentTitle\": \"Nr. Ew: 1728333\",\n\t\t\t\"paymentDueDt\": \"2022-07-20T15:15:15+02:00\"\n\t\t},\n\t\t\"payerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"ppeDetails\": [\n\t\t\t{\n\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\"ppeName\": \"Grzegorz Sikora\",\n\t\t\t\t\"ppeObName\": \"Sikora Grzegorz\",\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"tariffGroup\": \"G11\",\n\t\t\t\t\"contractedPower\": {\n\t\t\t\t\t\"value\": 12,\n\t\t\t\t\t\"unit\": \"kW\"\n\t\t\t\t}\n\t\t\t}\n\t\t],\n\t\t\"activeEnergyConsumed\": {\n\t\t\t\"energySell\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.405,\n\t\t\t\t\t\t\t\t\"netVal\": 40.5,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 20,\n\t\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"exciseTax\": 0.5,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 60.5\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"energyDistribution\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata dystr. zm. całodobowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.2223,\n\t\t\t\t\t\t\t\t\"netVal\": 22.23,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata dystrybucyjna stała\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 6.56,\n\t\t\t\t\t\t\t\t\"netVal\": 6.56,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata abonamentowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.75,\n\t\t\t\t\t\t\t\t\"netVal\": 0.75,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata OZE\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0009,\n\t\t\t\t\t\t\t\t\"netVal\": 0.09,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata kogeneracyjna\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0047,\n\t\t\t\t\t\t\t\t\"netVal\": 0.47,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata przejściowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.33,\n\t\t\t\t\t\t\t\t\"netVal\": 0.33,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata jakościowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 100,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 0.0095,\n\t\t\t\t\t\t\t\t\"netVal\": 0.95,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata mocowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 13.25,\n\t\t\t\t\t\t\t\t\"netVal\": 13.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"exciseTax\": 0,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 44.63\n\t\t\t\t}\n\t\t\t}\n\t\t},\n\t\t\"activeEnergyProduced\": {\n\t\t\t\"meters\": [\n\t\t\t\t{\n\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"01/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"30/06/2022\",\n\t\t\t\t\t\t\t\"production\": 90,\n\t\t\t\t\t\t\t\"grossVal\": 30.35\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"production\": 90,\n\t\t\t\t\"grossVal\": 30.35\n\t\t\t}\n\t\t},\n\t\t\"depositSummary\": {\n\t\t\t\"deposit\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"ppeSummary\": {\n\t\t\t\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\t\t\t\"depositNext\": 30.35\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"ppeSummaryTotal\": {\n\t\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\t\"depositNext\": 30.35\n\t\t\t}\n\t\t},\n\t\t\"excessSalesBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"grossVal\": 27.33\n\t\t\t}\n\t\t},\n\t\t\"sellSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 40.5,\n\t\t\t\t\t\"taxVal\": 2.03,\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Dystrybucja energii\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 44.63,\n\t\t\t\t\t\"taxVal\": 2.23,\n\t\t\t\t\t\"grossVal\": 46.86\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 105.13,\n\t\t\t\t\"taxVal\": 8.86,\n\t\t\t\t\"grossVal\": 113.99\n\t\t\t}\n\t\t},\n\t\t\"paymentSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 42.53\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Uwzględniona nadwyżka\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Dystrybucja energii\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 46.86\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"grossVal\": 98.79\n\t\t\t}\n\t\t},\n\t\t\"energyValueAnnualBalance\": {\n\t\t\t\"history\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii pobranej z sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t200,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t42.53,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t10.2,\n\t\t\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Środki na depozycie Prosumenta\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t30.35,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"energyAmountAnnualBalance\": {\n\t\t\t\"history\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii pobranej z sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t200,\n\t\t\t\t\t\t\t\t70,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"ppeSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"value\": 98.79,\n\t\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\t\"energyProduced\": 90\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"value\": 98.79,\n\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\"energyProduced\": 90\n\t\t\t}\n\t\t}\n\t}\n}"

	suite.ExpectedInvoiceVAT = InvoiceProsument{
		Header: Header{
			Version:  "1.0.0",
			Provider: "keno",
			Content: Content{
				Type:     "invoice",
				Category: "prosument",
			},
		},
		Payload: ProsumentPayload{
			InvoiceDetails: InvoiceDetails{
				Number:         "2022/06/1728333/SP/1",
				IssueDt:        time.Date(2022, 7, 6, 15, 15, 15, 0, time.Local),
				ServiceDt:      "30/06/2022",
				Type:           "VAT",
				CustomerId:     "1728333",
				BillingStartDt: "01/06/2022",
				BillingEndDt:   "30/06/2022",
				Catg:           "prosument",
				Status:         "issued",
			},
			SellerDetails: PartyDetails{
				LegalName:   "Keno Energia Sp. z o.o.",
				DisplayName: "Keno Energia Sp. z o.o.",
				Krs:         "0000935806",
				Nip:         "6312701187",
				Regon:       "520600217",
				Address: Address{
					Street:   "O. Jana Siemińskiego 22",
					PostCode: "44-100",
					City:     "Gliwice",
				},
				Contact: Contact{
					Address: Address{
						Street:   "O. Jana Siemińskiego 22",
						PostCode: "44-100",
						City:     "Gliwice",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "",
						},
						{
							Type:   "fix",
							Number: "32 230 25 71",
						},
					},
					Email: "biuro@keno-energia.com",
					WWW:   "www.keno-energia.com",
				},
			},
			CustomerDetails: PartyDetails{
				CustomerId:  "1728333",
				FirstName:   "Grzegorz",
				LastName:    "Sikora",
				DisplayName: "Grzegorz Sikora",
				Nip:         "3232701121",
				Regon:       "345600543",
				Address: Address{
					Street:   "Marylskiego 210",
					PostCode: "05-825",
					City:     "Grodzisk Mazowiecki",
				},
				Contact: Contact{
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "+48987654321",
						},
					},
					Email: "grzegorz.sikora@wp.pl",
					WWW:   "www.greg.com",
				},
			},
			PaymentDetails: PaymentDetails{
				BankDetails: BankDetails{
					Account: "31 2290 0005 0000 6000 4316 5540",
				},
				PaymentTitle: "Nr. Ew: 1728333",
				PaymentDueDt: time.Date(2022, 7, 20, 15, 15, 15, 0, time.Local),
			},
			PayerDetails: PartyDetails{
				CustomerId:  "1728333",
				FirstName:   "Grzegorz",
				LastName:    "Sikora",
				DisplayName: "Grzegorz Sikora",
				Nip:         "3232701121",
				Regon:       "345600543",
				Address: Address{
					Street:   "Marylskiego 210",
					PostCode: "05-825",
					City:     "Grodzisk Mazowiecki",
				},
				Contact: Contact{
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					PhoneNumbers: []PhoneNumber{
						{
							Type:   "mobile",
							Number: "+48987654321",
						},
					},
					Email: "grzegorz.sikora@wp.pl",
					WWW:   "www.greg.com",
				},
			},
			PpeDetails: []PpeItem{
				{
					PpeCode:   "PL_ZEWD_1405012128_06",
					PpeName:   "Grzegorz Sikora",
					PpeObName: "Sikora Grzegorz",
					Address: Address{
						Street:   "Marylskiego 210",
						PostCode: "05-825",
						City:     "Grodzisk Mazowiecki",
					},
					TariffGroup: "G11",
					ContractedPower: ContractedPower{
						Value: 12,
						Unit:  "kW",
					},
				},
			},
			ActiveEnergyConsumed: ActiveEnergyConsumed{
				EnergySell: EnergyReading{
					Meters: []Meter{
						{
							MeterNumber: "72421726",
							Items: []Item{
								{
									ItemName: "Sprzedaż energii elektrycznej (A)",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    100,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "15/06/2022",
										Value:    150,
										ReadType: "Z",
									},
									Factor:      1,
									Consumption: 50,
									NetVal:      10.25,
									VatRate:     5,
								},
								{
									ItemName: "Sprzedaż energii elektrycznej (A)",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "15/06/2022",
										Value:    150,
										ReadType: "Z",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    200,
										ReadType: "Z",
									},
									Factor:      1,
									Consumption: 50,
									NetVal:      30.25,
									VatRate:     23,
								},
								{
									ItemName: "Opłata handlowa",
									ItemCode: "",
									PrevMeterRead: MeterReading{
										Date:     "01/06/2022",
										Value:    0,
										ReadType: "",
									},
									CurrMeterRead: MeterReading{
										Date:     "30/06/2022",
										Value:    0,
										ReadType: "",
									},
									Factor:       1,
									NetUnitPrice: 20.00,
									NetVal:       20.00,
									VatRate:      23,
								},
							},
						},
					},
					ExciseTax: 0.5,
					Subtotal: ConsumptionSubtotal{
						Amount:   100,
						NetValue: 60.50,
					},
				},
			},
			ActiveEnergyProduced: ActiveEnergyProduced{
				Meters: []MeterProduction{
					{
						MeterNumber: "72421726",
						Items: []MeterProductionItem{
							{
								ItemName:   "Odkup energii",
								ItemCode:   "",
								DateFrom:   "01/06/2022",
								DateTo:     "15/06/2022",
								Production: 45,
								VatRate:    5,
								NetVal:     20.00,
								TaxVal:     1.00,
								GrossVal:   21.00,
							},
							{
								ItemName:   "Odkup energii",
								ItemCode:   "",
								DateFrom:   "15/06/2022",
								DateTo:     "30/06/2022",
								Production: 45,
								VatRate:    23,
								NetVal:     30.00,
								TaxVal:     6.90,
								GrossVal:   36.90,
							},
						},
					},
				},
				Summary: MeterProductionSummary{
					Production: 90,
					NetVal:     50,
					TaxVal:     7.90,
					GrossVal:   57.90,
				},
			},
			DepositSummary: PpeDeposit{
				Deposit: []PpeDepositItem{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						PpeDepositSummary: PpeDepositSummary{
							DepositCurrent:  15.20,
							DepositConsumed: 15.20,
							DepositNext:     57.90,
						},
					},
				},
				PpeDepositTotal: PpeDepositSummary{
					DepositCurrent:  15.20,
					DepositConsumed: 15.20,
					DepositNext:     57.90,
				},
			},
			ExcessSalesBalance: ExcessSalesBalance{
				Items: []ExcessSalesBalanceItem{
					{
						ItemName: "Wartość sprzedaży energii w bieżącym okresie rozliczeniowym",
						ItemCode: "",
						GrossVal: 47.97,
					},
					{
						ItemName: "Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu",
						ItemCode: "",
						GrossVal: -15.20,
					},
				},
				Summary: Total{
					GrossVal: 32.77,
				},
			},
			SellSummary: SellSummary{
				Items: []SellSummaryItem{
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  5,
						NetVal:   10.25,
						TaxVal:   0.51,
						GrossVal: 10.76,
					},
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  23,
						NetVal:   30.25,
						TaxVal:   6.96,
						GrossVal: 37.21,
					},
					{
						ItemName: "Opłata handlowa",
						ItemCode: "",
						VatRate:  23,
						NetVal:   20.00,
						TaxVal:   4.60,
						GrossVal: 24.60,
					},
				},
				Total: Total{
					NetVal:   60.50,
					TaxVal:   12.07,
					GrossVal: 72.57,
				},
			},
			PaymentSummary: PaymentSummary{
				Items: []SellSummaryItem{
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  5,
						NetVal:   10.25,
						TaxVal:   0.51,
						GrossVal: 10.76,
					},
					{
						ItemName: "Sprzedaż energii elektrycznej",
						ItemCode: "",
						VatRate:  23,
						NetVal:   30.25,
						TaxVal:   6.96,
						GrossVal: 37.21,
					},
					{
						ItemName: "Uwzględniona nadwyżka",
						ItemCode: "",
						VatRate:  5,
						NetVal:   -14.48,
						TaxVal:   -0.72,
						GrossVal: -15.20,
					},
					{
						ItemName: "Opłata handlowa",
						ItemCode: "",
						VatRate:  23,
						NetVal:   20.00,
						TaxVal:   4.60,
						GrossVal: 24.60,
					},
				},
				Total: Total{
					NetVal:   46.02,
					TaxVal:   11.35,
					GrossVal: 57.37,
				},
			},
			EnergyValueAnnualBalance: EnergyAnnualBalance{
				History: []DepositHistory{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						Items: []EnergyAnnualBalanceItem{
							{
								ItemName: "Wartość energii pobranej z sieci",
								ItemCode: "",
								Periods:  []float64{200, 40, 100, 90, 40, 47.97, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Wartość energii wprowadzonej do sieci",
								ItemCode: "",
								Periods:  []float64{10, 50, 40, 45, 10.20, 57.90, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Środki na depozycie Prosumenta",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 45, 15.20, 57.90, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Środki wykorzystane z depozytu Prosumenta",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 40, 15.20, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu",
								ItemCode: "",
								Periods:  []float64{0, 0, 0, 0, 0, 57.90, 0, 0, 0, 0, 0, 0, 0},
							},
						},
					},
				},
			},
			EnergyAmountAnnualBalance: EnergyAnnualBalance{
				History: []DepositHistory{
					{
						PpeCode: "PL_ZEWD_1405012128_06",
						Items: []EnergyAnnualBalanceItem{
							{
								ItemName: "Ilość energii pobranej z sieci",
								ItemCode: "",
								Periods:  []float64{200, 70, 100, 90, 30, 100, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość energii wprowadzonej do sieci",
								ItemCode: "",
								Periods:  []float64{10, 50, 40, 45, 30, 90, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 45, 45, 90, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość wprowadzonej energii do sieci wykorzystanej",
								ItemCode: "",
								Periods:  []float64{0, 10, 50, 40, 30, 45, 0, 0, 0, 0, 0, 0, 0},
							},
							{
								ItemName: "Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu",
								ItemCode: "",
								Periods:  []float64{0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0},
							},
						},
					},
				},
			},
			PpeSummary: PpeSummary{
				Items: []PpeSummaryItem{
					{
						PpeCode:        "PL_ZEWD_1405012128_06",
						Value:          57.37,
						EnergyConsumed: 100,
						EnergyProduced: 90,
					},
				},
				Total: PpeSummaryTotal{
					Value:          57.37,
					EnergyConsumed: 100,
					EnergyProduced: 90,
				},
			},
		},
	}
	//suite.ExpectedInvoiceVATAsString = "{\n\t\"header\": {\n\t\t\"version\": \"1.0.0\",\n\t\t\"provider\": \"keno\",\n\t\t\"content\": {\n\t\t\t\"type\": \"invoice\",\n\t\t\t\"catg\": \"prosument\"\n\t\t}\n\t},\n\t\"payload\": {\n\t\t\"invoiceDetails\": {\n\t\t\t\"number\": \"2022/06/1728333/SP/1\",\n\t\t\t\"issueDt\": \"06/07/2022\",\n\t\t\t\"serviceDt\": \"30/06/2022\",\n\t\t\t\"type\": \"VAT\",\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"billingStartDt\": \"01/06/2022\",\n\t\t\t\"billingEndDt\": \"30/06/2022\",\n\t\t\t\"catg\": \"prosument\",\n\t\t\t\"status\": \"issued\"\n\t\t},\n\t\t\"sellerDetails\": {\n\t\t\t\"legalName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"displayName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"krs\": \"0000935806\",\n\t\t\t\"nip\": \"6312701187\",\n\t\t\t\"regon\": \"520600217\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"\"\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"fix\",\n\t\t\t\t\t\t\"number\": \"32 230 25 71\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"biuro@keno-energia.com\",\n\t\t\t\t\"www\": \"www.keno-energia.com\"\n\t\t\t}\n\t\t},\n\t\t\"customerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"nip\": \"3232701121\",\n\t\t\t\"regon\": \"345600543\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"paymentDetails\": {\n\t\t\t\"bankDetails\": {\n\t\t\t\t\"account\": \"31 2290 0005 0000 6000 4316 5540\"\n\t\t\t},\n\t\t\t\"paymentTitle\": \"Nr. Ew: 1728333\",\n\t\t\t\"paymentDueDt\": \"20/07/2022\"\n\t\t},\n\t\t\"payerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"nip\": \"3232701121\",\n\t\t\t\"regon\": \"345600543\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"ppeDetails\": [\n\t\t\t{\n\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\"ppeName\": \"Grzegorz Sikora\",\n\t\t\t\t\"ppeObName\": \"Sikora Grzegorz\",\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"tariffGroup\": \"G11\",\n\t\t\t\t\"contractedPower\": {\n\t\t\t\t\t\"value\": 12,\n\t\t\t\t\t\"unit\": \"kW\"\n\t\t\t\t}\n\t\t\t}\n\t\t],\n\t\t\"activeEnergyConsumed\": {\n\t\t\t\"energySell\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"15/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 150,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 50,\n\t\t\t\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"15/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 150,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 50,\n\t\t\t\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 20,\n\t\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"exciseTax\": 0.5,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 60.5\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"energyDistribution\": {\n\t\t\t\t\"meters\": null,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 0,\n\t\t\t\t\t\"netVal\": 0\n\t\t\t\t}\n\t\t\t}\n\t\t},\n\t\t\"activeEnergyProduced\": {\n\t\t\t\"meters\": [\n\t\t\t\t{\n\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"01/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"15/06/2022\",\n\t\t\t\t\t\t\t\"production\": 45,\n\t\t\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\"taxVal\": 1,\n\t\t\t\t\t\t\t\"grossVal\": 21\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"15/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"30/06/2022\",\n\t\t\t\t\t\t\t\"production\": 45,\n\t\t\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\t\t\"netVal\": 30,\n\t\t\t\t\t\t\t\"taxVal\": 6.9,\n\t\t\t\t\t\t\t\"grossVal\": 36.9\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"production\": 90,\n\t\t\t\t\"netVal\": 50,\n\t\t\t\t\"taxVal\": 7.9,\n\t\t\t\t\"grossVal\": 57.9\n\t\t\t}\n\t\t},\n\t\t\"depositSummary\": {\n\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\"depositNext\": 57.9\n\t\t},\n\t\t\"excessSalesBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 47.97\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"grossVal\": 32.77\n\t\t\t}\n\t\t},\n\t\t\"sellSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\"taxVal\": 0.51,\n\t\t\t\t\t\"grossVal\": 10.76\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\"taxVal\": 6.96,\n\t\t\t\t\t\"grossVal\": 37.21\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 60.5,\n\t\t\t\t\"taxVal\": 12.07,\n\t\t\t\t\"grossVal\": 72.57\n\t\t\t}\n\t\t},\n\t\t\"paymentSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\"taxVal\": 0.51,\n\t\t\t\t\t\"grossVal\": 10.76\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\"taxVal\": 6.96,\n\t\t\t\t\t\"grossVal\": 37.21\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Uwzględniona nadwyżka\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": -14.48,\n\t\t\t\t\t\"taxVal\": -0.72,\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 46.02,\n\t\t\t\t\"taxVal\": 11.35,\n\t\t\t\t\"grossVal\": 57.37\n\t\t\t}\n\t\t},\n\t\t\"energyValueAnnualBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii pobranej z sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t200,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t47.97,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t10.2,\n\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki na depozycie Prosumenta\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"energyAmountAnnualBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii pobranej z sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t200,\n\t\t\t\t\t\t70,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t100,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t10,\n\t\t\t\t\t\t50,\n\t\t\t\t\t\t40,\n\t\t\t\t\t\t30,\n\t\t\t\t\t\t45,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t90,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0,\n\t\t\t\t\t\t0\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"ppeSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"value\": 57.37,\n\t\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\t\"energyProduced\": 90\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"value\": 57.37,\n\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\"energyProduced\": 90\n\t\t\t}\n\t\t}\n\t}\n}"
	suite.ExpectedInvoiceVATAsString = "{\n\t\"header\": {\n\t\t\"version\": \"1.0.0\",\n\t\t\"provider\": \"keno\",\n\t\t\"content\": {\n\t\t\t\"type\": \"invoice\",\n\t\t\t\"catg\": \"prosument\"\n\t\t}\n\t},\n\t\"payload\": {\n\t\t\"invoiceDetails\": {\n\t\t\t\"number\": \"2022/06/1728333/SP/1\",\n\t\t\t\"issueDt\": \"2022-07-06T15:15:15+02:00\",\n\t\t\t\"serviceDt\": \"30/06/2022\",\n\t\t\t\"type\": \"VAT\",\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"billingStartDt\": \"01/06/2022\",\n\t\t\t\"billingEndDt\": \"30/06/2022\",\n\t\t\t\"catg\": \"prosument\",\n\t\t\t\"status\": \"issued\"\n\t\t},\n\t\t\"sellerDetails\": {\n\t\t\t\"legalName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"displayName\": \"Keno Energia Sp. z o.o.\",\n\t\t\t\"krs\": \"0000935806\",\n\t\t\t\"nip\": \"6312701187\",\n\t\t\t\"regon\": \"520600217\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"O. Jana Siemińskiego 22\",\n\t\t\t\t\t\"postCode\": \"44-100\",\n\t\t\t\t\t\"city\": \"Gliwice\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"\"\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"fix\",\n\t\t\t\t\t\t\"number\": \"32 230 25 71\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"biuro@keno-energia.com\",\n\t\t\t\t\"www\": \"www.keno-energia.com\"\n\t\t\t}\n\t\t},\n\t\t\"customerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"nip\": \"3232701121\",\n\t\t\t\"regon\": \"345600543\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"paymentDetails\": {\n\t\t\t\"bankDetails\": {\n\t\t\t\t\"account\": \"31 2290 0005 0000 6000 4316 5540\"\n\t\t\t},\n\t\t\t\"paymentTitle\": \"Nr. Ew: 1728333\",\n\t\t\t\"paymentDueDt\": \"2022-07-20T15:15:15+02:00\"\n\t\t},\n\t\t\"payerDetails\": {\n\t\t\t\"customerId\": \"1728333\",\n\t\t\t\"firstName\": \"Grzegorz\",\n\t\t\t\"lastName\": \"Sikora\",\n\t\t\t\"displayName\": \"Grzegorz Sikora\",\n\t\t\t\"nip\": \"3232701121\",\n\t\t\t\"regon\": \"345600543\",\n\t\t\t\"address\": {\n\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t},\n\t\t\t\"contact\": {\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"phoneNumbers\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"type\": \"mobile\",\n\t\t\t\t\t\t\"number\": \"+48987654321\"\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"email\": \"grzegorz.sikora@wp.pl\",\n\t\t\t\t\"www\": \"www.greg.com\"\n\t\t\t}\n\t\t},\n\t\t\"ppeDetails\": [\n\t\t\t{\n\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\"ppeName\": \"Grzegorz Sikora\",\n\t\t\t\t\"ppeObName\": \"Sikora Grzegorz\",\n\t\t\t\t\"address\": {\n\t\t\t\t\t\"street\": \"Marylskiego 210\",\n\t\t\t\t\t\"postCode\": \"05-825\",\n\t\t\t\t\t\"city\": \"Grodzisk Mazowiecki\"\n\t\t\t\t},\n\t\t\t\t\"tariffGroup\": \"G11\",\n\t\t\t\t\"contractedPower\": {\n\t\t\t\t\t\"value\": 12,\n\t\t\t\t\t\"unit\": \"kW\"\n\t\t\t\t}\n\t\t\t}\n\t\t],\n\t\t\"activeEnergyConsumed\": {\n\t\t\t\"energySell\": {\n\t\t\t\t\"meters\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 100,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"15/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 150,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 50,\n\t\t\t\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 5\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"15/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 150,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\",\n\t\t\t\t\t\t\t\t\t\"value\": 200,\n\t\t\t\t\t\t\t\t\t\"readType\": \"Z\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"consumption\": 50,\n\t\t\t\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\t\"prevMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"01/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"currMeterRead\": {\n\t\t\t\t\t\t\t\t\t\"dt\": \"30/06/2022\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\"factor\": 1,\n\t\t\t\t\t\t\t\t\"netUnitPrice\": 20,\n\t\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\t\"vatRate\": 23\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t],\n\t\t\t\t\"exciseTax\": 0.5,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 100,\n\t\t\t\t\t\"netVal\": 60.5\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"energyDistribution\": {\n\t\t\t\t\"meters\": null,\n\t\t\t\t\"exciseTax\": 0,\n\t\t\t\t\"subtotal\": {\n\t\t\t\t\t\"amount\": 0,\n\t\t\t\t\t\"netVal\": 0\n\t\t\t\t}\n\t\t\t}\n\t\t},\n\t\t\"activeEnergyProduced\": {\n\t\t\t\"meters\": [\n\t\t\t\t{\n\t\t\t\t\t\"meterNumber\": \"72421726\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"01/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"15/06/2022\",\n\t\t\t\t\t\t\t\"production\": 45,\n\t\t\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\t\t\"taxVal\": 1,\n\t\t\t\t\t\t\t\"grossVal\": 21\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Odkup energii\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"dateFrom\": \"15/06/2022\",\n\t\t\t\t\t\t\t\"dateTo\": \"30/06/2022\",\n\t\t\t\t\t\t\t\"production\": 45,\n\t\t\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\t\t\"netVal\": 30,\n\t\t\t\t\t\t\t\"taxVal\": 6.9,\n\t\t\t\t\t\t\t\"grossVal\": 36.9\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"production\": 90,\n\t\t\t\t\"netVal\": 50,\n\t\t\t\t\"taxVal\": 7.9,\n\t\t\t\t\"grossVal\": 57.9\n\t\t\t}\n\t\t},\n\t\t\"depositSummary\": {\n\t\t\t\"deposit\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"ppeSummary\": {\n\t\t\t\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\t\t\t\"depositNext\": 57.9\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"ppeSummaryTotal\": {\n\t\t\t\t\"depositCurrent\": 15.2,\n\t\t\t\t\"depositConsumed\": 15.2,\n\t\t\t\t\"depositNext\": 57.9\n\t\t\t}\n\t\t},\n\t\t\"excessSalesBalance\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": 47.97\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"summary\": {\n\t\t\t\t\"grossVal\": 32.77\n\t\t\t}\n\t\t},\n\t\t\"sellSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\"taxVal\": 0.51,\n\t\t\t\t\t\"grossVal\": 10.76\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\"taxVal\": 6.96,\n\t\t\t\t\t\"grossVal\": 37.21\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 60.5,\n\t\t\t\t\"taxVal\": 12.07,\n\t\t\t\t\"grossVal\": 72.57\n\t\t\t}\n\t\t},\n\t\t\"paymentSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": 10.25,\n\t\t\t\t\t\"taxVal\": 0.51,\n\t\t\t\t\t\"grossVal\": 10.76\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Sprzedaż energii elektrycznej\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 30.25,\n\t\t\t\t\t\"taxVal\": 6.96,\n\t\t\t\t\t\"grossVal\": 37.21\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Uwzględniona nadwyżka\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"netVal\": -14.48,\n\t\t\t\t\t\"taxVal\": -0.72,\n\t\t\t\t\t\"grossVal\": -15.2\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemName\": \"Opłata handlowa\",\n\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"netVal\": 20,\n\t\t\t\t\t\"taxVal\": 4.6,\n\t\t\t\t\t\"grossVal\": 24.6\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"netVal\": 46.02,\n\t\t\t\t\"taxVal\": 11.35,\n\t\t\t\t\"grossVal\": 57.37\n\t\t\t}\n\t\t},\n\t\t\"energyValueAnnualBalance\": {\n\t\t\t\"history\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii pobranej z sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t200,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t47.97,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t10.2,\n\t\t\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Środki na depozycie Prosumenta\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t15.2,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t57.9,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"energyAmountAnnualBalance\": {\n\t\t\t\"history\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"items\": [\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii pobranej z sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t200,\n\t\t\t\t\t\t\t\t70,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t100,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t10,\n\t\t\t\t\t\t\t\t50,\n\t\t\t\t\t\t\t\t40,\n\t\t\t\t\t\t\t\t30,\n\t\t\t\t\t\t\t\t45,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n\t\t\t\t\t\t\t\"itemCode\": \"\",\n\t\t\t\t\t\t\t\"periods\": [\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t90,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0,\n\t\t\t\t\t\t\t\t0\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t\"ppeSummary\": {\n\t\t\t\"items\": [\n\t\t\t\t{\n\t\t\t\t\t\"ppeCode\": \"PL_ZEWD_1405012128_06\",\n\t\t\t\t\t\"value\": 57.37,\n\t\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\t\"energyProduced\": 90\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"total\": {\n\t\t\t\t\"value\": 57.37,\n\t\t\t\t\"energyConsumed\": 100,\n\t\t\t\t\"energyProduced\": 90\n\t\t\t}\n\t\t}\n\t}\n}"
}

func (suite *InvoiceTestSuite) TearDownSuite() {
}

func (suite *InvoiceTestSuite) SetupTest() {
}

func (suite *InvoiceTestSuite) TearDownTest() {
}

func (suite *InvoiceTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *InvoiceTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *InvoiceTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
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
func TestInvoiceTestSuite(t *testing.T) {
	suite.Run(t, new(InvoiceTestSuite))
}

///////////////////////////////////////////////////////////////////////
func (suite *InvoiceTestSuite) TestEchoPass() {
	//assert.Equal(suite.T(), 5, suite.SuiteVariable)
	//suite.Require().Equal(4, suite.TestVariable)
}

func (suite *InvoiceTestSuite) TestUnmarshalProsumentInvoiceNonVATOK() {
	require := require2.New(suite.T())

	invoiceAsString := "{\n  \"header\": {\n    \"version\": \"1.0.0\",\n    \"provider\": \"keno\",\n    \"content\": {\n      \"type\": \"invoice\",\n      \"catg\": \"prosument\"\n    }\n  },\n  \"payload\": {\n    \"invoiceDetails\": {\n      \"number\": \"2022/06/1728333/SP/1\",\n      \"issueDt\": \"2022-07-06T15:15:15+02:00\",\n      \"serviceDt\": \"30/06/2022\",\n      \"type\": \"VAT\",\n      \"customerId\": \"1728333\",\n      \"billingStartDt\": \"01/06/2022\",\n      \"billingEndDt\": \"30/06/2022\",\n      \"catg\": \"prosument\",\n      \"status\": \"issued\"\n    },\n    \"sellerDetails\": {\n      \"legalName\": \"Keno Energia Sp. z o.o.\",\n      \"displayName\": \"Keno Energia Sp. z o.o.\",\n      \"krs\": \"0000935806\",\n      \"nip\": \"6312701187\",\n      \"regon\": \"520600217\",\n      \"address\": {\n        \"street\": \"O. Jana Siemińskiego 22\",\n        \"postCode\": \"44-100\",\n        \"city\": \"Gliwice\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"O. Jana Siemińskiego 22\",\n          \"postCode\": \"44-100\",\n          \"city\": \"Gliwice\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"\"},\n          {\"type\": \"fix\", \"number\": \"32 230 25 71\"}\n        ],\n        \"email\": \"biuro@keno-energia.com\",\n        \"www\": \"www.keno-energia.com\"\n      }\n    },\n    \"customerDetails\": {\n      \"customerId\": \"1728333\",\n      \"firstName\": \"Grzegorz\",\n      \"lastName\": \"Sikora\",\n      \"displayName\": \"Grzegorz Sikora\",\n      \"address\": {\n        \"street\": \"Marylskiego 210\",\n        \"postCode\": \"05-825\",\n        \"city\": \"Grodzisk Mazowiecki\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"+48987654321\"}\n        ],\n        \"email\": \"grzegorz.sikora@wp.pl\",\n        \"www\": \"www.greg.com\"\n      }\n    },\n    \"paymentDetails\": {\n      \"bankDetails\": {\n        \"account\": \"31 2290 0005 0000 6000 4316 5540\"\n      },\n      \"paymentTitle\": \"Nr. Ew: 1728333\",\n      \"paymentDueDt\": \"2022-07-20T15:15:15+02:00\"\n    },\n    \"payerDetails\": {\n      \"customerId\": \"1728333\",\n      \"firstName\": \"Grzegorz\",\n      \"lastName\": \"Sikora\",\n      \"displayName\": \"Grzegorz Sikora\",\n      \"address\": {\n        \"street\": \"Marylskiego 210\",\n        \"postCode\": \"05-825\",\n        \"city\": \"Grodzisk Mazowiecki\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"+48987654321\"}\n        ],\n        \"email\": \"grzegorz.sikora@wp.pl\",\n        \"www\": \"www.greg.com\"\n      }\n    },\n    \"ppeDetails\": [\n      {\n        \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n        \"ppeName\": \"Grzegorz Sikora\",\n        \"ppeObName\": \"Sikora Grzegorz\",\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"tariffGroup\": \"G11\",\n        \"contractedPower\": {\"value\": 12, \"unit\": \"kW\"}\n      }\n    ],\n    \"activeEnergyConsumed\": {\n      \"energySell\": {\n        \"meters\": [\n          {\n            \"meterNumber\": \"72421726\",\n            \"items\": [\n              {\n                \"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 100,\n                \"netUnitPrice\": 0.40500,\n                \"netVal\": 40.50,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata handlowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\"},\n                \"currMeterRead\": {\"dt\": \"30/06/2022\"},\n                \"factor\": 1,\n                \"netUnitPrice\": 20.00000,\n                \"netVal\": 20.00,\n                \"vatRate\": 23\n              }\n            ]\n          }\n        ],\n        \"exciseTax\": 0.5,\n        \"subtotal\": {\"amount\": 100, \"netVal\": 60.50}\n      },\n      \"energyDistribution\": {\n        \"meters\": [\n          {\n            \"meterNumber\": \"72421726\",\n            \"items\": [\n              {\n                \"itemName\": \"Opłata dystr. zm. całodobowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 100,\n                \"netUnitPrice\": 0.22230,\n                \"netVal\": 22.23,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata dystrybucyjna stała\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"netUnitPrice\": 6.56000,\n                \"netVal\": 6.56,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata abonamentowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"netUnitPrice\": 0.75000,\n                \"netVal\": 0.75,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata OZE\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 100,\n                \"netUnitPrice\": 0.00090,\n                \"netVal\": 0.09,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata kogeneracyjna\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 100,\n                \"netUnitPrice\": 0.00470,\n                \"netVal\": 0.47,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata przejściowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"netUnitPrice\": 0.33000,\n                \"netVal\": 0.33,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata jakościowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 100,\n                \"netUnitPrice\": 0.00950,\n                \"netVal\": 0.95,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Opłata mocowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"netUnitPrice\": 13.25000,\n                \"netVal\": 13.25,\n                \"vatRate\": 5\n              }\n            ]\n          }\n        ],\n        \"subtotal\": {\"amount\": 100, \"netVal\": 44.63}\n      }\n    },\n    \"activeEnergyProduced\": {\n      \"meters\": [\n        {\n          \"meterNumber\": \"72421726\",\n          \"items\": [\n            {\n              \"itemName\": \"Odkup energii\",\n              \"itemCode\": \"\",\n              \"dateFrom\": \"01/06/2022\",\n              \"dateTo\": \"30/06/2022\",\n              \"production\": 90,\n              \"grossVal\": 30.35\n            }\n          ]\n        }\n      ],\n      \"summary\": {\n        \"production\": 90,\n        \"grossVal\": 30.35\n      }\n    },\n    \"depositSummary\": {\n      \"deposit\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"ppeSummary\": {\n            \"depositCurrent\": 15.20,\n            \"depositConsumed\": 15.20,\n            \"depositNext\": 30.35\n          }\n        }\n      ],\n      \"ppeSummaryTotal\": {\n        \"depositCurrent\": 15.20,\n        \"depositConsumed\": 15.20,\n        \"depositNext\": 30.35\n      }\n    },\n    \"excessSalesBalance\": {\n      \"items\": [\n        {\n          \"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n          \"itemCode\": \"\",\n          \"grossVal\": 42.53\n        },\n        {\n          \"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n          \"itemCode\": \"\",\n          \"grossVal\": -15.20\n        }\n      ],\n      \"summary\": {\"grossVal\": 27.33}\n    },\n    \"sellSummary\": {\n      \"items\": [\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"vatRate\": 5,\n          \"netVal\": 40.50,\n          \"taxVal\": 2.03,\n          \"grossVal\": 42.53\n        },\n        {\n          \"itemName\": \"Dystrybucja energii\",\n          \"itemCode\": \"\",\n          \"vatRate\": 5,\n          \"netVal\": 44.63,\n          \"taxVal\": 2.23,\n          \"grossVal\": 46.86\n        },\n        {\n          \"itemName\": \"Opłata handlowa\",\n          \"itemCode\": \"\",\n          \"vatRate\": 23,\n          \"netVal\": 20.00,\n          \"taxVal\": 4.60,\n          \"grossVal\": 24.60\n        }\n      ],\n      \"total\": {\"netVal\": 105.13, \"taxVal\": 8.86, \"grossVal\": 113.99}\n    },\n    \"paymentSummary\": {\n      \"items\": [\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"grossVal\": 42.53\n        },\n        {\n          \"itemName\": \"Uwzględniona nadwyżka\",\n          \"itemCode\": \"\",\n          \"grossVal\": -15.20\n        },\n        {\n          \"itemName\": \"Dystrybucja energii\",\n          \"itemCode\": \"\",\n          \"grossVal\": 46.86\n        },\n        {\n          \"itemName\": \"Opłata handlowa\",\n          \"itemCode\": \"\",\n          \"grossVal\": 24.60\n        }\n      ],\n      \"total\": {\"grossVal\": 98.79}\n    },\n    \"energyValueAnnualBalance\": {\n      \"history\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"items\": [\n            {\n              \"itemName\": \"Wartość energii pobranej z sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [200, 40, 100, 90, 40, 42.53, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Wartość energii wprowadzonej do sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [10, 50, 40, 45, 10.20, 30.35, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Środki na depozycie Prosumenta\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 45, 15.20, 30.35, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 40, 15.20, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 0, 0, 0, 0, 30.35, 0, 0, 0, 0, 0, 0, 0]\n            }\n          ]\n        }\n      ]\n    },\n    \"energyAmountAnnualBalance\": {\n      \"history\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"items\": [\n            {\n              \"itemName\": \"Ilość energii pobranej z sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [200, 70, 100, 90, 30, 100, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość energii wprowadzonej do sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [10, 50, 40, 45, 30, 90, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 45, 45, 90, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 30, 45, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0]\n            }\n          ]\n        }\n      ]\n    },\n    \"ppeSummary\": {\n      \"items\": [\n        {\"ppeCode\": \"PL_ZEWD_1405012128_06\", \"value\": 98.79, \"energyConsumed\": 100, \"energyProduced\": 90}\n      ],\n      \"total\": {\"value\": 98.79, \"energyConsumed\": 100, \"energyProduced\": 90}\n    }\n  }\n}\n"

	var invoice InvoiceProsument
	err := json.Unmarshal([]byte(invoiceAsString), &invoice)

	require.NoError(err, "should not happen")
	require.True(reflect.DeepEqual(suite.ExpectedInvoiceNonVAT, invoice))
}

func (suite *InvoiceTestSuite) TestMarshalProsumentInvoiceNonVATOK() {
	require := require2.New(suite.T())

	//arr, err := invoice.Encode()
	//buffer := new(bytes.Buffer)
	//_ = json.Compact(buffer, []byte(invoiceAsString))

	jsonInvoice := suite.ExpectedInvoiceNonVAT.String()

	//require.NoError(err, "should not happen")
	require.EqualValues(suite.ExpectedInvoiceNonVATAsString, jsonInvoice, "should be equal")

	fmt.Printf("%v\n", jsonInvoice)
}

func (suite *InvoiceTestSuite) TestUnmarshalProsumentInvoiceVATOK() {
	require := require2.New(suite.T())

	invoiceAsString := "{\n  \"header\": {\n    \"version\": \"1.0.0\",\n    \"provider\": \"keno\",\n    \"content\": {\n      \"type\": \"invoice\",\n      \"catg\": \"prosument\"\n    }\n  },\n  \"payload\": {\n    \"invoiceDetails\": {\n      \"number\": \"2022/06/1728333/SP/1\",\n      \"issueDt\": \"2022-07-06T15:15:15+02:00\",\n      \"serviceDt\": \"30/06/2022\",\n      \"type\": \"VAT\",\n      \"customerId\": \"1728333\",\n      \"billingStartDt\": \"01/06/2022\",\n      \"billingEndDt\": \"30/06/2022\",\n      \"catg\": \"prosument\",\n      \"status\": \"issued\"\n    },\n    \"sellerDetails\": {\n      \"legalName\": \"Keno Energia Sp. z o.o.\",\n      \"displayName\": \"Keno Energia Sp. z o.o.\",\n      \"krs\": \"0000935806\",\n      \"nip\": \"6312701187\",\n      \"regon\": \"520600217\",\n      \"address\": {\n        \"street\": \"O. Jana Siemińskiego 22\",\n        \"postCode\": \"44-100\",\n        \"city\": \"Gliwice\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"O. Jana Siemińskiego 22\",\n          \"postCode\": \"44-100\",\n          \"city\": \"Gliwice\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"\"},\n          {\"type\": \"fix\", \"number\": \"32 230 25 71\"}\n        ],\n        \"email\": \"biuro@keno-energia.com\",\n        \"www\": \"www.keno-energia.com\"\n      }\n    },\n    \"customerDetails\": {\n      \"customerId\": \"1728333\",\n      \"firstName\": \"Grzegorz\",\n      \"lastName\": \"Sikora\",\n      \"displayName\": \"Grzegorz Sikora\",\n      \"nip\": \"3232701121\",\n      \"regon\": \"345600543\",\n      \"address\": {\n        \"street\": \"Marylskiego 210\",\n        \"postCode\": \"05-825\",\n        \"city\": \"Grodzisk Mazowiecki\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"+48987654321\"}\n        ],\n        \"email\": \"grzegorz.sikora@wp.pl\",\n        \"www\": \"www.greg.com\"\n      }\n    },\n    \"paymentDetails\": {\n      \"bankDetails\": {\n        \"account\": \"31 2290 0005 0000 6000 4316 5540\"\n      },\n      \"paymentTitle\": \"Nr. Ew: 1728333\",\n      \"paymentDueDt\": \"2022-07-20T15:15:15+02:00\"\n    },\n    \"payerDetails\": {\n      \"customerId\": \"1728333\",\n      \"firstName\": \"Grzegorz\",\n      \"lastName\": \"Sikora\",\n      \"displayName\": \"Grzegorz Sikora\",\n      \"nip\": \"3232701121\",\n      \"regon\": \"345600543\",\n      \"address\": {\n        \"street\": \"Marylskiego 210\",\n        \"postCode\": \"05-825\",\n        \"city\": \"Grodzisk Mazowiecki\"\n      },\n      \"contact\": {\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"phoneNumbers\": [\n          {\"type\": \"mobile\", \"number\": \"+48987654321\"}\n        ],\n        \"email\": \"grzegorz.sikora@wp.pl\",\n        \"www\": \"www.greg.com\"\n      }\n    },\n    \"ppeDetails\": [\n      {\n        \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n        \"ppeName\": \"Grzegorz Sikora\",\n        \"ppeObName\": \"Sikora Grzegorz\",\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"tariffGroup\": \"G11\",\n        \"contractedPower\": {\"value\": 12, \"unit\": \"kW\"}\n      }\n    ],\n    \"activeEnergyConsumed\": {\n      \"energySell\": {\n        \"meters\": [\n          {\n            \"meterNumber\": \"72421726\",\n            \"items\": [\n              {\n                \"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\", \"value\": 100, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"15/06/2022\", \"value\": 150, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 50,\n                \"netVal\": 10.25,\n                \"vatRate\": 5\n              },\n              {\n                \"itemName\": \"Sprzedaż energii elektrycznej (A)\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"15/06/2022\", \"value\": 150, \"readType\": \"Z\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\", \"value\": 200, \"readType\": \"Z\"},\n                \"factor\": 1,\n                \"consumption\": 50,\n                \"netVal\": 30.25,\n                \"vatRate\": 23\n              },\n              {\n                \"itemName\": \"Opłata handlowa\",\n                \"itemCode\": \"\",\n                \"prevMeterRead\": { \"dt\": \"01/06/2022\" },\n                \"currMeterRead\": {\"dt\": \"30/06/2022\" },\n                \"factor\": 1,\n                \"netUnitPrice\": 20.00,\n                \"netVal\": 20.00,\n                \"vatRate\": 23\n              }\n            ]\n          }\n        ],\n        \"exciseTax\": 0.5,\n        \"subtotal\": {\"amount\": 100, \"netVal\": 60.50}\n      }\n    },\n    \"activeEnergyProduced\": {\n      \"meters\": [\n        {\n          \"meterNumber\": \"72421726\",\n          \"items\": [\n            {\n              \"itemName\": \"Odkup energii\",\n              \"itemCode\": \"\",\n              \"dateFrom\": \"01/06/2022\",\n              \"dateTo\": \"15/06/2022\",\n              \"production\": 45,\n              \"vatRate\": 5,\n              \"netVal\": 20.00,\n              \"taxVal\": 1.00,\n              \"grossVal\": 21.00\n            },\n            {\n              \"itemName\": \"Odkup energii\",\n              \"itemCode\": \"\",\n              \"dateFrom\": \"15/06/2022\",\n              \"dateTo\": \"30/06/2022\",\n              \"production\": 45,\n              \"vatRate\": 23,\n              \"netVal\": 30.00,\n              \"taxVal\": 6.90,\n              \"grossVal\": 36.90\n            }\n          ]\n        }\n      ],\n      \"summary\": {\n        \"production\": 90,\n        \"netVal\": 50.00,\n        \"taxVal\": 7.90,\n        \"grossVal\": 57.90\n      }\n    },\n    \"depositSummary\": {\n      \"deposit\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"ppeSummary\": {\n            \"depositCurrent\": 15.20,\n            \"depositConsumed\": 15.20,\n            \"depositNext\": 57.90\n          }\n        }\n      ],\n      \"ppeSummaryTotal\": {\n        \"depositCurrent\": 15.20,\n        \"depositConsumed\": 15.20,\n        \"depositNext\": 57.90\n      }\n    },\n    \"excessSalesBalance\": {\n      \"items\": [\n        {\n          \"itemName\": \"Wartość sprzedaży energii w bieżącym okresie rozliczeniowym\",\n          \"itemCode\": \"\",\n          \"grossVal\": 47.97\n        },\n        {\n          \"itemName\": \"Środki z depozytu Prosumenta wykorzystane w bieżącym rozliczeniu\",\n          \"itemCode\": \"\",\n          \"grossVal\": -15.20\n        }\n      ],\n      \"summary\": {\"grossVal\": 32.77}\n    },\n    \"sellSummary\": {\n      \"items\": [\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"vatRate\": 5,\n          \"netVal\": 10.25,\n          \"taxVal\": 0.51,\n          \"grossVal\": 10.76\n        },\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"vatRate\": 23,\n          \"netVal\": 30.25,\n          \"taxVal\": 6.96,\n          \"grossVal\": 37.21\n        },\n        {\n          \"itemName\": \"Opłata handlowa\",\n          \"itemCode\": \"\",\n          \"vatRate\": 23,\n          \"netVal\": 20.00,\n          \"taxVal\": 4.60,\n          \"grossVal\": 24.60\n        }\n      ],\n      \"total\": {\"netVal\": 60.50, \"taxVal\": 12.07, \"grossVal\": 72.57}\n    },\n    \"paymentSummary\": {\n      \"items\": [\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"vatRate\": 5,\n          \"netVal\": 10.25,\n          \"taxVal\": 0.51,\n          \"grossVal\": 10.76\n        },\n        {\n          \"itemName\": \"Sprzedaż energii elektrycznej\",\n          \"itemCode\": \"\",\n          \"vatRate\": 23,\n          \"netVal\": 30.25,\n          \"taxVal\": 6.96,\n          \"grossVal\": 37.21\n        },\n        {\n          \"itemName\": \"Uwzględniona nadwyżka\",\n          \"itemCode\": \"\",\n          \"vatRate\": 5,\n          \"netVal\": -14.48,\n          \"taxVal\": -0.72,\n          \"grossVal\": -15.20\n        },\n        {\n          \"itemName\": \"Opłata handlowa\",\n          \"itemCode\": \"\",\n          \"vatRate\": 23,\n          \"netVal\": 20.00,\n          \"taxVal\": 4.60,\n          \"grossVal\": 24.60\n        }\n      ],\n      \"total\": {\"netVal\": 46.02, \"taxVal\": 11.35, \"grossVal\": 57.37}\n    },\n    \"energyValueAnnualBalance\": {\n      \"history\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"items\": [\n            {\n              \"itemName\": \"Wartość energii pobranej z sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [200, 40, 100, 90, 40, 47.97, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Wartość energii wprowadzonej do sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [10, 50, 40, 45, 10.20, 57.90, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Środki na depozycie Prosumenta\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 45, 15.20, 57.90, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Środki wykorzystane z depozytu Prosumenta\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 40, 15.20, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Wartość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 0, 0, 0, 0, 57.90, 0, 0, 0, 0, 0, 0, 0]\n            }\n          ]\n        }\n      ]\n    },\n    \"energyAmountAnnualBalance\": {\n      \"history\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"items\": [\n            {\n              \"itemName\": \"Ilość energii pobranej z sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [200, 70, 100, 90, 30, 100, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość energii wprowadzonej do sieci\",\n              \"itemCode\": \"\",\n              \"periods\": [10, 50, 40, 45, 30, 90, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Suma energii wprowadzonej do sieci jaka pozostaje do wykorzystania\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 45, 45, 90, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość wprowadzonej energii do sieci wykorzystanej\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 10, 50, 40, 30, 45, 0, 0, 0, 0, 0, 0, 0]\n            },\n            {\n              \"itemName\": \"Ilość energii wprowadzonej do sieci jaka pozostała do wykorzystania w danym miesiącu\",\n              \"itemCode\": \"\",\n              \"periods\": [0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0]\n            }\n          ]\n        }\n      ]\n    },\n    \"ppeSummary\": {\n      \"items\": [\n        {\"ppeCode\": \"PL_ZEWD_1405012128_06\", \"value\": 57.37, \"energyConsumed\": 100, \"energyProduced\": 90}\n      ],\n      \"total\": {\"value\": 57.37, \"energyConsumed\": 100, \"energyProduced\": 90}\n    }\n  }\n}\n"

	var invoice InvoiceProsument
	err := json.Unmarshal([]byte(invoiceAsString), &invoice)

	require.NoError(err, "should not happen")
	require.True(reflect.DeepEqual(suite.ExpectedInvoiceVAT, invoice))
}

func (suite *InvoiceTestSuite) TestMarshalProsumentInvoiceVATOK() {
	require := require2.New(suite.T())

	//arr, err := invoice.Encode()
	//buffer := new(bytes.Buffer)
	//_ = json.Compact(buffer, []byte(invoiceAsString))

	jsonInvoice := suite.ExpectedInvoiceVAT.String()

	//require.NoError(err, "should not happen")
	require.EqualValues(suite.ExpectedInvoiceVATAsString, jsonInvoice, "should be equal")

	fmt.Printf("%v\n", jsonInvoice)
}
