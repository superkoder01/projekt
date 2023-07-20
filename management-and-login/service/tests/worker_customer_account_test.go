package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"testing"
)

func TestListWorkerCustomerAccounts(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	supervisorModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(supervisorModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)
	supervisor := model.(*worker.Worker)

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
		Supervisor:           supervisor.ID,
	}

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)
	subordinate := model2.(*worker.Worker)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
		WorkerID:         supervisor.ID,
	}

	firstName2 := "Marian"
	lastName2 := "Nowak"

	customerAccountModel2 := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName2,
		LastName:         lastName2,
		WorkerID:         subordinate.ID,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	customer1, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, customer1)

	customer2, err := customerAccountService.Create(customerAccountModel2)
	assert.Nil(t, err)
	assert.NotNil(t, customer2)

	workerCustomerAccountService := sf.New(service.WORKER_CUSTOMER_ACCOUNT).(impl.WorkerCustomerAccountService)

	count, ens, err := workerCustomerAccountService.ListWorkerCustomerAccounts(int(enum.TRADER), 1, supervisor.ID, &apiUtils.Query{})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, count)
}
