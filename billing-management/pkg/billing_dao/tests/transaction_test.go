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
package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	"testing"
)

func TestTransactionCommit(t *testing.T) {
	mariaDBFlushData(s)

	newSession, err := newBillingSession()
	assert.Nil(t, err)
	assert.NotNil(t, newSession)

	// Begin transaction
	tx := newSession.Begin()
	assert.NotNil(t, tx)

	err = tx.Create(&entity.User{
		Login:    "testLogin",
		Password: "testPassword",
		Email:    "testemail",
	}).Error()
	assert.Nil(t, err)

	err = tx.Create(&entity.User{
		Login:    "testLogin2",
		Password: "testPassword2",
		Email:    "testemail2",
	}).Error()
	assert.Nil(t, err)

	err = tx.Commit().Error()
	assert.Nil(t, err)

	var users []*entity.User
	err = newSession.Find(&users).Error()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestTransactionRollback(t *testing.T) {
	mariaDBFlushData(s)

	newSession, err := newBillingSession()
	assert.Nil(t, err)
	assert.NotNil(t, newSession)

	// Begin transaction
	tx := newSession.Begin()
	assert.NotNil(t, tx)

	err = tx.Create(&entity.User{
		Login:    "testLogin",
		Password: "testPassword",
		Email:    "testemail",
	}).Error()
	assert.Nil(t, err)

	err = tx.Create(&entity.User{
		Login:    "testLogin2",
		Password: "testPassword2",
		Email:    "testemail2",
	}).Error()
	assert.Nil(t, err)

	err = tx.Rollback().Error()
	assert.Nil(t, err)

	var users []*entity.User
	err = newSession.Find(&users).Error()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}
