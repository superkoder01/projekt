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

type RatingPlanTypeName string

const (
	RATING_PLAN_TYPE = "RATING_PLAN_TYPE"

	RPT_SALE       RatingPlanTypeName = "SALE"
	RPT_REPURCHASE RatingPlanTypeName = "REPURCHASE"
	RPT_RDN        RatingPlanTypeName = "RDN"
)

// TODO: error handling
func (c RatingPlanTypeName) Value() string {
	return string(c)
}

type RatingPlanType struct {
	ID   int                `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name RatingPlanTypeName `gorm:"column:NAME" sql:"type:ENUM('SALE', 'REPURCHASE', 'RDN')"`
}

func NewRatingPlanType() *RatingPlanType {
	return &RatingPlanType{}
}

func (rpt *RatingPlanType) TableName() string {
	return RATING_PLAN_TYPE
}
