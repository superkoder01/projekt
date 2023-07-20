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
