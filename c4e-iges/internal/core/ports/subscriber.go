package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/worker"
)

type (
	InvoiceEventSubscriber interface {
		Subscribe(ctx context.Context) chan error
		Close() error
	}

	InvoiceEventSubscriberFactory interface {
		MakeSubscriber(worker.Worker) InvoiceEventSubscriber
	}
)
