package pricing

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type PricingPayload struct {
	Name          string                       `json:"name" bson:"name"`
	Id            string                       `json:"id" bson:"id"`
	Type          string                       `json:"type" bson:"type"`
	Osd           string                       `json:"osd,omitempty" bson:"osd,omitempty"`
	TariffGroup   string                       `json:"tariffGroup,omitempty" bson:"tariffGroup,omitempty"`
	Zones         []billing.Zone               `json:"zones,omitempty" bson:"zones,omitempty"`
	CommercialFee []billing.OfferCommercialFee `json:"commercialFee,omitempty" bson:"commercialFee,omitempty"`
	Price         billing.Price                `json:"price,omitempty" bson:"price,omitempty"`
}

func (p *PricingPayload) String() string {
	return fmt.Sprintf("%s", *p)
}
