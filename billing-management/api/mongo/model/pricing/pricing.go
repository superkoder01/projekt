package pricing

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type Pricing struct {
	Id      string         `json:"id,omitempty" bson:"_id,omitempty"`
	Header  billing.Header `json:"header" bson:"header"`
	Payload PricingPayload `json:"payload" bson:"payload"`
}

func (c *Pricing) String() string {
	return fmt.Sprintf("%s", *c)
}
