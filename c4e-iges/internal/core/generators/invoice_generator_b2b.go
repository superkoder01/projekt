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
package generators

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type invoiceGeneratorB2B struct {
	log           logger.Logger
	cfg           *config.AppConfig
	contract      *billing.Contract
	event         *invoice.InvoiceEvent
	invoiceNumber string
}

func NewInvoiceGeneratorB2B(contract *billing.Contract, event *invoice.InvoiceEvent, log logger.Logger, cfg *config.AppConfig) ports.InvoiceGenerator {
	return &invoiceGeneratorB2B{
		log:      log,
		cfg:      cfg,
		contract: contract,
		event:    event,
	}
}

func (service *invoiceGeneratorB2B) GenerateInvoice(ctx context.Context, number string) (*billing.InvoiceProsument, interface{}, error) {
	service.invoiceNumber = number
	service.log.Infof("start generating invoice, contract: %s, customerId: %s, invoiceNumber: %s",
		service.contract.Payload.ContractDetails.Number,
		service.contract.Payload.CustomerDetails.CustomerId,
		service.invoiceNumber)

	//TODO implement me
	return &billing.InvoiceProsument{}, "invoice details b2b", nil
}
