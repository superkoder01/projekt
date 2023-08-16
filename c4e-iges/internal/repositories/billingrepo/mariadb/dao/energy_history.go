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
package mariadb_dao

import (
	"encoding/json"
	"fmt"
)

type EnergyHistory struct {
	Period                   int  `gorm:"column:PERIOD"`
	CurrentEnergyConsumption int  `gorm:"column:CURRENT_ENERGY_CONSUMPTION"`
	CurrentEnergyExcess      int  `gorm:"column:CURRENT_ENERGY_EXCESS"`
	PreviousEnergyExcess     int  `gorm:"column:PREVIOUS_ENERGY_EXCESS"`
	EnergyExcessUsage        int  `gorm:"column:ENERGY_EXCESS_USAGE"`
	EnergyExcessTransfer     int  `gorm:"column:ENERGY_EXCESS_TRANSFER"`
	ServiceAccessPointId     uint `gorm:"column:SERVICE_ACCESS_POINT_ID"`
}

func (EnergyHistory) TableName() string {
	return "ENERGY_EXCESS"
}

func (p EnergyHistory) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%s", data)
}
