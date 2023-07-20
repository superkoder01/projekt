package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"time"
)

type (
	ContractRepo interface {
		GetContractByContractNumber(ctx context.Context, contractNumber string) (*billing.Contract, error)
	}

	ContractReposFactory interface {
		MakeRepo() ContractRepo
	}

	InvoiceStoreRepo interface {
		StoreOne(ctx context.Context, document interface{}) error
		StoreMany(ctx context.Context, documents ...interface{}) error
		StoreManyWithinTransaction(ctx context.Context, documents ...interface{}) error
	}

	InvoiceCountRepo interface {
		CountInvoices(ctx context.Context, customerId string, from string, to string) (int64, error)
		CountRepurchaseInvoices(ctx context.Context, customerId string, from string, to string) (int64, error)
	}

	InvoiceFetchRepo interface {
		GetInvoiceSummaryByIssueDate(ctx context.Context, issueStartDate time.Time, issueEndDate time.Time) (*billing.InvoiceSummary, error)
	}

	InvoiceRepo interface {
		InvoiceStoreRepo
		InvoiceCountRepo
		InvoiceFetchRepo
	}

	InvoiceRepoFactory interface {
		MakeRepo() InvoiceRepo
	}
)
