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
