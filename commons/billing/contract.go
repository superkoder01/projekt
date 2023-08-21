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

type Contract struct {
	Id      string          `bson:"_id"`
	Header  Header          `json:"header" bson:"header"`
	Payload ContractPayload `json:"payload" bson:"payload"`
}

type ContractPayload struct {
	ContractDetails     ContractDetails       `json:"contractDetails" bson:"contractDetails,omitempty"`
	SellerDetails       SellerDetails         `json:"sellerDtls" bson:"sellerDtls,omitempty"`
	CustomerDetails     CustomerDetails       `json:"customerDtls" bson:"customerDtls,omitempty"`
	ContractConditions  ContractConditions    `json:"conditions" bson:"conditions"`
	ServiceAccessPoints []ServiceAccessPoints `json:"serviceAccessPoints" bson:"serviceAccessPoints"`
	PriceList           []PriceList           `json:"priceList,omitempty" bson:"priceList,omitempty"`
	Repurchase          Repurchase            `json:"repurchase,omitempty" bson:"repurchase,omitempty"`
}

type ContractDetails struct {
	Title         string `json:"title" bson:"title,omitempty"`
	TypeName      string `json:"type" bson:"type,omitempty"`
	ClientType    string `json:"clientType" bson:"clientType,omitempty"`
	OfferId       string `json:"offerId,omitempty" bson:"offerId,omitempty"`
	TpaParameter  string `json:"tpaParameter" bson:"tpaParameter,omitempty"`
	Number        string `json:"number" bson:"number,omitempty"`
	CreationDate  string `json:"creationDate" bson:"creationDate,omitempty"`
	State         string `json:"state" bson:"state,omitempty"`
	CustomerId    string `json:"customerId" bson:"customerId,omitempty"`
	TransactionId string `json:"transactionId" bson:"transactionId,omitempty"`
	ReferenceId   string `json:"referenceId" bson:"referenceId,omitempty"`
	TariffGroup   string `json:"tariffGroup" bson:"tariffGroup,omitempty"`
	AgreementType string `json:"agreementType" bson:"agreementType,omitempty"`
}

type ContractConditions struct {
	SignatureDate                         string   `json:"signatureDate" bson:"signatureDate,omitempty"`
	StartDate                             string   `json:"startDate" bson:"startDate"`
	EndDate                               string   `json:"endDate" bson:"endDate"`
	Duration                              Duration `json:"duration" bson:"duration"`
	BillingPeriod                         Duration `json:"billingPeriod" bson:"billingPeriod"`
	InvoiceDueDate                        string   `json:"invoiceDueDate" bson:"invoiceDueDate"`
	EstimatedAnnualElectricityConsumption Energy   `json:"estimatedAnnualElectricityConsumption,omitempty" bson:"estimatedAnnualElectricityConsumption,omitempty"`
	EstimatedAnnualElectricityProduction  Energy   `json:"estimatedAnnualElectricityProduction,omitempty" bson:"estimatedAnnualElectricityProduction,omitempty"`
}

type ServiceAccessPoints struct {
	ObjectName           string        `json:"objectName" bson:"objectName"`
	Address              string        `json:"address" bson:"address"`
	SapCode              string        `json:"sapCode" bson:"sapCode"`
	MeterNumber          string        `json:"meterNumber" bson:"meterNumber"`
	Osd                  Osd           `json:"osd" bson:"osd"`
	TariffGroup          string        `json:"tariffGroup" bson:"tariffGroup"`
	EstimatedEnergyUsage Energy        `json:"estimatedEnergyUsage" bson:"estimatedEnergyUsage"`
	DeclaredEnergyUsage  Energy        `json:"declaredEnergyUsage" bson:"declaredEnergyUsage"`
	ConnectionPower      Energy        `json:"connectionPower" bson:"connectionPower"`
	ContractedPower      Energy        `json:"contractedPower" bson:"contractedPower"`
	CurrentSeller        CurrentSeller `json:"currentSeller" bson:"currentSeller"`
	Phase                string        `json:"phase" bson:"phase"`
	SourceType           string        `json:"sourceType" bson:"sourceType"`
	SourcePower          Energy        `json:"sourcePower" bson:"sourcePower"`
}

type Osd struct {
	Name   string `json:"name" bson:"name"`
	Branch string `json:"branch,omitempty" bson:"branch,omitempty"`
}

type CurrentSeller struct {
	Name         string `json:"name" bson:"name"`
	NoticePeriod string `json:"noticePeriod" bson:"noticePeriod"`
}

type PriceList struct {
	Name          string        `json:"name" bson:"name"`
	Id            string        `json:"objectName" bson:"id"`
	Type          string        `json:"type" bson:"type,omitempty"`
	StartDate     string        `json:"startDate" bson:"startDate,omitempty"`
	EndDate       string        `json:"endDate" bson:"endDate,omitempty"`
	Osd           string        `json:"osd" bson:"osd,omitempty"`
	TariffGroup   string        `json:"tariffGroup" bson:"tariffGroup,omitempty"`
	Zones         []Zone        `json:"zones" bson:"zones,omitempty"`
	CommercialFee CommercialFee `json:"commercialFee" bson:"commercialFee,omitempty"`
}

type Repurchase struct {
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
	Id    string `json:"id" bson:"id"`
	Price Price  `json:"price,omitempty" bson:"price,omitempty"`
}

type Zone struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Unit     string `json:"unit,omitempty" bson:"unit,omitempty"`
	Cost     string `json:"cost,omitempty" bson:"cost,omitempty"`
	Currency string `json:"currency,omitempty" bson:"currency,omitempty"`
}

type Price struct {
	Unit     string `json:"unit,omitempty" bson:"unit,omitempty"`
	Cost     string `json:"cost,omitempty" bson:"cost,omitempty"`
	Currency string `json:"currency,omitempty" bson:"currency,omitempty"`
}

type Energy struct {
	Unit   string `json:"unit,omitempty" bson:"unit,omitempty"`
	Amount string `json:"amount,omitempty" bson:"amount,omitempty"`
}

type Duration struct {
	CalendarUnit string `json:"calendarUnit,omitempty" bson:"calendarUnit,omitempty"`
	Number       string `json:"number,omitempty" bson:"number,omitempty"`
}

type CommercialFee struct {
	CalendarUnit string `json:"calendarUnit" bson:"calendarUnit,omitempty"`
	Cost         string `json:"cost,omitempty" bson:"cost,omitempty"`
	Currency     string `json:"currency,omitempty" bson:"currency,omitempty"`
}
