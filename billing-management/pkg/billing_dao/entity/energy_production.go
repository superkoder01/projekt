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

import "time"

const (
	ENERGY_PRODUCTION = "ENERGY_PRODUCTION"
)

type EnergyProductionEntity interface {
	SetContractID(i int)
	SetServiceAccessPointID(i int)
	SetStartDate(t time.Time)
	SetEndDate(t time.Time)
	SetEnergyAmount(i int)
	SetEnergyAmountUnits(epau Units)
	SetNetPrice(f float64)
	SetNetValue(f float64)
}

type EnergyProduction struct {
	ID                   int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	StartDate            time.Time `gorm:"column:START_DATE"`
	EndDate              time.Time `gorm:"column:END_DATE"`
	EnergyAmount         int       `gorm:"column:ENERGY_AMOUNT;size:11"`
	EnergyAmountUnits    Units     `gorm:"column:ENERGY_AMOUNT_UNITS;size:45"`
	NetPrice             float64   `gorm:"column:NET_PRICE"`
	NetValue             float64   `gorm:"column:NET_VALUE"`
	ContractID           int       `gorm:"column:CONTRACT_ID;size:11;not null"`
	ServiceAccessPointID int       `gorm:"column:SERVICE_ACCESS_POINT_ID;size:11;not null"`
}

func NewEnergyProduction() *EnergyProduction {
	return &EnergyProduction{}
}

func (ep *EnergyProduction) TableName() string {
	return ENERGY_PRODUCTION
}

func (ep *EnergyProduction) SetContractID(i int) {
	ep.ContractID = i
}

func (ep *EnergyProduction) SetServiceAccessPointID(i int) {
	ep.ServiceAccessPointID = i
}

func (ep *EnergyProduction) SetStartDate(t time.Time) {
	ep.StartDate = t
}

func (ep *EnergyProduction) SetEndDate(t time.Time) {
	ep.EndDate = t
}

func (ep *EnergyProduction) SetEnergyAmount(i int) {
	ep.EnergyAmount = i
}

func (ep *EnergyProduction) SetEnergyAmountUnits(epau Units) {
	ep.EnergyAmountUnits = epau
}

func (ep *EnergyProduction) SetNetPrice(f float64) {
	ep.NetPrice = f
}

func (ep *EnergyProduction) SetNetValue(f float64) {
	ep.NetValue = f
}
