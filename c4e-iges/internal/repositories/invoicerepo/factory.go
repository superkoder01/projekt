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
