package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
)

type (
	InvoiceNumberGenerator interface {
		GetNumber(ctx context.Context, customerId string, event invoice.InvoiceEvent) (string, error)
	}
)
