package email

type EmailsUseCase interface {
	SendEmail(message []byte, messageId, workerName string) error
}
