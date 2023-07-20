package validators

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
)

type invoiceEventValidator struct {
	log logger.Logger
	cfg *config.AppConfig
}

func NewInvoiceEventValidator(log logger.Logger, cfg *config.AppConfig) ports.InvoiceEventValidator {
	return &invoiceEventValidator{log: log, cfg: cfg}
}

func (i *invoiceEventValidator) ValidateInvoiceEvent(ctx context.Context, invoiceEvent *invoice.InvoiceEvent) error {
	if invoiceEvent == nil {
		return fmt.Errorf("invoice event is nil")
	}

	i.log.Infof("validating invoice event received for contract: %s", invoiceEvent.Contract)

	if invoiceEvent.Contract == "" {
		return fmt.Errorf("invoice event contains empty contract number")
	}

	if invoiceEvent.StartDate == "" {
		return fmt.Errorf("invoice event contains empty billing start date")
	}

	if invoiceEvent.EndDate == "" {
		return fmt.Errorf("invoice event contains empty billing end date")
	}

	if invoiceEvent.ServiceAccessPoints == nil || len(invoiceEvent.ServiceAccessPoints) == 0 {
		return fmt.Errorf("invoice event contains no service access point data")
	}

	return nil
}
