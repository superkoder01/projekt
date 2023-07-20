package invoice

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type Invoice struct {
	Id      string         `json:"id" bson:"_id,omitempty"`
	Header  billing.Header `json:"header" bson:"header"`
	Payload interface{}    `json:"payload" bson:"payload"`
}

func (c *Invoice) String() string {
	return fmt.Sprintf("%s", *c)
}
