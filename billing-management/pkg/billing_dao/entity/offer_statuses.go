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
	OFFER_STATUSES = "OFFER_STATUSES"
)

type OfferStatuses struct {
	Name string `gorm:"column:NAME;primaryKey;size:45"`
}

func NewOfferStatuses() *OfferStatuses {
	return &OfferStatuses{}
}

func (tt *OfferStatuses) TableName() string {
	return OFFER_STATUSES
}
