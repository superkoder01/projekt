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

type Units string

const (
	ZONE = "ZONE"

	kWh Units = "kWh"
	mWh Units = "mWh"
)

// TODO: error handling
func (c Units) Value() string {
	return string(c)
}

type ZoneEntity interface {
	SetProviderID(int)
	SetRatingPlanID(int)
	SetDescription(string)
	SetHours(string)
	SetNetPrice(float64)
	SetUnits(Units)
}

type Zone struct {
	ID           int     `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	RatingPlanID int     `gorm:"column:RATING_PLAN_ID;size:11;not null"`
	Description  string  `gorm:"column:DESCRIPTION;size:100;default:null"`
	Hours        string  `gorm:"column:HOURS;size:45;default:null"`
	NetPrice     float64 `gorm:"column:NET_PRICE;"`
	Units        Units   `gorm:"column:UNITS" sql:"type:ENUM('kWh','mWh')"`
}

func NewZone() *Zone {
	return &Zone{}
}

func (z *Zone) TableName() string {
	return ZONE
}

func (z *Zone) SetRatingPlanID(i int) {
	z.RatingPlanID = i
}

func (z *Zone) SetDescription(s string) {
	z.Description = s
}

func (z *Zone) SetHours(s string) {
	z.Hours = s
}

func (z *Zone) SetNetPrice(s float64) {
	z.NetPrice = s
}

func (z *Zone) SetUnits(s Units) {
	z.Units = s
}
