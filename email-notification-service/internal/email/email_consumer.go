package email

type EmailConsumer interface {
	Execute(message []byte, messageId, consumerName string) error
}
