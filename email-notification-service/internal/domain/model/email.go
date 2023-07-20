package model

import (
	"bytes"
	"encoding/json"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
)

func deserializeEmail(b []byte) (*model.Email, error) {
	var msg model.Email
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return &msg, err
}
