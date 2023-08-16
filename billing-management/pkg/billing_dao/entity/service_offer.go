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
	SERVICE_OFFER = "SERVICE_OFFER"
)

type ServiceOfferEntity interface {
	SetServiceID(int)
	SetProviderID(int)
	SetRatingPlanID(int)
	SetServiceOfferGroupID(int)
	SetName(string)
}

type ServiceOffer struct {
	ID                  int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ServiceID           int    `gorm:"column:SERVICE_ID;size:11;not null"`
	ProviderID          int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	RatingPlanID        int    `gorm:"column:RATING_PLAN_ID;size:11;not null"`
	ServiceOfferGroupID int    `gorm:"column:SERVICE_OFFER_GROUP_ID;size:11;not null"`
	Name                string `gorm:"column:NAME;size:45;default:null"`
}

func NewServiceOffer() *ServiceOffer {
	return &ServiceOffer{}
}

func (so *ServiceOffer) TableName() string {
	return SERVICE_OFFER
}

func (so *ServiceOffer) SetServiceID(i int) {
	so.ServiceID = i
}

func (so *ServiceOffer) SetProviderID(i int) {
	so.ProviderID = i
}

func (so *ServiceOffer) SetRatingPlanID(i int) {
	so.RatingPlanID = i
}

func (so *ServiceOffer) SetServiceOfferGroupID(i int) {
	so.ServiceOfferGroupID = i
}

func (so *ServiceOffer) SetName(s string) {
	so.Name = s
}
