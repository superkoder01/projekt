package model

import (
	"bytes"
	"encoding/json"
)

type SmsWrapper struct {
	MessageWrapper
	Message *SmsMessage
}

type SmsMessage struct {
	Msisdn []string `json:"msisdn" `
	Text   string   `json:"text"`
}

func deserializeSmsFromJson(b []byte) (*SmsMessage, error) {
	var msg SmsMessage
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return &msg, err
}
