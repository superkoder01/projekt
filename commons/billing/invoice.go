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
	"encoding/json"
	"fmt"
)

type InvoiceProsument struct {
	Header  Header           `json:"header" bson:"header"`
	Payload ProsumentPayload `json:"payload" bson:"payload"`
}

type ProsumentPayload struct {
	InvoiceDetails            InvoiceDetails       `json:"invoiceDetails" bson:"invoiceDetails"`
	SellerDetails             PartyDetails         `json:"sellerDetails" bson:"sellerDetails"`
	CustomerDetails           PartyDetails         `json:"customerDetails" bson:"customerDetails"`
	PaymentDetails            PaymentDetails       `json:"paymentDetails" bson:"paymentDetails"`
	PayerDetails              PartyDetails         `json:"payerDetails" bson:"payerDetails"`
	PpeDetails                []PpeItem            `json:"ppeDetails" bson:"ppeDetails"`
	ActiveEnergyConsumed      ActiveEnergyConsumed `json:"activeEnergyConsumed" bson:"activeEnergyConsumed"`
	ActiveEnergyProduced      ActiveEnergyProduced `json:"activeEnergyProduced,omitempty" bson:"activeEnergyProduced,omitempty"`
	DepositSummary            PpeDeposit           `json:"depositSummary" bson:"depositSummary"`
	ExcessSalesBalance        ExcessSalesBalance   `json:"excessSalesBalance" bson:"excessSalesBalance"`
	SellSummary               SellSummary          `json:"sellSummary" bson:"sellSummary"`
	PaymentSummary            PaymentSummary       `json:"paymentSummary" bson:"paymentSummary"`
	EnergyValueAnnualBalance  EnergyAnnualBalance  `json:"energyValueAnnualBalance,omitempty" bson:"energyValueAnnualBalance,omitempty"`
	EnergyAmountAnnualBalance EnergyAnnualBalance  `json:"energyAmountAnnualBalance,omitempty" bson:"energyAmountAnnualBalance,omitempty"`
	PpeSummary                PpeSummary           `json:"ppeSummary,omitempty" bson:"ppeSummary,omitempty"`
}

func (i InvoiceProsument) String() string {
	data, _ := json.MarshalIndent(i, "", "\t")
	return fmt.Sprintf("%s", data)
}

func (i InvoiceProsument) Encode() ([]byte, error) {
	return json.Marshal(i)
}

type InvoiceProsumentOption func(invoiceProsument *InvoiceProsument)

func (i *InvoiceProsument) Create(opts ...InvoiceProsumentOption) {
	for _, opt := range opts {
		opt(i)
	}
}

// InvoiceSummary invoice summary
type InvoiceSummary struct {
	Header  Header                `json:"header" bson:"header"`
	Payload InvoiceSummaryPayload `json:"payload" bson:"payload"`
}

type InvoiceSummaryPayload struct {
	InvoiceDetails  InvoiceDetails  `json:"invoiceDetails" bson:"invoiceDetails"`
	CustomerDetails CustomerDetails `json:"customerDetails" bson:"customerDetails"`
	PaymentDetails  PaymentDetails  `json:"paymentDetails" bson:"paymentDetails"`
	SellSummary     SellSummary     `json:"sellSummary" bson:"sellSummary"`
}

func (i InvoiceSummary) String() string {
	data, _ := json.MarshalIndent(i, "", "\t")
	return fmt.Sprintf("%s", data)
}

func (i InvoiceSummary) Encode() ([]byte, error) {
	return json.Marshal(i)
}
