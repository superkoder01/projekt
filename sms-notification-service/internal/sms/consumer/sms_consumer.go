package consumer

import (
	"NotificationSmsService/internal/sms"
)

// SmsConsumer Message Consumer implementation for Sms Service
type SmsConsumer struct {
	smsUc        sms.SmsUseCase
	consumerName string
}

func NewSmsConsumer(smsUc sms.SmsUseCase, consumerName string) *SmsConsumer {
	return &SmsConsumer{
		smsUc:        smsUc,
		consumerName: consumerName,
	}
}

func (consumer *SmsConsumer) Execute(message []byte, messageId, consumerName string) error {
	return consumer.smsUc.SendSms(message, messageId, consumerName)
}
