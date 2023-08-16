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
package validators

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/coreutils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type contractValidator struct {
	log logger.Logger
	cfg *config.AppConfig
}

func NewContractValidator(log logger.Logger, cfg *config.AppConfig) ports.ContractValidator {
	return &contractValidator{log: log, cfg: cfg}
}

func (c *contractValidator) ValidateContract(ctx context.Context, invoiceEvent *invoice.InvoiceEvent, contract *billing.Contract) error {
	if contract == nil {
		return fmt.Errorf("contract is nil")
	}

	if invoiceEvent == nil {
		return fmt.Errorf("invoice event is nil")
	}

	c.log.Infof("validating contract: %s, customerId: %s", contract.Payload.ContractDetails.Number, contract.Payload.CustomerDetails.CustomerId)

	sapCodes := make([]string, 0, len(invoiceEvent.ServiceAccessPoints))
	for k, _ := range invoiceEvent.ServiceAccessPoints {
		sapCodes = append(sapCodes, k)
	}

	if len(sapCodes) == 0 {
		return fmt.Errorf("invoice event does not contain any service access point data")
	}

	var foundSap int64
	for _, sapCode := range sapCodes {
		sap := coreutils.FindSap(sapCode, contract.Payload.ServiceAccessPoints)
		if sap == nil {
			c.log.Warnf("seems like service access point %s does not exist in client contract %s", sapCode, contract.Payload.ContractDetails.Number)
		} else {
			foundSap += 1
		}
	}

	if foundSap < 1 {
		return fmt.Errorf("neither invoice event service access point found in client contract")
	}

	return nil
}
