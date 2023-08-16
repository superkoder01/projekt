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
	SERVICE_ACCESS_POINT_CONTRACT = "SERVICE_ACCESS_POINT_CONTRACT"
)

type ServiceAccessPointContractEntity interface {
	SetContractID(int)
	SetEnergyPointID(int)
}

type ServiceAccessPointContract struct {
	ID                     int   `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ContractID             int   `gorm:"column:CONTRACT_ID;size:11;not null"`
	AnnualEnergyUsage      int   `gorm:"column:ANNUAL_ENERGY_USAGE;size:11"`
	AnnualEnergyUsageUnits Units `gorm:"column:ANNUAL_ENERGY_USAGE_UNITS"`
	ServiceAccessPointID   int   `gorm:"column:SERVICE_ACCESS_POINT_ID;size:11;not null"`
}

func NewServiceAccessPointContract() *ServiceAccessPointContract {
	return &ServiceAccessPointContract{}
}

func (sapc *ServiceAccessPointContract) TableName() string {
	return SERVICE_ACCESS_POINT_CONTRACT
}

func (sapc *ServiceAccessPointContract) SetContractID(i int) {
	sapc.ContractID = i
}

func (sapc *ServiceAccessPointContract) SetAnnualEnergyUsage(i int) {
	sapc.AnnualEnergyUsage = i
}

func (sapc *ServiceAccessPointContract) SetAnnualEnergyUsageUnits(i Units) {
	sapc.AnnualEnergyUsageUnits = i
}

func (sapc *ServiceAccessPointContract) SetServiceAccessPointID(i int) {
	sapc.ServiceAccessPointID = i
}
