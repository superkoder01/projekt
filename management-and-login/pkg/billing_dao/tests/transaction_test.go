package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
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
