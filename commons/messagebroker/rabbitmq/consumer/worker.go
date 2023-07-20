package consumer

// Worker interface have to be implemented by consumer to process incoming messages
type Worker interface {
	Execute(message []byte, messageId, consumerName string) error
}
