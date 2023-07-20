package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"testing"
)

func TestCustomerAccountServiceQuery(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountModel3 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "CONSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	model3, err := customerAccountService.Create(customerAccountModel3)
	assert.Nil(t, err)
	assert.NotNil(t, model3)

	queryUrl := "https://example.com/api/management/customerAccounts?&sort=asc:apartmentNumber&filterFields=customerTypeName&filterValues=PROSUMER"
	query := api_utils.ParseQuery("CUSTOMER_ACCOUNT", queryUrl)
	count, mdls, err := customerAccountService.Query(customer_account.CustomerAccount{ProviderID: 1}, query)
	assert.Nil(t, err)
	assert.NotNil(t, mdls)
	assert.Equal(t, 2, count)
}

func TestCustomerAccountServiceCreate(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	customerAccountModelResp, ok := model.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp)
	assert.Equal(t, firstName, customerAccountModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountModelResp.LastName)
}

func TestCustomerAccountServiceDeleteByID(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	customerAccountModelResp, ok := model.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp)
	assert.Equal(t, firstName, customerAccountModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountModelResp.LastName)

	customerAccountModelResp2, ok := model2.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp2)
	assert.Equal(t, firstName2, customerAccountModelResp2.FirstName)
	assert.Equal(t, lastName2, customerAccountModelResp2.LastName)

	err = customerAccountService.DeleteByID(customerAccountModelResp.ID)
	assert.Nil(t, err)

	getModel2, err := customerAccountService.GetByID(customerAccountModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	getModel1, err := customerAccountService.GetByID(customerAccountModelResp.ID)
	assert.Error(t, err)
	assert.Nil(t, getModel1)
}

func TestCustomerAccountServiceGetByID(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	customerAccountModelResp, ok := model.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp)
	assert.Equal(t, firstName, customerAccountModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountModelResp.LastName)

	customerAccountModelResp2, ok := model2.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp2)
	assert.Equal(t, firstName2, customerAccountModelResp2.FirstName)
	assert.Equal(t, lastName2, customerAccountModelResp2.LastName)

	getModel1, err := customerAccountService.GetByID(customerAccountModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel1)

	customerAccountGetModelResp, ok := getModel1.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.Equal(t, firstName, customerAccountGetModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountGetModelResp.LastName)

	getModel2, err := customerAccountService.GetByID(customerAccountModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)

	customerAccountGetModelResp2, ok := getModel2.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.Equal(t, firstName2, customerAccountGetModelResp2.FirstName)
	assert.Equal(t, lastName2, customerAccountGetModelResp2.LastName)
}

func TestCustomerAccountServiceGetWithFilter(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := customerAccountService.GetWithFilter("FIRST_NAME = ?", firstName2)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))

	customerAccountModelResp, ok := ens[0].(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.Equal(t, firstName2, customerAccountModelResp.FirstName)
	assert.Equal(t, lastName2, customerAccountModelResp.LastName)
}

func TestCustomerAccountServiceList(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := customerAccountService.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestCustomerAccountServiceUpdateByID(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"
	email := "jan.kowalski@email.com"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	customerAccountModelResp, ok := model.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp)
	assert.Equal(t, firstName, customerAccountModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountModelResp.LastName)

	customerAccountModelResp2, ok := model2.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.NotNil(t, customerAccountModelResp2)
	assert.Equal(t, firstName2, customerAccountModelResp2.FirstName)
	assert.Equal(t, lastName2, customerAccountModelResp2.LastName)

	updateModel, err := customerAccountService.UpdateByID(customerAccountModelResp.ID, &customer_account.CustomerAccount{
		Email: email,
	})
	assert.Nil(t, err)
	assert.NotNil(t, updateModel)

	getModel, err := customerAccountService.GetByID(customerAccountModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel)

	customerAccountGetModelResp, ok := getModel.(*customer_account.CustomerAccount)
	assert.True(t, ok)
	assert.Equal(t, firstName, customerAccountGetModelResp.FirstName)
	assert.Equal(t, lastName, customerAccountGetModelResp.LastName)
	assert.Equal(t, email, customerAccountGetModelResp.Email)
}
