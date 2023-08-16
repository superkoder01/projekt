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

import "database/sql/driver"

type invoiceStatus string

const (
	ISSUED invoiceStatus = "ISSUED"
	PAID   invoiceStatus = "PAID"
)

func (ct *invoiceStatus) Scan(value interface{}) error {
	*ct = invoiceStatus(value.([]byte))
	return nil
}

func (ct invoiceStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Invoice struct {
	InvoiceNumber string        `gorm:"column:INVOICE_NUMBER"`
	ContractId    uint          `gorm:"column:CONTRACT_ID"`
	InvoiceStatus invoiceStatus `gorm:"column:STATUS"`
}

func (Invoice) TableName() string {
	return "INVOICE"
}
