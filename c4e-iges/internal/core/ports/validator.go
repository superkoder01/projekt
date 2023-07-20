package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type (
	ContractValidator interface {
		ValidateContract(ctx context.Context, invoiceEvent *invoice.InvoiceEvent, contract *billing.Contract) error
	}

	InvoiceEventValidator interface {
		ValidateInvoiceEvent(ctx context.Context, invoiceEvent *invoice.InvoiceEvent) error
	}
)
