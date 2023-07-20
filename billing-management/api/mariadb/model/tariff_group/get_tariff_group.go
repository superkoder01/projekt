package tariff_group

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/fee"
)

type GetTariffGroup struct {
	TariffGroup
	Fees []fee.Fee `json:"fees"`
}

func (f *GetTariffGroup) String() string {
	return fmt.Sprintf("%s", *f)
}

func (f *GetTariffGroup) SetFees(i []fee.Fee) {
	f.Fees = i
}
