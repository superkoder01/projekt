package billing

import (
	"encoding/json"
	"fmt"
)

type InvoiceProsumentRepurchase struct {
	Header  Header                     `json:"header" bson:"header"`
	Payload ProsumentRepurchasePayload `json:"payload" bson:"payload"`
}

type ProsumentRepurchasePayload struct {
	InvoiceDetails       InvoiceDetails       `json:"invoiceDetails" bson:"invoiceDetails"`
	SellerDetails        PartyDetails         `json:"sellerDetails" bson:"sellerDetails"`
	CustomerDetails      PartyDetails         `json:"customerDetails" bson:"customerDetails"`
	PpeDetails           []PpeItem            `json:"ppeDetails" bson:"ppeDetails"`
	ActiveEnergyConsumed ActiveEnergyConsumed `json:"activeEnergyConsumed" bson:"activeEnergyConsumed"`
	SellSummary          SellSummary          `json:"sellSummary" bson:"sellSummary"`
}

func (i InvoiceProsumentRepurchase) String() string {
	data, _ := json.MarshalIndent(i, "", "\t")
	return fmt.Sprintf("%s", data)
}

func (i InvoiceProsumentRepurchase) Encode() ([]byte, error) {
	return json.Marshal(i)
}

type InvoiceProsumentRepurchaseOption func(invoiceRepurchaseProsument *InvoiceProsumentRepurchase)

func (i *InvoiceProsumentRepurchase) Create(opts ...InvoiceProsumentRepurchaseOption) {
	for _, opt := range opts {
		opt(i)
	}
}
