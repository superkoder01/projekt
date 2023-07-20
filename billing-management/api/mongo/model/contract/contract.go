package contract

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type Contract struct {
	Id      string                  `json:"id" bson:"_id,omitempty"`
	Header  billing.Header          `json:"header" bson:"header"`
	Payload billing.ContractPayload `json:"payload" bson:"payload"`
}

func (c *Contract) String() string {
	return fmt.Sprintf("%s", *c)
}
