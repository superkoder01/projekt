package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type (
	InvoiceGenerator interface {
		GenerateInvoice(ctx context.Context, number string) (*billing.InvoiceProsument, interface{}, error)
	}

	RepurchaseInvoiceGenerator interface {
		GenerateRepurchaseInvoice(ctx context.Context, number string) (*billing.InvoiceProsumentRepurchase, *billing.InvoiceProsumentRepurchaseDetails, error)
	}
)
