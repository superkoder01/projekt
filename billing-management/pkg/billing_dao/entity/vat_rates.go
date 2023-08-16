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

type VatRateGroup string

const (
	VAT_RATES = "VAT_RATES"

	Gxx       VatRateGroup = "Gxx"
	OWN_USAGE VatRateGroup = "OWN_USAGE"
	RESALE    VatRateGroup = "RESALE"
)

// TODO: error handling
func (c VatRateGroup) Value() string {
	return string(c)
}

type VatRatesEntity interface {
	SetDescription(string)
	SetValue(int)
	SetGroup(VatRateGroup)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}

type VatRates struct {
	ID          int          `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Description string       `gorm:"column:DESCRIPTION;size:45;default:null"`
	Value       int          `gorm:"column:VALUE;size:11;not null"`
	Group       VatRateGroup `gorm:"column:GROUP" sql:"type:ENUM('Gxx', 'OWN_USAGE', 'RESALE')"`
	StartDate   time.Time    `gorm:"column:START_DATE"`
	EndDate     time.Time    `gorm:"column:END_DATE"`
}

func NewVatRates() *VatRates {
	return &VatRates{}
}

func (vr *VatRates) TableName() string {
	return VAT_RATES
}

func (vr *VatRates) SetGroup(s VatRateGroup) {
	vr.Group = s
}

func (vr *VatRates) SetDescription(s string) {
	vr.Description = s
}

func (vr *VatRates) SetValue(i int) {
	vr.Value = i
}

func (vr *VatRates) SetStartDate(t time.Time) {
	vr.StartDate = t
}

func (vr *VatRates) SetEndDate(t time.Time) {
	vr.EndDate = t
}
