package generators

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type repurchaseInvoiceGeneratorB2B struct {
	log           logger.Logger
	cfg           *config.AppConfig
	contract      *billing.Contract
	event         *invoice.InvoiceEvent
	invoiceNumber string
}

func NewRepurchaseInvoiceGeneratorB2B(contract *billing.Contract, event *invoice.InvoiceEvent, log logger.Logger, cfg *config.AppConfig) ports.RepurchaseInvoiceGenerator {
	return &repurchaseInvoiceGeneratorB2B{
		log:      log,
		cfg:      cfg,
		contract: contract,
		event:    event,
	}
}

func (service *repurchaseInvoiceGeneratorB2B) GenerateRepurchaseInvoice(ctx context.Context, number string) (*billing.InvoiceProsumentRepurchase, *billing.InvoiceProsumentRepurchaseDetails, error) {
	service.invoiceNumber = number
	service.log.Infof("start generating invoice, contract: %s, customerId: %s, invoiceNumber: %s",
		service.contract.Payload.ContractDetails.Number,
		service.contract.Payload.CustomerDetails.CustomerId,
		service.invoiceNumber)

	//TODO implement me
	return nil, nil, nil
}
