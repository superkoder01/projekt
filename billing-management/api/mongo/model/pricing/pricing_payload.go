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
