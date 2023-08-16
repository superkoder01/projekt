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
	USER = "USER"
)

type UserEntity interface {
	SetLogin(string)
	SetPassword(string)
	SetEmail(string)
	SetProviderID(int)
	SetWorkerID(int)
	SetCustomerAccountID(int)
	SetRoleID(int)
	SetActive(bool)
	SetActivationCode(string)
	SetMustChangePassword(bool)
	SetCustomerAccount(*CustomerAccount)
	SetBlockchainAccAddress(string)
	SetBlockchainPubKey(string)
}

type User struct {
	ID                   int              `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Login                string           `gorm:"column:LOGIN;size:255;not null;unique"`
	Password             string           `gorm:"column:PASSWORD;size:255;not null"`
	Email                string           `gorm:"column:EMAIL;size:45;unique"`
	ProviderID           int              `gorm:"column:PROVIDER_ID;size:11;default:null"`
	CustomerAccountID    int              `gorm:"column:CUSTOMER_ACCOUNT_ID;size:11;default:null"`
	RoleID               int              `gorm:"column:ROLE_ID;size:11;default:null"`
	Role                 *Role            `gorm:"foreignKey:ROLE_ID;references:ID"`
	Active               bool             `gorm:"column:ACTIVE;default:0"`
	ActivationCode       string           `gorm:"column:ACTIVATION_CODE;default:null"`
	MustChangePassword   bool             `gorm:"column:MUST_CHANGE_PASSWORD;default:0"`
	CustomerAccount      *CustomerAccount `gorm:"foreignKey:CUSTOMER_ACCOUNT_ID;references:ID"`
	WorkerID             int              `gorm:"column:WORKER_ID;size:11;default:null"`
	Worker               *Worker          `gorm:"foreignKey:WORKER_ID;references:ID"`
	BlockchainAccAddress string           `gorm:"column:BLOCKCHAIN_ACC_ADDRESS;default:null"`
	BlockchainPubKey     string           `gorm:"column:BLOCKCHAIN_PUB_KEY;default:null"`
	AddedDate            time.Time        `gorm:"column:ADDED_DATE;not null"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) TableName() string {
	return USER
}

func (u *User) SetLogin(s string) {
	u.Login = s
}

func (u *User) SetPassword(s string) {
	u.Password = s
}

func (u *User) SetEmail(s string) {
	u.Email = s
}

func (u *User) SetProviderID(i int) {
	u.ProviderID = i
}

func (u *User) SetCustomerAccountID(i int) {
	u.CustomerAccountID = i
}

func (u *User) SetCustomerAccount(c *CustomerAccount) {
	u.CustomerAccount = c
}

func (u *User) SetRoleID(i int) {
	u.RoleID = i
}

func (u *User) SetActive(b bool) {
	u.Active = b
}

func (u *User) SetActivationCode(s string) {
	u.ActivationCode = s
}

func (u *User) SetMustChangePassword(b bool) {
	u.MustChangePassword = b
}

func (u *User) SetWorkerID(i int) {
	u.WorkerID = i
}

func (u *User) SetBlockchainAccAddress(s string) {
	u.BlockchainAccAddress = s
}

func (u *User) SetBlockchainPubKey(s string) {
	u.BlockchainPubKey = s
}
