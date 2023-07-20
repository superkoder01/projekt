package model

import "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"

type EmailWrapper struct {
	Message        *model.Email
	FileAttachment map[string][]byte
	messageId      string
	worker         string
}

func CreateWrappedEmail(b []byte, messageId, workerName string) (*EmailWrapper, error) {
	email, err := deserializeEmail(b)

	if err != nil {
		return nil, err
	} else {
		return &EmailWrapper{
			messageId:      messageId,
			worker:         workerName,
			Message:        email,
			FileAttachment: make(map[string][]byte),
		}, nil
	}
}

func (wrapper *EmailWrapper) GetMessageIdentifiers() string {
	return "[" + wrapper.worker + ":" + wrapper.messageId + "]"
}

func (wrapper *EmailWrapper) GetMessageId() string {
	return wrapper.messageId
}
