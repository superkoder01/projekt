/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
