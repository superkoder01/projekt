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
