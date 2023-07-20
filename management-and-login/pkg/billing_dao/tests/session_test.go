package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"testing"
)

func TestSession(t *testing.T) {
	mariaDBFlushData(s)

	testSession, err := newBillingSession()
	assert.Nil(t, err)
	assert.NotNil(t, testSession)

	// Test insert
	err = testSession.Create(&entity.User{
		Login:    "testLogin",
		Password: "testPassword",
	}).Error()
	assert.Nil(t, err)

	// Test select
	var ts entity.User
	err = testSession.First(&ts).Error()
	assert.Nil(t, err)
	assert.Equal(t, entity.User{
		ID:       1,
		Login:    "testLogin",
		Password: "testPassword",
	}, ts)
}
