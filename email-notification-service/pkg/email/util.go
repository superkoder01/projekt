package email

import (
	"bytes"
	"encoding/json"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
)

func ToBytes(email *model.Email) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(email)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
