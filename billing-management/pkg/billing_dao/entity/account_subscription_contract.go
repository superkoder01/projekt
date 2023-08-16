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
	ACCOUNT_SUBSCRIPTION_CONTRACT = "ACCOUNT_SUBSCRIPTION_CONTRACT"
)

type AccountSubscriptionContractEntity interface {
	SetAccountSubscriptionServiceOfferGroupID(int)
	SetAccountSubscriptionAccountID(int)
	SetAccountSubscriptionProviderID(int)
}

type AccountSubscriptionContract struct {
	ContractID                             int `gorm:"column:CONTRACT_ID;size:11;not null"`
	AccountSubscriptionServiceOfferGroupID int `gorm:"column:ACCOUNT_SUBSCRIPTION_SERVICE_OFFER_GROUP_ID;size:11;not null"`
	AccountSubscriptionAccountID           int `gorm:"column:ACCOUNT_SUBSCRIPTION_ACCOUNT_ID;size:11;not null"`
	AccountSubscriptionProviderID          int `gorm:"column:ACCOUNT_SUBSCRIPTION_PROVIDER_ID;size:11;not null"`
}

func NewAccountSubscriptionContract() *AccountSubscriptionContract {
	return &AccountSubscriptionContract{}
}

func (as *AccountSubscriptionContract) TableName() string {
	return ACCOUNT_SUBSCRIPTION_CONTRACT
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionServiceOfferGroupID(i int) {
	as.AccountSubscriptionServiceOfferGroupID = i
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionAccountID(i int) {
	as.AccountSubscriptionAccountID = i
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionProviderID(i int) {
	as.AccountSubscriptionProviderID = i
}
