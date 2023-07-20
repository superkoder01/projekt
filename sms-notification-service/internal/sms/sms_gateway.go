package sms

import (
	"NotificationSmsService/internal/domain/model"
)

type SmsGateway interface {
	Send(sms *model.SmsWrapper) error
}
