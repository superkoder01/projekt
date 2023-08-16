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
package rabbit_publisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
)

func CreateEmailMessage(destination []string, titleTemplate string, bodyTemplate string, invoice billing.InvoiceProsument) (model.Message, error) {
	var inter interface{}
	inter = invoice

	body := fmt.Sprintf(bodyTemplate, invoice.Payload.InvoiceDetails.Number, invoice.Payload.SellerDetails.LegalName)
	title := fmt.Sprintf(titleTemplate, invoice.Payload.InvoiceDetails.Number)

	var email = model.Email{
		Destination: destination,
		Title:       title,
		Body:        body,
		InvoiceData: &inter,
	}
	var be bytes.Buffer
	encoder := json.NewEncoder(&be)
	err := encoder.Encode(email)
	if err != nil {
		return model.Message{}, err
	}

	return model.Message{Payload: be.Bytes()}, nil
}
