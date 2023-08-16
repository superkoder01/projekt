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

import (
	"time"
)

type ServiceOfferGroupStatus string

const (
	SERVICE_OFFER_GROUP = "SERVICE_OFFER_GROUP"

	SOGS_DRAFT       ServiceOfferGroupStatus = "DRAFT"
	SOGS_FINAL       ServiceOfferGroupStatus = "FINAL"
	SOGS_SENT        ServiceOfferGroupStatus = "SENT"
	SOGS_DELIVERED   ServiceOfferGroupStatus = "DELIVERED"
	SOGS_NEGOTIATION ServiceOfferGroupStatus = "NEGOTIATION"
)

// TODO: error handling
func (c ServiceOfferGroupStatus) Value() string {
	return string(c)
}

type ServiceOfferGroupEntity interface {
	SetProviderID(int)
	SetName(string)
	SetIsActive(bool)
	SetStatus(ServiceOfferGroupStatus)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}

type ServiceOfferGroup struct {
	ID         int                     `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID int                     `gorm:"column:PROVIDER_ID;size:11;not null"`
	Name       string                  `gorm:"column:NAME;size:45;default:null"`
	IsActive   bool                    `gorm:"column:IS_ACTIVE;default:0"`
	Status     ServiceOfferGroupStatus `gorm:"column:STATUS" sql:"type:ENUM('DRAFT', 'FINAL', 'SENT', 'DELIVERED', 'NEGOTIATION')"`
	StartDate  time.Time               `gorm:"column:START_DATE"`
	EndDate    time.Time               `gorm:"column:END_DATE"`
}

func NewServiceOfferGroup() *ServiceOfferGroup {
	return &ServiceOfferGroup{}
}

func (sog *ServiceOfferGroup) TableName() string {
	return SERVICE_OFFER_GROUP
}

func (sog *ServiceOfferGroup) SetProviderID(i int) {
	sog.ProviderID = i
}

func (sog *ServiceOfferGroup) SetName(s string) {
	sog.Name = s
}

func (sog *ServiceOfferGroup) SetIsActive(b bool) {
	sog.IsActive = b
}

func (sog *ServiceOfferGroup) SetStatus(sogs ServiceOfferGroupStatus) {
	sog.Status = sogs
}

func (sog *ServiceOfferGroup) SetStartDate(t time.Time) {
	sog.StartDate = t
}

func (sog *ServiceOfferGroup) SetEndDate(t time.Time) {
	sog.EndDate = t
}
