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

type ServiceAccessPoint struct {
	ID          uint   `gorm:"column:ID"`
	SapCode     string `gorm:"column:SAP_CODE"`
	MeterNumber string `gorm:"column:METER_NUMBER"`
	City        string `gorm:"column:CITY"`
	Address     string `gorm:"column:ADDRESS"`
	AccountId   int    `gorm:"column:ACCOUNT_ID"`
}

func (ServiceAccessPoint) TableName() string {
	return "SERVICE_ACCESS_POINT"
}

func (p ServiceAccessPoint) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%s", data)
}
