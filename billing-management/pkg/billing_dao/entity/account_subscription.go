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
	ACCOUNT_SUBSCRIPTION = "ACCOUNT_SUBSCRIPTION"
)

type AccountSubscriptionEntity interface {
	SetServiceOfferGroupID(int)
	SetAccountID(int)
	SetProviderID(int)
	SetStatus(bool)
	SetPriority(int)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}

type AccountSubscription struct {
	Name                string    `gorm:"column:NAME;size:45;primaryKey;not null"`
	ServiceOfferGroupID int       `gorm:"column:SERVICE_OFFER_GROUP_ID;size:11;not null"`
	AccountID           int       `gorm:"column:ACCOUNT_ID;size:11;not null"`
	ProviderID          int       `gorm:"column:PROVIDER_ID;size:11;not null"`
	Status              bool      `gorm:"column:STATUS;default:1"`
	Priority            int       `gorm:"column:PRIORITY;size:11;not null"`
	StartDate           time.Time `gorm:"column:START_DATE"`
	EndDate             time.Time `gorm:"column:END_DATE"`
}

func NewAccountSubscription() *AccountSubscription {
	return &AccountSubscription{}
}

func (as *AccountSubscription) TableName() string {
	return ACCOUNT_SUBSCRIPTION
}

func (as *AccountSubscription) SetServiceOfferGroupID(i int) {
	as.ServiceOfferGroupID = i
}

func (as *AccountSubscription) SetAccountID(i int) {
	as.AccountID = i
}

func (as *AccountSubscription) SetProviderID(i int) {
	as.ProviderID = i
}

func (as *AccountSubscription) SetStatus(b bool) {
	as.Status = b
}

func (as *AccountSubscription) SetPriority(i int) {
	as.Priority = i
}

func (as *AccountSubscription) SetStartDate(t time.Time) {
	as.StartDate = t
}

func (as *AccountSubscription) SetEndDate(t time.Time) {
	as.EndDate = t
}
