package invoicerepo

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/invoicerepo/mongo"
)

type InvoiceRepoType string

const (
	MONGO InvoiceRepoType = "mongo"
)

type invoiceRepoFactory struct {
	context.Context
	invoiceRepoType InvoiceRepoType
	log             logger.Logger
	cfg             *config.AppConfig
}

func NewInvoiceRepoFactory(ctx context.Context, invoiceRepoType InvoiceRepoType, log logger.Logger, cfg *config.AppConfig) *invoiceRepoFactory {
	return &invoiceRepoFactory{ctx, invoiceRepoType, log, cfg}
}

func (f *invoiceRepoFactory) MakeRepo() ports.InvoiceRepo {
	switch f.invoiceRepoType {
	case MONGO:
		return mongo.NewInvoiceRepo(f.log, f.cfg)
	default:
		panic("unknown repo: " + f.invoiceRepoType)
	}
}
