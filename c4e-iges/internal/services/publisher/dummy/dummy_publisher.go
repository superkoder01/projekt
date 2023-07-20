package dummy_publisher

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type dummyPublisher struct {
	log logger.Logger
	cfg *config.AppConfig
}

func NewPublisher(log logger.Logger, cfg *config.AppConfig) ports.InvoicePublisher {
	return &dummyPublisher{log: log, cfg: cfg}
}

func (dp *dummyPublisher) Publish(ctx context.Context, invoice *billing.InvoiceProsument) error {
	dp.log.Debugf("invoice successfully published: %v", invoice)
	return nil
}

func (dp *dummyPublisher) Close() error {
	dp.log.Infof("dummy invoice publisher stopping")
	return nil
}
