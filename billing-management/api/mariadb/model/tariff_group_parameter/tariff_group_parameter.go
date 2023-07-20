package tariff_group_parameter

import "fmt"

type TariffGroupParameter struct {
	ID              int     `json:"id,omitempty"`
	TariffGroupID   int     `json:"tariffGroupId"`
	ParameterNameID int     `json:"parameterNameId"`
	Price           float64 `json:"price"`
}

func (tgp *TariffGroupParameter) String() string {
	return fmt.Sprintf("%s", *tgp)
}

func (tgp *TariffGroupParameter) SetTariffGroupID(i int) {
	tgp.TariffGroupID = i
}

func (tgp *TariffGroupParameter) SetParameterNameID(i int) {
	tgp.ParameterNameID = i
}

func (tgp *TariffGroupParameter) SetPrice(f float64) {
	tgp.Price = f
}
