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
