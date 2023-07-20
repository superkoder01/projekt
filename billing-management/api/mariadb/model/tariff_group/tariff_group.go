package tariff_group

import (
	"fmt"
	"time"
)

type TariffGroup struct {
	ID                            int       `json:"id,omitempty"`
	DistributionNetworkOperatorID int       `json:"distributionNetworkOperatorId"`
	TariffGroupLabelName          string    `json:"tariffGroupLabelName"`
	Name                          string    `json:"name"`
	StartDate                     time.Time `json:"startDate"`
	EndDate                       time.Time `json:"endDate"`
}

func (tg *TariffGroup) String() string {
	return fmt.Sprintf("%s", *tg)
}

func (tg *TariffGroup) SetTariffGroupLabelName(i string) {
	tg.TariffGroupLabelName = i
}

func (tg *TariffGroup) SetDistributionNetworkOperatorID(i int) {
	tg.DistributionNetworkOperatorID = i
}

func (tg *TariffGroup) SetName(s string) {
	tg.Name = s
}

func (tg *TariffGroup) SetStartDate(t time.Time) {
	tg.StartDate = t
}

func (tg *TariffGroup) SetEndDate(t time.Time) {
	tg.EndDate = t
}
