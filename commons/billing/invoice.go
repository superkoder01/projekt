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
