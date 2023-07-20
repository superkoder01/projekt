package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type (
	InvoicePublisher interface {
		Publish(ctx context.Context, invoice *billing.InvoiceProsument) error
		Close() error
	}

	InvoicePublisherFactory interface {
		MakePublisher() InvoicePublisher
	}
)
