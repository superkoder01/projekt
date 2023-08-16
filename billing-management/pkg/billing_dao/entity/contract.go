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

type ContractStatus string

const (
	CONTRACT = "CONTRACT"

	CS_DRAFT     ContractStatus = "DRAFT"
	CS_FINAL     ContractStatus = "FINAL"
	CS_SENT      ContractStatus = "SENT"
	CS_DELIVERED ContractStatus = "DELIVERED"
	CS_ACCEPTED  ContractStatus = "ACCEPTED"
)

// TODO: error handling
func (c ContractStatus) Value() string {
	return string(c)
}

type ContractEntity interface {
	SetDistributionNetworkOperatorID(int)
}

type Contract struct {
	ID     int            `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Number string         `gorm:"column:NUMBER;size:45;default:null"`
	Status ContractStatus `gorm:"column:STATUS" sql:"type:ENUM('DRAFT', 'FINAL', 'SENT', 'DELIVERED', 'ACCEPTED')"`
}

func NewContract() *Contract {
	return &Contract{}
}

func (c *Contract) TableName() string {
	return CONTRACT
}

func (c *Contract) SetStatus(cs ContractStatus) {
	c.Status = cs
}

func (c *Contract) SetNumber(s string) {
	c.Number = s
}
