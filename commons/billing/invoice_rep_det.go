package billing

import (
	"encoding/json"
	"fmt"
)

type InvoiceProsumentRepurchaseDetails struct {
	Header  Header                            `json:"header" bson:"header"`
	Payload ProsumentRepurchaseDetailsPayload `json:"payload" bson:"payload"`
}

type ProsumentRepurchaseDetailsPayload struct {
	InvoiceDetails    InvoiceDetails   `json:"invoiceDetails" bson:"invoiceDetails"`
	RepurchaseDetails []RdnMeterRecord `json:"repurchaseDetails" bson:"repurchaseDetails"`
}

func (i InvoiceProsumentRepurchaseDetails) String() string {
	data, _ := json.MarshalIndent(i, "", "\t")
	return fmt.Sprintf("%s", data)
}

func (i InvoiceProsumentRepurchaseDetails) Encode() ([]byte, error) {
	return json.Marshal(i)
}

type InvoiceProsumentRepurchaseDetailsOption func(invoiceRepurchaseDetailsProsument *InvoiceProsumentRepurchaseDetails)

func (i *InvoiceProsumentRepurchaseDetails) Create(opts ...InvoiceProsumentRepurchaseDetailsOption) {
	for _, opt := range opts {
		opt(i)
	}
}
