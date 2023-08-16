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

type InvoiceStatus string

const (
	INVOICE = "INVOICE"

	ISSUED InvoiceStatus = "ISSUED"
	PAID   InvoiceStatus = "PAID"
)

type InvoiceEntity interface {
	SetContractID(int)
	SetInvoiceNumber(string)
	SetStatus(InvoiceStatus)
}

// TODO: error handling
func (c InvoiceStatus) Value() string {
	return string(c)
}

type Invoice struct {
	ID            int           `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ContractID    int           `gorm:"column:CONTRACT_ID;size:11;not null"`
	InvoiceNumber string        `gorm:"column:INVOICE_NUMER;size:55;default:null"`
	Status        InvoiceStatus `gorm:"column:STATUS" sql:"type:ENUM('ISSUED','PAID')"`
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (inv *Invoice) TableName() string {
	return INVOICE
}

func (inv *Invoice) SetContractID(i int) {
	inv.ContractID = i
}

func (inv *Invoice) SetInvoiceNumber(s string) {
	inv.InvoiceNumber = s
}

func (inv *Invoice) SetStatus(invs InvoiceStatus) {
	inv.Status = invs
}
