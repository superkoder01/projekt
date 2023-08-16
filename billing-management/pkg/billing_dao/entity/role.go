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
package entity

import "database/sql/driver"

type RoleName string

const (
	// table name
	ROLE = "ROLE"

	SUPER_ADMIN         RoleName = "SUPER_ADMIN"
	ADMINISTRATOR_FULL  RoleName = "ADMINISTRATOR_FULL"
	ADMINISTRATOR_BASIC RoleName = "ADMINISTRATOR_BASIC"
	TRADER              RoleName = "TRADER"
	SUPER_AGENT         RoleName = "SUPER_AGENT"
	AGENT               RoleName = "AGENT"
	PROSUMER            RoleName = "PROSUMER"
	ACCOUNTER           RoleName = "ACCOUNTER"
)

func (r *RoleName) Scan(value interface{}) error {
	*r = RoleName(value.([]byte))
	return nil
}

func (r RoleName) Value() (driver.Value, error) {
	return string(r), nil
}

type Role struct {
	ID   int      `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name RoleName `gorm:"column:NAME" sql:"type:ENUM('SUPER_ADMIN', 'ADMINISTRATOR_FULL','ADMINISTRATOR_BASIC','TRADER','SUPER_AGENT','AGENT','PROSUMER', 'ACCOUNTER')"`
}

func (r *Role) TableName() string {
	return ROLE
}
