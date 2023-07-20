package tests

import (
	"github.com/stretchr/testify/assert"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"testing"
)

func TestCreateCustomerAccount(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)
	customerAccount := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccountEntity, err := customerAccountDao.Create(customerAccount)
	customerAccountEntityE := customerAccountEntity.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity)
	assert.Equal(t, customerAccount, customerAccountEntityE)
}

func TestGetCustomerAccounts(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount1 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccount2 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "34",
	}

	customerAccountEntity1, err := customerAccountDao.Create(customerAccount1)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity1)

	customerAccountEntity2, err := customerAccountDao.Create(customerAccount2)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity2)

	ens, err := customerAccountDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	assert.Equal(t, customerAccount1, customerAccountEntity1)
	assert.Equal(t, customerAccount2, customerAccountEntity2)
}

func TestUpdateCustomerAccountByID(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount1 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccount2 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		Country:          "Poland",
	}

	customerAccountEntity, err := customerAccountDao.Create(customerAccount1)
	customerAccountEntityE := customerAccountEntity.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity)
	assert.Equal(t, customerAccount1, customerAccountEntityE)

	customerAccountEntity2, err := customerAccountDao.Create(customerAccount2)
	customerAccountEntityE2 := customerAccountEntity2.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity2)
	assert.Equal(t, customerAccount2, customerAccountEntityE2)

	customerAccount1.SetApartmentNumber("34")
	customerAccountEntity, err = customerAccountDao.UpdateByID(customerAccountEntityE.ID, customerAccount1)
	customerAccountEntityE = customerAccountEntity.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity)
	assert.Equal(t, customerAccount1, customerAccountEntityE)

	customerAccountEntity2, err = customerAccountDao.GetByID(customerAccount2.ID)
	assert.Nil(t, err)
	assert.Equal(t, "", customerAccountEntityE2.ApartmentNumber)
}

func TestDeleteCustomerAccountByID(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
		City:             "Przemysl",
	}

	customerAccount2 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "34",
		City:             "Krakow",
	}

	customerAccount3 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "35",
		City:             "Warszawa",
	}

	customerAccountEntity, err := customerAccountDao.Create(customerAccount)
	customerAccountEntityE := customerAccountEntity.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity)
	assert.Equal(t, customerAccount, customerAccountEntityE)

	customerAccountEntity2, err := customerAccountDao.Create(customerAccount2)
	customerAccountEntityE2 := customerAccountEntity2.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity2)
	assert.Equal(t, customerAccount2, customerAccountEntityE2)

	customerAccountEntity3, err := customerAccountDao.Create(customerAccount3)
	customerAccountEntityE3 := customerAccountEntity3.(*entity.CustomerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity3)
	assert.Equal(t, customerAccount3, customerAccountEntityE3)

	err = customerAccountDao.DeleteByID(customerAccountEntityE2.ID)
	assert.Nil(t, err)

	ens, err := customerAccountDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestGetCustomerAccountsByFilter(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount1 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccount2 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "34",
	}

	customerAccount3 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "CONSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccountEntity1, err := customerAccountDao.Create(customerAccount1)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity1)

	customerAccountEntity2, err := customerAccountDao.Create(customerAccount2)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity2)

	customerAccountEntity3, err := customerAccountDao.Create(customerAccount3)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity3)

	ens, err := customerAccountDao.GetByFilter("APARTMENT_NUMBER = ?", "33")
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	for _, en := range ens {
		ca := en.(*entity.CustomerAccount)
		assert.Equal(t, "33", ca.ApartmentNumber)
	}

	ens2, err := customerAccountDao.GetByFilter("CUSTOMER_TYPE_NAME = ?", "CONSUMER")
	assert.Nil(t, err)
	assert.NotNil(t, ens2)
	assert.Equal(t, 1, len(ens2))

	for _, en := range ens2 {
		ca := en.(*entity.CustomerAccount)
		assert.Equal(t, "CONSUMER", ca.CustomerTypeName)
	}

}

func TestQueryCustomerAccounts(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount1 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "33",
	}

	customerAccount2 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Status:           false,
		ApartmentNumber:  "34",
	}

	customerAccount3 := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "CONSUMER",
		Status:           false,
		ApartmentNumber:  "35",
	}

	customerAccountEntity1, err := customerAccountDao.Create(customerAccount1)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity1)

	customerAccountEntity2, err := customerAccountDao.Create(customerAccount2)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity2)

	customerAccountEntity3, err := customerAccountDao.Create(customerAccount3)
	assert.Nil(t, err)
	assert.NotNil(t, customerAccountEntity3)

	// "SELECT * FROM `CUSTOMER_ACCOUNT` WHERE `CUSTOMER_TYPE_NAME` = 'PROSUMER' ORDER BY `APARTMENT_NUMBER` ASC"
	count, ens, err := customerAccountDao.Query(
		entity.CustomerAccount{ProviderID: 1},
		&mysql.Query{
			Limit:  10,
			Offset: 0,
			Order:  "`APARTMENT_NUMBER` ASC",
			Filter: "`CUSTOMER_TYPE_NAME` = 'PROSUMER'",
		})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, count)
	assert.Equal(t, 2, len(ens))

	assert.Equal(t, customerAccountEntity1, ens[0])
	assert.Equal(t, customerAccountEntity2, ens[1])
}
