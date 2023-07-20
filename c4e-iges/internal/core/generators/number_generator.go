package generators

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"time"
)

type layout string

const (
	yearMonth layout = "2006/01"
)

type base struct {
	log  logger.Logger
	cfg  *config.AppConfig
	repo ports.InvoiceCountRepo
}

// invoiceNumberGenerator number generator for invoices
type invoiceNumberGenerator struct {
	base
}

func NewInvoiceNumberGenerator(repo ports.InvoiceCountRepo, log logger.Logger, cfg *config.AppConfig) ports.InvoiceNumberGenerator {
	return &invoiceNumberGenerator{
		base: base{
			repo: repo,
			log:  log,
			cfg:  cfg,
		},
	}
}

func (g *invoiceNumberGenerator) GetNumber(ctx context.Context, customerId string, event invoice.InvoiceEvent) (string, error) {
	g.log.Infof("generating invoice number for contract: %s, customerId: %s", event.Contract, customerId)

	t, err := time.Parse("02/01/2006", event.StartDate)
	if err != nil {
		g.log.Warnf("unable to parse billing start date for contract %s, reason: %v", event.Contract, err)
		t = time.Now()
	}

	from := event.StartDate
	to := event.EndDate
	count, err := g.repo.CountInvoices(ctx, customerId, from, to)
	if err != nil {
		g.log.Warnf("error when counting customerId %s invoices: %v", customerId, err)
		g.log.Warnf("assuming no invoices has been generated in current month for customerId %s", customerId)
		count = 0
	}

	result := fmt.Sprintf("%s/%s/%s/%d", t.Format(string(yearMonth)), customerId, domain.InvoiceTypeSell.Get(), count+1)
	g.log.Infof("generated invoice number: %s", result)

	return result, nil
}

// repurchaseInvoiceNumberGenerator number generator for repurchase invoices
type repurchaseInvoiceNumberGenerator struct {
	base
}

func NewRepurchaseInvoiceNumberGenerator(repo ports.InvoiceCountRepo, log logger.Logger, cfg *config.AppConfig) ports.InvoiceNumberGenerator {
	return &repurchaseInvoiceNumberGenerator{
		base: base{
			repo: repo,
			log:  log,
			cfg:  cfg,
		},
	}
}

func (g *repurchaseInvoiceNumberGenerator) GetNumber(ctx context.Context, customerId string, event invoice.InvoiceEvent) (string, error) {
	g.log.Infof("generating repurchase invoice number for contract: %s, customerId: %s", event.Contract, customerId)

	t, err := time.Parse("02/01/2006", event.StartDate)
	if err != nil {
		g.log.Warnf("unable to parse billing start date for contract %s, reason: %v", event.Contract, err)
		t = time.Now()
	}

	from := event.StartDate
	to := event.EndDate
	count, err := g.repo.CountRepurchaseInvoices(ctx, customerId, from, to)
	if err != nil {
		g.log.Warnf("error when counting customerId %s repurchase invoices: %v", customerId, err)
		g.log.Warnf("assuming no repurchase invoices has been generated in current month for customerId %s", customerId)
		count = 0
	}

	result := fmt.Sprintf("%s/%s/%s/%d", t.Format(string(yearMonth)), customerId, domain.InvoiceTypeRepurchase.Get(), count+1)
	g.log.Infof("generated repurchase invoice number: %s", result)

	return result, nil
}
