package model

import (
	"encoding/xml"
)

type Sms struct {
	XMLName xml.Name `xml:"sendSms"`
	Host    string   `xml:"xmlns,attr"`
	Sender  string   `xml:"sender>ccn"`
	Msisdn  []string `xml:"recipient>msisdn"`
	Text    string   `xml:"text"`
}

func SmsFromMessage(wrapper *SmsWrapper, host, sender string) *Sms {
	return &Sms{
		Text:   wrapper.Message.Text,
		Msisdn: wrapper.Message.Msisdn,
		Host:   host,
		Sender: sender,
	}
}

func SerializeSmsToXml(sms *SmsWrapper, apiHost, sender string) ([]byte, error) {
	output, err := xml.Marshal(SmsFromMessage(sms, apiHost, sender))
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func CreateWrappedSms(b []byte, messageId, workerName string) (*SmsWrapper, error) {
	sms, err := deserializeSmsFromJson(b)

	if err != nil {
		return nil, err
	} else {
		return &SmsWrapper{
			MessageWrapper: MessageWrapper{
				messageId: messageId,
				worker:    workerName,
			},
			Message: sms,
		}, nil
	}
}
