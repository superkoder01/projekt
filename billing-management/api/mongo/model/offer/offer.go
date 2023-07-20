package offer

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type Offer struct {
	Id      string               `json:"id" bson:"_id,omitempty"`
	Header  billing.Header       `json:"header" bson:"header"`
	Payload billing.OfferPayload `json:"payload" bson:"payload"`
}

func (c *Offer) String() string {
	return fmt.Sprintf("%s", *c)
}
