package validation

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/contract"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/enum"
	"strings"
)

func ValidateIfCanUpdateContract(oldContract *contract.Contract) bool {
	return !strings.EqualFold(oldContract.Payload.ContractDetails.State, enum.CS_ACCEPTED.Name())
}
