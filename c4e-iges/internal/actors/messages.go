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
