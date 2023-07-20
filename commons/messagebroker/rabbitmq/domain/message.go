package domain

/*todo create base model with builder*/

// Message used by producer to insert messages to queues
// Payload is Sms/Email object deserialized to byte slice
type Message struct {
	Payload []byte
	MessageParams
}

// MessageParams used to specify parameters for rabbitMQ Message
//contentType specifies message type (default: "application/json"),
//messageId is message identifier (optional),
//isTransient specifies if message should be persistent in queue (default: "false")
type MessageParams struct {
	ContentType string
	MessageId   string
	IsTransient bool
	Headers     map[string]interface{}
}
