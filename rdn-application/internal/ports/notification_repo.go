package ports

import (
	"context"
	"database/sql/driver"
)

type NotificationType string

const (
	Email NotificationType = "email"
	Sms   NotificationType = "sms"
)

func (nType *NotificationType) Scan(value interface{}) error {
	*nType = NotificationType(value.([]byte))
	return nil
}

func (nType NotificationType) Value() (driver.Value, error) {
	return string(nType), nil
}

type NotificationRepo interface {
	PublishToUser(ctx context.Context, messageBody string, userAddress []string, notificationType NotificationType) error
}

type NotificationRepoFactory interface {
	MakeService() NotificationRepo
}
