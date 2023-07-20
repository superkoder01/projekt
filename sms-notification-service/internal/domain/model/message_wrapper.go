package model

type MessageWrapper struct {
	messageId string
	worker    string
}

func (wrapper *MessageWrapper) GetMessageIdentifiers() string {
	return "[" + wrapper.worker + ":" + wrapper.messageId + "]"
}

func (wrapper *MessageWrapper) GetMessageId() string {
	return wrapper.messageId
}
