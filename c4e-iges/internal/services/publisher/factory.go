package publisher

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	dummy_publisher "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/publisher/dummy"
	rabbit_publisher "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/publisher/rabbit"
)

type PublisherType string

const (
	RABBITMQ PublisherType = "rabbitmq"
	DUMMY                  = "dummy"
)

// publisher factory
type publisherFactory struct {
	context.Context
	publisherType PublisherType
	log           logger.Logger
	cfg           *config.AppConfig
}

func NewPublisherFactory(ctx context.Context, publisherType PublisherType, log logger.Logger, cfg *config.AppConfig) *publisherFactory {
	return &publisherFactory{ctx, publisherType, log, cfg}
}

func (f *publisherFactory) MakePublisher() ports.InvoicePublisher {
	switch f.publisherType {
	case RABBITMQ:
		return rabbit_publisher.NewPublisher(f.log, f.cfg)
	case DUMMY:
		return dummy_publisher.NewPublisher(f.log, f.cfg)
	default:
		panic("unknown publisher type: " + f.publisherType)
	}
}
