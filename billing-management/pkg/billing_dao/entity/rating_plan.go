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
	RATING_PLAN = "RATING_PLAN"
)

type RatingPlanEntity interface {
	SetInternalID(string)
	SetName(string)
	SetNetCommercialFee(float64)
	SetProviderID(int)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
	SetRatingPlanTypeID(int)
	SetTariffGroupLabelName(string)
}

type RatingPlan struct {
	ID                   int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	InternalID           string    `gorm:"column:INTERNAL_ID;size:45"`
	Name                 string    `gorm:"column:NAME;size:45;default:null"`
	NetCommercialFee     float64   `gorm:"column:NET_COMMERCIAL_FEE;size:11;not null"`
	ProviderID           int       `gorm:"column:PROVIDER_ID;size:11;not null"`
	StartDate            time.Time `gorm:"column:START_DATE"`
	EndDate              time.Time `gorm:"column:END_DATE"`
	RatingPlanTypeID     int       `gorm:"column:RATING_PLAN_TYPE_ID;size:11;not null"`
	TariffGroupLabelName string    `gorm:"column:TARIFF_GROUP_LABEL_NAME;size:45;not null"`
}

func NewRatingPlan() *RatingPlan {
	return &RatingPlan{}
}

func (rp *RatingPlan) TableName() string {
	return RATING_PLAN
}

func (rp *RatingPlan) SetProviderID(i int) {
	rp.ProviderID = i
}

func (rp *RatingPlan) SetRatingPlanTypeID(i int) {
	rp.RatingPlanTypeID = i
}

func (rp *RatingPlan) SetTariffGroupLabelName(s string) {
	rp.TariffGroupLabelName = s
}

func (rp *RatingPlan) SetName(s string) {
	rp.Name = s
}

func (rp *RatingPlan) SetInternalID(s string) {
	rp.InternalID = s
}

func (rp *RatingPlan) SetNetCommercialFee(f float64) {
	rp.NetCommercialFee = f
}

func (rp *RatingPlan) SetStartDate(t time.Time) {
	rp.StartDate = t
}

func (rp *RatingPlan) SetEndDate(t time.Time) {
	rp.EndDate = t
}
