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
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type vatGroup string

const (
	Gxx vatGroup = "Gxx"
)

func (vg *vatGroup) Scan(value interface{}) error {
	*vg = vatGroup(value.([]byte))
	return nil
}

func (vg vatGroup) Value() (driver.Value, error) {
	return string(vg), nil
}

type VatRate struct {
	Value int `gorm:"column:RATE"`
}

func (VatRate) TableName() string {
	return "VAT_RATES"
}

func (v VatRate) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(v, "", "\t")
	return fmt.Sprintf("%s", data)
}
