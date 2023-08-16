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
	SERVICE_ACCESS_POINT = "SERVICE_ACCESS_POINT"
)

type ServiceAccessPointEntity interface {
	SetAccountID(int)
	SetCity(string)
	SetMeterNumber(string)
	SetAddress(string)
	SetSapCode(string)
	SetProviderID(int)
	SetName(string)
}

type ServiceAccessPoint struct {
	ID          int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	AccountID   int    `gorm:"column:ACCOUNT_ID;size:11;not null"`
	City        string `gorm:"column:CITY;size:45;default:null"`
	Address     string `gorm:"column:ADDRESS;size:45;default:null"`
	MeterNumber string `gorm:"column:METER_NUMBER;size:45;default:null"`
	SapCode     string `gorm:"column:SAP_CODE;size:45;default:null"`
	ProviderID  int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	Name        string `gorm:"column:NAME;size:45;default:null"`
}

func NewServiceAccessPoint() *ServiceAccessPoint {
	return &ServiceAccessPoint{}
}

func (sap *ServiceAccessPoint) TableName() string {
	return SERVICE_ACCESS_POINT
}

func (sap *ServiceAccessPoint) SetAccountID(i int) {
	sap.AccountID = i
}

func (sap *ServiceAccessPoint) SetCity(s string) {
	sap.City = s
}

func (sap *ServiceAccessPoint) SetMeterNumber(s string) {
	sap.MeterNumber = s
}

func (sap *ServiceAccessPoint) SetAddress(s string) {
	sap.Address = s
}

func (sap *ServiceAccessPoint) SetSapCode(s string) {
	sap.SapCode = s
}

func (sap *ServiceAccessPoint) SetProviderID(i int) {
	sap.ProviderID = i
}

func (sap *ServiceAccessPoint) SetName(s string) {
	sap.Name = s
}
