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
package billing

import (
	"time"
)

type Header struct {
	Version  string  `json:"version" bson:"version"`
	Provider string  `json:"provider" bson:"provider"`
	Content  Content `json:"content" bson:"content"`
}

type Content struct {
	Type     string `json:"type" bson:"type"`
	Category string `json:"catg" bson:"catg"`
}

type Contact struct {
	Address      Address       `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty" bson:"phoneNumbers,omitempty"`
	Email        string        `json:"email,omitempty" bson:"email,omitempty"`
	WWW          string        `json:"www,omitempty" bson:"www,omitempty"`
}

type Address struct {
	Street   string `json:"street,omitempty" bson:"street,omitempty"`
	PostCode string `json:"postCode,omitempty" bson:"postCode,omitempty"`
	City     string `json:"city,omitempty" bson:"city,omitempty"`
}

type PhoneNumber struct {
	Type   string `json:"type" bson:"type"`
	Number string `json:"number" bson:"number"`
}

type BankDetails struct {
	Account string `json:"account" bson:"account"`
}

type ContractedPower struct {
	Value float64 `json:"value" bson:"value"`
	Unit  string  `json:"unit" bson:"unit"`
}

type InvoiceDetails struct {
	Number         string    `json:"number" bson:"number"`
	IssueDt        time.Time `json:"issueDt" bson:"issueDt"`
	ServiceDt      string    `json:"serviceDt,omitempty" bson:"serviceDt,omitempty"`
	Type           string    `json:"type" bson:"type"`
	CustomerId     string    `json:"customerId" bson:"customerId"`
	BillingStartDt string    `json:"billingStartDt" bson:"billingStartDt"`
	BillingEndDt   string    `json:"billingEndDt" bson:"billingEndDt"`
	Catg           string    `json:"catg" bson:"catg"`
	Status         string    `json:"status" bson:"status"`
	TimeZone       string    `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
}

type PartyDetails struct {
	CustomerId        string  `json:"customerId,omitempty" bson:"customerId,omitempty"`
	FirstName         string  `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName          string  `json:"lastName,omitempty" bson:"lastName,omitempty"`
	LegalName         string  `json:"legalName,omitempty" bson:"legalName,omitempty"`
	DisplayName       string  `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Pesel             string  `json:"pesel,omitempty" bson:"pesel,omitempty"`
	Krs               string  `json:"krs,omitempty" bson:"krs,omitempty"`
	Nip               string  `json:"nip,omitempty" bson:"nip,omitempty"`
	Regon             string  `json:"regon,omitempty" bson:"regon,omitempty"`
	BankAccountNumber string  `json:"bankAccountNumber,omitempty" bson:"bankAccountNumber,omitempty"`
	Address           Address `json:"address,omitempty" bson:"address,omitempty"`
	Contact           Contact `json:"contact,omitempty" bson:"contact,omitempty"`
}

type SellerDetails struct {
	LegalName         string  `json:"legalName" bson:"legalName"`
	DisplayName       string  `json:"displayName" bson:"displayName"`
	Krs               string  `json:"krs,omitempty" bson:"krs,omitempty"`
	Nip               string  `json:"nip,omitempty" bson:"nip,omitempty"`
	Regon             string  `json:"regon,omitempty" bson:"regon,omitempty"`
	BankAccountNumber string  `json:"bankAccountNumber,omitempty" bson:"bankAccountNumber,omitempty"`
	Address           Address `json:"address,omitempty" bson:"address,omitempty"`
	Contact           Contact `json:"contact,omitempty" bson:"contact,omitempty"`
}

type CustomerDetails struct {
	CustomerId  string  `json:"customerId" bson:"customerId"`
	FirstName   string  `json:"firstName" bson:"firstName"`
	LastName    string  `json:"lastName" bson:"lastName"`
	DisplayName string  `json:"displayName" bson:"displayName"`
	Pesel       string  `json:"pesel,omitempty" bson:"pesel,omitempty"`
	Nip         string  `json:"nip,omitempty" bson:"nip,omitempty"`
	Regon       string  `json:"regon,omitempty" bson:"regon,omitempty"`
	Address     Address `json:"address,omitempty" bson:"address,omitempty"`
	Contact     Contact `json:"contact,omitempty" bson:"contact,omitempty"`
}

type PaymentDetails struct {
	BankDetails  BankDetails `json:"bankDetails" bson:"bankDetails"`
	PaymentTitle string      `json:"paymentTitle" bson:"paymentTitle"`
	PaymentDueDt time.Time   `json:"paymentDueDt" bson:"paymentDueDt"`
}

type PayerDetails struct {
	CustomerId  string  `json:"customerId" bson:"customerId"`
	FirstName   string  `json:"firstName" bson:"firstName"`
	LastName    string  `json:"lastName" bson:"lastName"`
	DisplayName string  `json:"displayName" bson:"displayName"`
	Nip         string  `json:"nip,omitempty" bson:"nip,omitempty"`
	Regon       string  `json:"regon,omitempty" bson:"regon,omitempty"`
	Address     Address `json:"address,omitempty" bson:"address,omitempty"`
	Contact     Contact `json:"contact,omitempty" bson:"contact,omitempty"`
}

type PpeItem struct {
	PpeCode         string          `json:"ppeCode" bson:"ppeCode"`
	PpeName         string          `json:"ppeName" bson:"ppeName"`
	PpeObName       string          `json:"ppeObName,omitempty" bson:"ppeObName,omitempty"`
	Address         Address         `json:"address,omitempty" bson:"address,omitempty"`
	TariffGroup     string          `json:"tariffGroup" bson:"tariffGroup"`
	ContractedPower ContractedPower `json:"contractedPower" bson:"contractedPower"`
}

type ActiveEnergyConsumed struct {
	EnergySell         EnergyReading `json:"energySell" bson:"energySell"`
	EnergyDistribution EnergyReading `json:"energyDistribution,omitempty" bson:"energyDistribution,omitempty"`
}

type EnergyReading struct {
	Meters    []Meter             `json:"meters" bson:"meters"`
	ExciseTax float64             `json:"exciseTax" bson:"exciseTax"`
	Subtotal  ConsumptionSubtotal `json:"subtotal" bson:"subtotal"`
}

type Meter struct {
	MeterNumber string `json:"meterNumber" bson:"meterNumber"`
	Items       []Item `json:"items" bson:"items"`
}

type Item struct {
	ItemName      string       `json:"itemName" bson:"itemName"`
	ItemCode      string       `json:"itemCode" bson:"itemCode"`
	PrevMeterRead MeterReading `json:"prevMeterRead" bson:"prevMeterRead"`
	CurrMeterRead MeterReading `json:"currMeterRead" bson:"currMeterRead"`
	Factor        float64      `json:"factor" bson:"factor"`
	Consumption   float64      `json:"consumption,omitempty" bson:"consumption,omitempty"`
	NetUnitPrice  float64      `json:"netUnitPrice,omitempty" bson:"netUnitPrice,omitempty"`
	NetVal        float64      `json:"netVal,omitempty" bson:"netVal,omitempty"`
	VatRate       int          `json:"vatRate" bson:"vatRate"`
}

type MeterReading struct {
	Date     string  `json:"dt" bson:"dt"`
	Value    float64 `json:"value,omitempty" bson:"value,omitempty"`
	ReadType string  `json:"readType,omitempty" bson:"readType,omitempty"`
}

type ConsumptionSubtotal struct {
	Amount   float64 `json:"amount" bson:"amount"`
	NetValue float64 `json:"netVal" bson:"netVal"`
}

type MeterProductionItem struct {
	ItemName   string  `json:"itemName" bson:"itemName"`
	ItemCode   string  `json:"itemCode" bson:"itemCode"`
	DateFrom   string  `json:"dateFrom" bson:"dateFrom"`
	DateTo     string  `json:"dateTo" bson:"dateTo"`
	Production float64 `json:"production" bson:"production"`
	VatRate    int     `json:"vatRate,omitempty" bson:"vatRate,omitempty"`
	NetVal     float64 `json:"netVal,omitempty" bson:"netVal,omitempty"`
	TaxVal     float64 `json:"taxVal,omitempty" bson:"taxVal,omitempty"`
	GrossVal   float64 `json:"grossVal" bson:"grossVal"`
}

type MeterProduction struct {
	MeterNumber string                `json:"meterNumber" bson:"meterNumber"`
	Items       []MeterProductionItem `json:"items" bson:"items"`
}

type MeterProductionSummary struct {
	Production float64 `json:"production" bson:"production"`
	NetVal     float64 `json:"netVal,omitempty" bson:"netVal,omitempty"`
	TaxVal     float64 `json:"taxVal,omitempty" bson:"taxVal,omitempty"`
	GrossVal   float64 `json:"grossVal" bson:"grossVal"`
}

type ActiveEnergyProduced struct {
	Meters  []MeterProduction      `json:"meters" bson:"meters"`
	Summary MeterProductionSummary `json:"summary" bson:"summary"`
}

type PpeDepositItem struct {
	PpeCode           string            `json:"ppeCode" bson:"ppeCode"`
	PpeDepositSummary PpeDepositSummary `json:"ppeSummary" bson:"ppeSummary"`
}

type PpeDeposit struct {
	Deposit         []PpeDepositItem  `json:"deposit" bson:"deposit"`
	PpeDepositTotal PpeDepositSummary `json:"ppeSummaryTotal" bson:"ppeSummaryTotal"`
}

type PpeDepositSummary struct {
	DepositCurrent  float64 `json:"depositCurrent" bson:"depositCurrent"`
	DepositConsumed float64 `json:"depositConsumed" bson:"depositConsumed"`
	DepositNext     float64 `json:"depositNext" bson:"depositNext"`
}

type ExcessSalesBalanceItem struct {
	ItemName string  `json:"itemName" bson:"itemName"`
	ItemCode string  `json:"itemCode" bson:"itemCode"`
	GrossVal float64 `json:"grossVal" bson:"grossVal"`
}

type ExcessSalesBalance struct {
	Items   []ExcessSalesBalanceItem `json:"items" bson:"items"`
	Summary Total                    `json:"summary" bson:"summary"`
}

type SellSummaryItem struct {
	ItemName string  `json:"itemName" bson:"itemName"`
	ItemCode string  `json:"itemCode" bson:"itemCode"`
	VatRate  int     `json:"vatRate,omitempty" bson:"vatRate,omitempty"`
	NetVal   float64 `json:"netVal,omitempty" bson:"netVal,omitempty"`
	TaxVal   float64 `json:"taxVal,omitempty" bson:"taxVal,omitempty"`
	GrossVal float64 `json:"grossVal" bson:"grossVal"`
}

type Total struct {
	NetVal   float64 `json:"netVal,omitempty" bson:"netVal,omitempty"`
	TaxVal   float64 `json:"taxVal,omitempty" bson:"taxVal,omitempty"`
	GrossVal float64 `json:"grossVal" bson:"grossVal"`
}

type SellSummary struct {
	Items []SellSummaryItem `json:"items" bson:"items"`
	Total Total             `json:"total" bson:"total"`
}

type PaymentSummary struct {
	Items []SellSummaryItem `json:"items" bson:"items"`
	Total Total             `json:"total" bson:"total"`
}

type EnergyAnnualBalanceItem struct {
	ItemName string    `json:"itemName" bson:"itemName"`
	ItemCode string    `json:"itemCode" bson:"itemCode"`
	Periods  []float64 `json:"periods" bson:"periods"`
}

type DepositHistory struct {
	PpeCode string                    `json:"ppeCode" bson:"ppeCode"`
	Items   []EnergyAnnualBalanceItem `json:"items" bson:"items"`
}

type EnergyAnnualBalance struct {
	History []DepositHistory `json:"history" bson:"history"`
}

type PpeSummaryItem struct {
	PpeCode        string  `json:"ppeCode" bson:"ppeCode"`
	Value          float64 `json:"value" bson:"value"`
	EnergyConsumed float64 `json:"energyConsumed" bson:"energyConsumed"`
	EnergyProduced float64 `json:"energyProduced" bson:"energyProduced"`
}

type PpeSummaryTotal struct {
	Value          float64 `json:"value" bson:"value"`
	EnergyConsumed float64 `json:"energyConsumed" bson:"energyConsumed"`
	EnergyProduced float64 `json:"energyProduced" bson:"energyProduced"`
}

type PpeSummary struct {
	Items []PpeSummaryItem `json:"items" bson:"items"`
	Total PpeSummaryTotal  `json:"total" bson:"total"`
}

type RdnMeterRecord struct {
	MeterNumber string    `json:"meterNumber" bson:"meterNumber"`
	RdnItems    []RdnItem `json:"rdnItems" bson:"rdnItems"`
}

type RdnItem struct {
	Date         string  `json:"date" bson:"date"`
	Hour         string  `json:"hour" bson:"hour"`
	EnergyAmount float64 `json:"energyAmount" bson:"energyAmount"`
	NetPrice     float64 `json:"netPrice" bson:"netPrice"`
	VatRate      int     `json:"vatRate,omitempty" bson:"vatRate,omitempty"`
	NetVal       float64 `json:"netVal" bson:"netVal"`
	TaxVal       float64 `json:"taxVal,omitempty" bson:"taxVal,omitempty"`
	GrossVal     float64 `json:"grossVal" bson:"grossVal"`
}
