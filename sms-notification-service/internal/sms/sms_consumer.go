package sms

type SmsConsumer interface {
	Execute(message []byte, messageId, consumerName string) error
}
