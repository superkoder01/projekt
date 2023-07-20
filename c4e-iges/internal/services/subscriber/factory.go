package subscriber

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	rabbit_subscriber "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/subscriber/rabbit"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/worker"
)

type SubscriberType string

const (
	RABBITMQ SubscriberType = "rabbitmq"
)

// subscriber factory
type invoiceEventSubscriberFactory struct {
	context.Context
	subscribeType SubscriberType
	log           logger.Logger
	cfg           *config.AppConfig
}

func NewInvoiceEventSubscriberFactory(ctx context.Context, subscribeType SubscriberType, log logger.Logger, cfg *config.AppConfig) *invoiceEventSubscriberFactory {
	return &invoiceEventSubscriberFactory{ctx, subscribeType, log, cfg}
}

func (f *invoiceEventSubscriberFactory) MakeSubscriber(worker worker.Worker) ports.InvoiceEventSubscriber {
	switch f.subscribeType {
	case RABBITMQ:
		return rabbit_subscriber.NewRabbitSubscriber(f, worker, f.log, f.cfg)
	default:
		panic("unknown subscriber type: " + f.subscribeType)
	}
}
