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

const (
	ACCOUNT_BALANCE = "ACCOUNT_BALANCE"
)

type AccountBalanceEntity interface {
	SetAccountID(int)
	SetProviderID(int)
	SetBalanceTypeName(string)
	SetName(string)
	SetStatus(bool)
	SetUnits(int)
}

type AccountBalance struct {
	ID              int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID      int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	AccountID       int    `gorm:"column:ACCOUNT_ID;size:11;not null"`
	BalanceTypeName string `gorm:"column:BALANCE_TYPE_NAME;size:45;not null"`
	Name            string `gorm:"column:NAME;size:45;default:null"`
	Status          bool   `gorm:"column:STATUS;default:1"`
	Units           int    `gorm:"column:UNITS;size:11;not null"`
}

func NewAccountBalance() *AccountBalance {
	return &AccountBalance{}
}

func (ab *AccountBalance) TableName() string {
	return ACCOUNT_BALANCE
}

func (ab *AccountBalance) SetProviderID(i int) {
	ab.ProviderID = i
}

func (ab *AccountBalance) SetAccountID(i int) {
	ab.AccountID = i
}

func (ab *AccountBalance) SetBalanceTypeName(s string) {
	ab.BalanceTypeName = s
}

func (ab *AccountBalance) SetName(s string) {
	ab.Name = s
}

func (ab *AccountBalance) SetStatus(b bool) {
	ab.Status = b
}

func (ab *AccountBalance) SetUnits(i int) {
	ab.Units = i
}
