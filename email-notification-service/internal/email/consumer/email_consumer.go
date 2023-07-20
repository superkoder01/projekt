package consumer

import "NotificationEmailService/internal/email"

// EmailConsumer Message Consumer implementation for Email Service
type EmailConsumer struct {
	emailUC      email.EmailsUseCase
	consumerName string
}

func NewEmailConsumer(emailUC email.EmailsUseCase, consumerName string) *EmailConsumer {
	return &EmailConsumer{
		emailUC:      emailUC,
		consumerName: consumerName,
	}
}

func (consumer *EmailConsumer) Execute(message []byte, messageId, consumerName string) error {
	return consumer.emailUC.SendEmail(message, messageId, consumerName)
}
