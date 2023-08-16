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

type CustomerTypeName string

const (
	// table name
	CUSTOMER_TYPE = "CUSTOMER_TYPE"

	CT_CONSUMER CustomerTypeName = "CONSUMER"
	CT_PROSUMER CustomerTypeName = "PROSUMER"
	CT_PRODUCER CustomerTypeName = "PRODUCER"
)

func (c *CustomerTypeName) Scan(value interface{}) error {
	*c = CustomerTypeName(value.([]byte))
	return nil
}

func (c CustomerTypeName) Value() (driver.Value, error) {
	return string(c), nil
}

type CustomerType struct {
	ID   int              `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name CustomerTypeName `gorm:"column:NAME" sql:"type:ENUM('CONSUMER', 'PROSUMER','PRODUCER')"`
}

func (c *CustomerType) TableName() string {
	return CUSTOMER_TYPE
}
