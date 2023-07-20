package sms

type SmsUseCase interface {
	SendSms(message []byte, messageId, workerName string) error
}
