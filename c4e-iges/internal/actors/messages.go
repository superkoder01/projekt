package actors

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type (
	InvoiceGenerateRequest struct {
		number string
	}
	InvoiceGenerateResponse struct {
		invoice        *billing.InvoiceProsument
		invoiceDetails interface{}
	}
	InvoiceGenerateError struct {
		Error error
	}

	RepurchaseInvoiceGenerateRequest struct {
		number string
	}
	RepurchaseInvoiceGenerateResponse struct {
		repurchaseInvoice        interface{}
		repurchaseInvoiceDetails interface{}
	}
	RepurchaseInvoiceGenerateError struct {
		Error error
	}

	SettlementMessage struct {
		invoice                  *billing.InvoiceProsument
		invoiceDetails           interface{}
		repurchaseInvoice        interface{}
		repurchaseInvoiceDetails interface{}
	}

	BillingMessage struct {
		InvoiceEvent *invoice.InvoiceEvent
	}
)
