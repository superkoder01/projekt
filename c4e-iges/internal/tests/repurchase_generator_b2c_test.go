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
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	require2 "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/generators"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/mocks"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"os"
	"reflect"
	"testing"
	"time"
)

type RepurchaseGeneratorB2CTestSuite struct {
	suite.Suite
	ExpectedRepurchaseInvoiceNonVAT        billing.InvoiceProsumentRepurchase
	ExpectedRepurchaseInvoiceDetailsNonVAT billing.InvoiceProsumentRepurchaseDetails
	InvoiceEvent                           invoice.InvoiceEvent
	Contract                               billing.Contract
	Cfg                                    *config.AppConfig
}

func (suite *RepurchaseGeneratorB2CTestSuite) SetupSuite() {
	repurchaseInvoiceAsString := "{\n    \"header\": {\n      \"version\": \"1.0.0\",\n      \"provider\": \"keno\",\n      \"content\": {\n        \"type\": \"repurchase\",\n        \"catg\": \"prosument\"\n      }\n    },\n    \"payload\": {\n      \"invoiceDetails\": {\n        \"number\": \"2022/06/1/OD/1\",\n        \"issueDt\": \"2022-07-06T15:15:15+02:00\",\n        \"serviceDt\": \"30/06/2022\",\n        \"type\": \"nonVAT\",\n        \"customerId\": \"1\",\n        \"billingStartDt\": \"01/06/2022\",\n        \"billingEndDt\": \"30/06/2022\",\n        \"catg\": \"prosument\",\n        \"status\": \"issued\"\n      },\n      \"sellerDetails\": {\n        \"customerId\": \"1\",\n        \"firstName\": \"Grzegorz\",\n        \"lastName\": \"Sikora\",\n        \"displayName\": \"Grzegorz Sikora\",\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"contact\": {\n          \"address\": {\n            \"street\": \"Marylskiego 210\",\n            \"postCode\": \"05-825\",\n            \"city\": \"Grodzisk Mazowiecki\"\n          },\n          \"phoneNumbers\": [\n            {\n              \"type\": \"mobile\",\n              \"number\": \"+48987654321\"\n            }\n          ],\n          \"email\": \"grzegorz.sikora@wp.pl\",\n          \"www\": \"www.greg.pl\"\n        }\n      },\n      \"customerDetails\": {\n        \"legalName\": \"Keno Energia Sp. z o.o.\",\n        \"displayName\": \"Keno Energia Sp. z o.o.\",\n        \"krs\": \"0000935806\",\n        \"nip\": \"6312701187\",\n        \"regon\": \"520600217\",\n        \"address\": {\n          \"street\": \"O. Jana Siemińskiego 22\",\n          \"postCode\": \"44-100\",\n          \"city\": \"Gliwice\"\n        },\n        \"contact\": {\n          \"address\": {\n            \"street\": \"O. Jana Siemińskiego 22\",\n            \"postCode\": \"44-100\",\n            \"city\": \"Gliwice\"\n          },\n          \"phoneNumbers\": [\n            {\n              \"type\": \"mobile\",\n              \"number\": \"\"\n            },\n            {\n              \"type\": \"fix\",\n              \"number\": \"32 230 25 71\"\n            }\n          ],\n          \"email\": \"biuro@keno-energia.com\",\n          \"www\": \"www.keno-energia.com\"\n        }\n      },\n      \"ppeDetails\": [\n        {\n          \"ppeCode\": \"PL_ZEWD_1405012128_06\",\n          \"ppeName\": \"Sikora Grzegorz\",\n          \"ppeObName\": \"Sikora Grzegorz\",\n          \"address\": {\n            \"street\": \"Marylskiego 210\",\n            \"postCode\": \"05-825\",\n            \"city\": \"Grodzisk Mazowiecki\"\n          },\n          \"tariffGroup\": \"G11\",\n          \"contractedPower\": {\n            \"value\": 40,\n            \"unit\": \"kW\"\n          }\n        }\n      ],\n      \"activeEnergyConsumed\": {\n        \"energySell\": {\n          \"meters\": [\n            {\n              \"meterNumber\": \"12312312\",\n              \"items\": [\n                {\n                  \"itemName\": \"Sprzedaż energii elektrycznej\",\n                  \"itemCode\": \"ITEM_00\",\n                  \"prevMeterRead\": {\n                    \"dt\": \"01/06/2022\"\n                  },\n                  \"currMeterRead\": {\n                    \"dt\": \"30/06/2022\"\n                  },\n                  \"factor\": 1,\n                  \"consumption\": 90,\n                  \"netVal\": 30.35,\n                  \"vatRate\": 0\n                }\n              ]\n            }\n          ],\n          \"subtotal\": {\n            \"amount\": 90,\n            \"netVal\": 30.35\n          }\n        },\n        \"energyDistribution\": {\n          \"meters\": null,\n          \"subtotal\": {\n            \"amount\": 0,\n            \"netVal\": 0\n          }\n        }\n      },\n      \"sellSummary\": {\n        \"items\": [\n          {\n            \"itemName\": \"Sprzedaż energii elektrycznej\",\n            \"itemCode\": \"ITEM_00\",\n            \"netVal\": 30.35,\n            \"grossVal\": 30.35\n          }\n        ],\n        \"total\": {\n          \"netVal\": 30.35,\n          \"grossVal\": 30.35\n        }\n      }\n    }\n  }"
	json.Unmarshal([]byte(repurchaseInvoiceAsString), &suite.ExpectedRepurchaseInvoiceNonVAT)

	repurchaseInvoiceDetailsAsString := "{\n    \"header\": {\n      \"version\": \"1.0.0\",\n      \"provider\": \"keno\",\n      \"content\": {\n        \"type\": \"repurchase-details\",\n        \"catg\": \"prosument\"\n      }\n    },\n    \"payload\": {\n      \"invoiceDetails\": {\n        \"number\": \"2022/06/1/OD/1\",\n        \"issueDt\": \"2022-07-06T15:15:15+02:00\",\n        \"serviceDt\": \"30/06/2022\",\n        \"type\": \"attachment\",\n        \"customerId\": \"1\",\n        \"billingStartDt\": \"01/06/2022\",\n        \"billingEndDt\": \"30/06/2022\",\n        \"catg\": \"prosument\",\n        \"status\": \"issued\"\n      },\n      \"repurchaseDetails\": [\n        {\n          \"meterNumber\": \"12312312\",\n          \"rdnItems\": [\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"0-1\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.651,\n              \"netVal\": 1.3,\n              \"grossVal\": 1.3\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"1-2\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.5803,\n              \"netVal\": 1.16,\n              \"grossVal\": 1.16\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"2-3\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.55402,\n              \"netVal\": 1.11,\n              \"grossVal\": 1.11\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"3-4\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.5425,\n              \"netVal\": 1.09,\n              \"grossVal\": 1.09\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"4-5\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.532,\n              \"netVal\": 1.06,\n              \"grossVal\": 1.06\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"5-6\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.5383,\n              \"netVal\": 1.08,\n              \"grossVal\": 1.08\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"6-7\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.623,\n              \"netVal\": 1.25,\n              \"grossVal\": 1.25\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"7-8\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.595,\n              \"netVal\": 1.19,\n              \"grossVal\": 1.19\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"8-9\",\n              \"energyAmount\": 4,\n              \"netPrice\": 0.5712,\n              \"netVal\": 2.28,\n              \"grossVal\": 2.28\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"9-10\",\n              \"energyAmount\": 5,\n              \"netPrice\": 0.441,\n              \"netVal\": 2.21,\n              \"grossVal\": 2.21\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"10-11\",\n              \"energyAmount\": 6,\n              \"netPrice\": 0.36464,\n              \"netVal\": 2.19,\n              \"grossVal\": 2.19\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"11-12\",\n              \"energyAmount\": 7,\n              \"netPrice\": 0.3395,\n              \"netVal\": 2.38,\n              \"grossVal\": 2.38\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"12-13\",\n              \"energyAmount\": 8,\n              \"netPrice\": 0.3612,\n              \"netVal\": 2.89,\n              \"grossVal\": 2.89\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"13-14\",\n              \"energyAmount\": 4,\n              \"netPrice\": 0.31996,\n              \"netVal\": 1.28,\n              \"grossVal\": 1.28\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"14-15\",\n              \"energyAmount\": 3,\n              \"netPrice\": 0.315,\n              \"netVal\": 0.95,\n              \"grossVal\": 0.95\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"15-16\",\n              \"energyAmount\": 4,\n              \"netPrice\": 0.31996,\n              \"netVal\": 1.28,\n              \"grossVal\": 1.28\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"16-17\",\n              \"energyAmount\": 5,\n              \"netPrice\": 0.38514,\n              \"netVal\": 1.93,\n              \"grossVal\": 1.93\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"17-18\",\n              \"energyAmount\": 7,\n              \"netPrice\": 0.4263,\n              \"netVal\": 2.98,\n              \"grossVal\": 2.98\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"18-19\",\n              \"energyAmount\": 4,\n              \"netPrice\": 0.595,\n              \"netVal\": 2.38,\n              \"grossVal\": 2.38\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"19-20\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.728,\n              \"netVal\": 1.46,\n              \"grossVal\": 1.46\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"20-21\",\n              \"energyAmount\": 1,\n              \"netPrice\": 0.78365,\n              \"netVal\": 0.78,\n              \"grossVal\": 0.78\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"21-22\",\n              \"energyAmount\": 2,\n              \"netPrice\": 0.735,\n              \"netVal\": 1.47,\n              \"grossVal\": 1.47\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"22-23\",\n              \"energyAmount\": 4,\n              \"netPrice\": 0.7343,\n              \"netVal\": 2.94,\n              \"grossVal\": 2.94\n            },\n            {\n              \"date\": \"01/06/2022\",\n              \"hour\": \"23-24\",\n              \"energyAmount\": 5,\n              \"netPrice\": 0.6308,\n              \"netVal\": 3.15,\n              \"grossVal\": 3.15\n            }\n          ]\n        }\n      ]\n    }\n  }"
	json.Unmarshal([]byte(repurchaseInvoiceDetailsAsString), &suite.ExpectedRepurchaseInvoiceDetailsNonVAT)

	invoiceEventString := "{\n\t\"contract\": \"GSv0.1/4/05/07/2022\",\n\t\"startDate\": \"01/06/2022\",\n\t\"endDate\": \"30/06/2022\",\n\t\"serviceAccessPoints\": {\n\t\t\"PL_ZEWD_1405012128_06\": {\n\t\t\t\"excise\": 0.5,\n\t\t\t\"sale\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 0,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.405,\n\t\t\t\t\t\"net\": 40.5,\n\t\t\t\t\t\"gross\": 42.53,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 2.03,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 2,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 100\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 1,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 20,\n\t\t\t\t\t\"net\": 20,\n\t\t\t\t\t\"gross\": 24.6,\n\t\t\t\t\t\"vatRate\": 23,\n\t\t\t\t\t\"vat\": 4.6,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 0,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 0\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"distribution\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 4,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.75,\n\t\t\t\t\t\"net\": 0.75,\n\t\t\t\t\t\"gross\": 0.79,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.04,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 0,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 0\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 3,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 6.56,\n\t\t\t\t\t\"net\": 6.56,\n\t\t\t\t\t\"gross\": 6.89,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.33,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 0,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 0\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 7,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.33,\n\t\t\t\t\t\"net\": 0.33,\n\t\t\t\t\t\"gross\": 0.35,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.02,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 0,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 0\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 8,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.0095,\n\t\t\t\t\t\"net\": 0.95,\n\t\t\t\t\t\"gross\": 1,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.05,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 2,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 100\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 2,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.2223,\n\t\t\t\t\t\"net\": 22.23,\n\t\t\t\t\t\"gross\": 23.34,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 1.11,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 2,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 100\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 5,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.0009,\n\t\t\t\t\t\"net\": 0.09,\n\t\t\t\t\t\"gross\": 0.09,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 2,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 100\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 6,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 0.0047,\n\t\t\t\t\t\"net\": 0.47,\n\t\t\t\t\t\"gross\": 0.49,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.02,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 2,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 100\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 9,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"unitPrice\": 13.25,\n\t\t\t\t\t\"net\": 13.25,\n\t\t\t\t\t\"gross\": 13.91,\n\t\t\t\t\t\"vatRate\": 5,\n\t\t\t\t\t\"vat\": 0.66,\n\t\t\t\t\t\"readType\": 0,\n\t\t\t\t\t\"factor\": {\n\t\t\t\t\t\t\"factorType\": 0,\n\t\t\t\t\t\t\"value\": 1\n\t\t\t\t\t},\n\t\t\t\t\t\"consumed\": 0\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"repurchase\": [\n\t\t\t\t{\n\t\t\t\t\t\"itemCode\": 10,\n\t\t\t\t\t\"meterNumber\": \"12312312\",\n\t\t\t\t\t\"from\": \"01/06/2022\",\n\t\t\t\t\t\"to\": \"30/06/2022\",\n\t\t\t\t\t\"net\": 30.35,\n\t\t\t\t\t\"gross\": 30.35,\n\t\t\t\t\t\"excess\": 90\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"deposit\": {\n\t\t\t\t\"records\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 200, \"outcome\": 10, \"deposit\": 0, \"usedValue\": 0, \"residualValue\": 0},\n\t\t\t\t\t\t\"amount\": {\"income\": 200, \"outcome\": 10, \"deposit\": 0, \"usedValue\": 0, \"residualValue\": 0}\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 40, \"outcome\": 50, \"deposit\": 10, \"usedValue\": 10, \"residualValue\": 0},\n\t\t\t\t\t\t\"amount\": {\"income\": 70, \"outcome\": 50, \"deposit\": 10, \"usedValue\": 10, \"residualValue\": 0}\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 100, \"outcome\": 40, \"deposit\": 50, \"usedValue\": 50, \"residualValue\": 0},\n\t\t\t\t\t\t\"amount\": {\"income\": 100, \"outcome\": 40, \"deposit\": 50, \"usedValue\": 50, \"residualValue\": 0}\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 90, \"outcome\": 45, \"deposit\": 40, \"usedValue\": 40, \"residualValue\": 0},\n\t\t\t\t\t\t\"amount\": {\"income\": 90, \"outcome\": 45, \"deposit\": 40, \"usedValue\": 40, \"residualValue\": 0}\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 40, \"outcome\": 10.20, \"deposit\": 45, \"usedValue\": 40, \"residualValue\": 0},\n\t\t\t\t\t\t\"amount\": {\"income\": 30, \"outcome\": 30, \"deposit\": 45, \"usedValue\": 30, \"residualValue\": 0}\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"value\": {\"income\": 42.53, \"outcome\": 30.35, \"deposit\": 15.20, \"usedValue\": 15.20, \"residualValue\": 30.35},\n\t\t\t\t\t\t\"amount\": {\"income\": 100, \"outcome\": 90, \"deposit\": 45, \"usedValue\": 45, \"residualValue\": 90}\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t},\n\t\t\t\"repurchaseDetails\": [\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"0-1\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.651,\n\t\t\t\t\t\"net\": 1.3,\n\t\t\t\t\t\"gross\": 1.3\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"1-2\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.5803,\n\t\t\t\t\t\"net\": 1.16,\n\t\t\t\t\t\"gross\": 1.16\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"2-3\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.55402,\n\t\t\t\t\t\"net\": 1.11,\n\t\t\t\t\t\"gross\": 1.11\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"3-4\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.5425,\n\t\t\t\t\t\"net\": 1.09,\n\t\t\t\t\t\"gross\": 1.09\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"4-5\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.532,\n\t\t\t\t\t\"net\": 1.06,\n\t\t\t\t\t\"gross\": 1.06\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"5-6\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.5383,\n\t\t\t\t\t\"net\": 1.08,\n\t\t\t\t\t\"gross\": 1.08\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"6-7\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.623,\n\t\t\t\t\t\"net\": 1.25,\n\t\t\t\t\t\"gross\": 1.25\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"7-8\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.595,\n\t\t\t\t\t\"net\": 1.19,\n\t\t\t\t\t\"gross\": 1.19\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"8-9\",\n\t\t\t\t\t\"units\": 4,\n\t\t\t\t\t\"unitPrice\": 0.5712,\n\t\t\t\t\t\"net\": 2.28,\n\t\t\t\t\t\"gross\": 2.28\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"9-10\",\n\t\t\t\t\t\"units\": 5,\n\t\t\t\t\t\"unitPrice\": 0.441,\n\t\t\t\t\t\"net\": 2.21,\n\t\t\t\t\t\"gross\": 2.21\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"10-11\",\n\t\t\t\t\t\"units\": 6,\n\t\t\t\t\t\"unitPrice\": 0.36464,\n\t\t\t\t\t\"net\": 2.19,\n\t\t\t\t\t\"gross\": 2.19\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"11-12\",\n\t\t\t\t\t\"units\": 7,\n\t\t\t\t\t\"unitPrice\": 0.3395,\n\t\t\t\t\t\"net\": 2.38,\n\t\t\t\t\t\"gross\": 2.38\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"12-13\",\n\t\t\t\t\t\"units\": 8,\n\t\t\t\t\t\"unitPrice\": 0.3612,\n\t\t\t\t\t\"net\": 2.89,\n\t\t\t\t\t\"gross\": 2.89\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"13-14\",\n\t\t\t\t\t\"units\": 4,\n\t\t\t\t\t\"unitPrice\": 0.31996,\n\t\t\t\t\t\"net\": 1.28,\n\t\t\t\t\t\"gross\": 1.28\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"14-15\",\n\t\t\t\t\t\"units\": 3,\n\t\t\t\t\t\"unitPrice\": 0.315,\n\t\t\t\t\t\"net\": 0.95,\n\t\t\t\t\t\"gross\": 0.95\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"15-16\",\n\t\t\t\t\t\"units\": 4,\n\t\t\t\t\t\"unitPrice\": 0.31996,\n\t\t\t\t\t\"net\": 1.28,\n\t\t\t\t\t\"gross\": 1.28\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"16-17\",\n\t\t\t\t\t\"units\": 5,\n\t\t\t\t\t\"unitPrice\": 0.38514,\n\t\t\t\t\t\"net\": 1.93,\n\t\t\t\t\t\"gross\": 1.93\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"17-18\",\n\t\t\t\t\t\"units\": 7,\n\t\t\t\t\t\"unitPrice\": 0.4263,\n\t\t\t\t\t\"net\": 2.98,\n\t\t\t\t\t\"gross\": 2.98\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"18-19\",\n\t\t\t\t\t\"units\": 4,\n\t\t\t\t\t\"unitPrice\": 0.595,\n\t\t\t\t\t\"net\": 2.38,\n\t\t\t\t\t\"gross\": 2.38\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"19-20\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.728,\n\t\t\t\t\t\"net\": 1.46,\n\t\t\t\t\t\"gross\": 1.46\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"20-21\",\n\t\t\t\t\t\"units\": 1,\n\t\t\t\t\t\"unitPrice\": 0.78365,\n\t\t\t\t\t\"net\": 0.78,\n\t\t\t\t\t\"gross\": 0.78\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"21-22\",\n\t\t\t\t\t\"units\": 2,\n\t\t\t\t\t\"unitPrice\": 0.735,\n\t\t\t\t\t\"net\": 1.47,\n\t\t\t\t\t\"gross\": 1.47\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"22-23\",\n\t\t\t\t\t\"units\": 4,\n\t\t\t\t\t\"unitPrice\": 0.7343,\n\t\t\t\t\t\"net\": 2.94,\n\t\t\t\t\t\"gross\": 2.94\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"date\": \"01/06/2022\",\n\t\t\t\t\t\"hourPeriod\": \"23-24\",\n\t\t\t\t\t\"units\": 5,\n\t\t\t\t\t\"unitPrice\": 0.6308,\n\t\t\t\t\t\"net\": 3.15,\n\t\t\t\t\t\"gross\": 3.15\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t}\n}"
	json.Unmarshal([]byte(invoiceEventString), &suite.InvoiceEvent)

	contractString := "{\n    \"Id\": {\"$oid\": \"62c41119c6232c55691fe8a8\"},\n    \"header\": {\n      \"version\": \"1.0.0\",\n      \"provider\": \"keno\",\n      \"content\": {\n        \"type\": \"contract\",\n        \"catg\": \"prosumer\"\n      }\n    },\n    \"payload\": {\n      \"contractDetails\": {\n        \"title\": \"GSv0.1\",\n        \"type\": \"comprehensive\",\n        \"number\": \"GSv0.1/4/05/07/2022\",\n        \"creationDate\": \"05/07/2022\",\n        \"state\": \"accepted\",\n        \"customerId\": \"1\",\n        \"tariffGroup\": \"G11\",\n        \"agreementType\": \"B2C\"\n      },\n      \"sellerDtls\": {\n        \"legalName\": \"Keno Energia Sp. z o.o.\",\n        \"displayName\": \"Keno Energia Sp. z o.o.\",\n        \"krs\": \"0000935806\",\n        \"nip\": \"6312701187\",\n        \"regon\": \"520600217\",\n        \"bankAccountNumber\": \"31 2290 0005 0000 6000 4316 5540\",\n        \"address\": {\n          \"street\": \"O. Jana Siemińskiego 22\",\n          \"postCode\": \"44-100\",\n          \"city\": \"Gliwice\"\n        },\n        \"contact\": {\n          \"address\": {\n            \"street\": \"O. Jana Siemińskiego 22\",\n            \"postCode\": \"44-100\",\n            \"city\": \"Gliwice\"\n          },\n          \"phoneNumbers\": [\n            {\n              \"type\": \"mobile\",\n              \"number\": \"\"\n            },\n            {\n              \"type\": \"fix\",\n              \"number\": \"32 230 25 71\"\n            }\n          ],\n          \"email\": \"biuro@keno-energia.com\",\n          \"www\": \"www.keno-energia.com\"\n        }\n      },\n      \"customerDtls\": {\n        \"customerId\": \"1\",\n        \"firstName\": \"Grzegorz\",\n        \"lastName\": \"Sikora\",\n        \"displayName\": \"Grzegorz Sikora\",\n        \"address\": {\n          \"street\": \"Marylskiego 210\",\n          \"postCode\": \"05-825\",\n          \"city\": \"Grodzisk Mazowiecki\"\n        },\n        \"contact\": {\n          \"address\": {\n            \"street\": \"Marylskiego 210\",\n            \"postCode\": \"05-825\",\n            \"city\": \"Grodzisk Mazowiecki\"\n          },\n          \"phoneNumbers\": [\n            {\n              \"type\": \"mobile\",\n              \"number\": \"+48987654321\"\n            }\n          ],\n          \"email\": \"grzegorz.sikora@wp.pl\",\n          \"www\": \"www.greg.pl\"\n        }\n      },\n      \"conditions\": {\n        \"startDate\": \"05/07/2022\",\n        \"endDate\": \"05/07/2024\",\n        \"duration\": {\n          \"calendarUnit\": \"month\",\n          \"number\": \"24\"\n        },\n        \"billingPeriod\": {\n          \"calendarUnit\": \"month\",\n          \"number\": \"1\"\n        },\n        \"invoiceDueDate\": \"14\",\n        \"estimatedAnnualElectricityConsumption\": {\n          \"unit\": \"MWh\",\n          \"amount\": \"40\"\n        }\n      },\n      \"serviceAccessPoints\": [\n        {\n          \"objectName\": \"Sikora Grzegorz\",\n          \"address\": \"Chęcińska\",\n          \"sapCode\": \"PL_ZEWD_1405012128_06\",\n          \"meterNumber\": \"12312312\",\n          \"osd\": {\n            \"name\": \"ENID\"\n          },\n          \"tariffGroup\": \"G11\",\n          \"estimatedEnergyUsage\": {\n            \"unit\": \"MWh\",\n            \"amount\": \"40\"\n          },\n          \"declaredEnergyUsage\": {\n            \"unit\": \"MWh\",\n            \"amount\": \"40\"\n          },\n          \"connectionPower\": {\n            \"unit\": \"kW\",\n            \"amount\": \"40\"\n          },\n          \"contractedPower\": {\n            \"unit\": \"kW\",\n            \"amount\": \"40\"\n          },\n          \"currentSeller\": {\n            \"name\": \"PGE\",\n            \"noticePeriod\": \"3\"\n          },\n          \"phase\": \"1\",\n          \"sourceType\": \"PV\",\n          \"sourcePower\": {\n            \"unit\": \"kWp\",\n            \"amount\": \"2.42\"\n          }\n        }\n      ],\n      \"priceList\": [\n        {\n          \"name\": \"Super Cennik G11\",\n          \"id\": \"\",\n          \"type\": \"fixed\",\n          \"startDate\": \"12/04/2022\",\n          \"endDate\": \"11/04/2024\",\n          \"osd\": \"ENID\",\n          \"tariffGroup\": \"G11\",\n          \"zones\": [\n            {\n              \"id\": \"1\",\n              \"name\": \"całodobowa\",\n              \"unit\": \"kWh\",\n              \"cost\": \"232.323\",\n              \"currency\": \"pln\"\n            }\n          ],\n          \"commercialFee\": {\n            \"calendarUnit\": \"month\",\n            \"cost\": \"10.0\",\n            \"currency\": \"pln\"\n          }\n        }\n      ],\n      \"repurchase\": {\n        \"name\": \"SuperKenoOdkup\",\n        \"type\": \"rdn\",\n        \"id\": \"CEN1G11T_01.01.2022-31.03.2022\",\n        \"price\": {\n        }\n      }\n    }\n  }"
	json.Unmarshal([]byte(contractString), &suite.Contract)

	os.Setenv("CONFIGPATH", "../../config/")
	suite.Cfg, _ = config.LoadConfig()
}

func (suite *RepurchaseGeneratorB2CTestSuite) TearDownSuite() {
}

func (suite *RepurchaseGeneratorB2CTestSuite) SetupTest() {
}

func (suite *RepurchaseGeneratorB2CTestSuite) TearDownTest() {
}

func (suite *RepurchaseGeneratorB2CTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *RepurchaseGeneratorB2CTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *RepurchaseGeneratorB2CTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
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
func TestRepurchaseGeneratorB2CTestSuite(t *testing.T) {
	suite.Run(t, new(RepurchaseGeneratorB2CTestSuite))
}

///////////////////////////////////////////////////////////////////////
func (suite *RepurchaseGeneratorB2CTestSuite) TestGenerateRepurchaseInvoiceOK() {
	require := require2.New(suite.T())

	ctx := context.Background()

	tests := []struct {
		invoiceNumber                    string
		expectedRepurchaseInvoice        *billing.InvoiceProsumentRepurchase
		expectedRepurchaseInvoiceDetails *billing.InvoiceProsumentRepurchaseDetails
		invoiceEvent                     *invoice.InvoiceEvent
		contract                         *billing.Contract
	}{
		{"2022/06/1/OD/1", &suite.ExpectedRepurchaseInvoiceNonVAT, &suite.ExpectedRepurchaseInvoiceDetailsNonVAT, &suite.InvoiceEvent, &suite.Contract},
	}

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Debugf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Warnf", mock.Anything, mock.Anything, mock.Anything).Maybe()
	loggerMock.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Maybe()

	for _, test := range tests {
		repurchaseInvoiceGenerator := generators.NewRepurchaseInvoiceGeneratorB2C(test.contract, test.invoiceEvent, loggerMock, suite.Cfg)
		actualRepurchaseInvoice, actualRepurchaseInvoiceDetails, err := repurchaseInvoiceGenerator.GenerateRepurchaseInvoice(ctx, test.invoiceNumber)

		require.NoError(err, "should not happen")
		require.NotNilf(actualRepurchaseInvoice, "for b2c repurchase invoice shall be generated")
		require.NotNilf(actualRepurchaseInvoiceDetails, "for b2c repurchase invoice details shall be generated")

		// Workaround for IssueDt and PaymentDueDt that are dynamically generated
		issueDt, _ := time.Parse(time.RFC3339, "2022-07-06T15:15:15+02:00")
		actualRepurchaseInvoice.Payload.InvoiceDetails.IssueDt = issueDt
		actualRepurchaseInvoiceDetails.Payload.InvoiceDetails.IssueDt = issueDt

		require.True(reflect.DeepEqual(test.expectedRepurchaseInvoice, actualRepurchaseInvoice))
		require.True(reflect.DeepEqual(test.expectedRepurchaseInvoiceDetails, actualRepurchaseInvoiceDetails))
	}
}
