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
	DISTRIBUTION_NETWORK_OPERATOR = "DISTRIBUTION_NETWORK_OPERATOR"
)

type DistributionNetworkOperatorEntity interface {
	SetName(string)
	SetNip(string)
	SetAddress(string)
	SetCity(string)
}

type DistributionNetworkOperator struct {
	ID      int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name    string `gorm:"column:NAME;size:45;default:null"`
	Nip     string `gorm:"column:NIP;size:45;default:null"`
	Address string `gorm:"column:ADDRESS;size:45;default:null"`
	City    string `gorm:"column:CITY;size:45;default:null"`
}

func NewDistributionNetworkOperator() *DistributionNetworkOperator {
	return &DistributionNetworkOperator{}
}

func (dno *DistributionNetworkOperator) TableName() string {
	return DISTRIBUTION_NETWORK_OPERATOR
}

func (dno *DistributionNetworkOperator) SetName(s string) {
	dno.Name = s
}

func (dno *DistributionNetworkOperator) SetNip(s string) {
	dno.Nip = s
}

func (dno *DistributionNetworkOperator) SetAddress(s string) {
	dno.Address = s
}

func (dno *DistributionNetworkOperator) SetCity(s string) {
	dno.City = s
}
