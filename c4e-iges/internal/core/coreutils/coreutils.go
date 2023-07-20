package coreutils

import "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"

func FindSap(sapCode string, saps []billing.ServiceAccessPoints) *billing.ServiceAccessPoints {
	for _, sap := range saps {
		if sap.SapCode == sapCode {
			return &sap
		}
	}
	return nil
}
