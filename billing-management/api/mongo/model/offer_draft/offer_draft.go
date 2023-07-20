package offer_draft

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type OfferDraft struct {
	Id      string            `json:"id" bson:"_id,omitempty"`
	Header  billing.Header    `json:"header" bson:"header"`
	Payload OfferDraftPayload `json:"payload" bson:"payload"`
}

func (c *OfferDraft) String() string {
	return fmt.Sprintf("%s", *c)
}

type OfferDraftPayload struct {
	OfferDetails    OfferDraftDetails      `json:"offerDetails" bson:"offerDetails,omitempty"`
	OfferConditions OfferDraftConditions   `json:"conditions" bson:"conditions,omitempty"`
	PriceList       billing.OfferPriceList `json:"priceList" bson:"priceList,omitempty"`
	Repurchase      billing.Repurchase     `json:"repurchase" bson:"repurchase,omitempty"`
}

type OfferDraftDetails struct {
	Title         string `json:"title" bson:"title,omitempty"`
	Type          string `json:"type" bson:"type,omitempty"`
	CreationDate  string `json:"creationDate" bson:"creationDate,omitempty"`
	TariffGroup   string `json:"tariffGroup" bson:"tariffGroup,omitempty"`
	AgreementType string `json:"agreementType" bson:"agreementType,omitempty"`
}

type OfferDraftConditions struct {
	StartDate      string           `json:"startDate" bson:"startDate"`
	EndDate        string           `json:"endDate" bson:"endDate"`
	Duration       billing.Duration `json:"duration" bson:"duration"`
	BillingPeriod  billing.Duration `json:"billingPeriod" bson:"billingPeriod"`
	InvoiceDueDate string           `json:"invoiceDueDate" bson:"invoiceDueDate"`
}
