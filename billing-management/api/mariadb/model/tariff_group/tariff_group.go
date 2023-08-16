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
