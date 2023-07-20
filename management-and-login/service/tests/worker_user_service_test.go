package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	wu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"testing"
)

func TestCreateWorkerUserService(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"

	workerUserModel := &wu.WorkerUser{
		ProviderID:           1,
		Email:                "jan.kowalski@email.com",
		BlockchainAccAddress: "jan.kowalski",
		Login:                &login,
		Password:             &password,
	}

	workerUserService := sf.New(service.WORKER_USER)

	model, err := workerUserService.Create(workerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	workerUserModelResp, ok := model.(*wu.WorkerUser)
	assert.True(t, ok)
	assert.NotNil(t, workerUserModelResp)
	assert.NotNil(t, workerUserModelResp.WorkerID)
	assert.NotNil(t, workerUserModelResp.UserID)
	assert.Equal(t, workerUserModel.ProviderID, workerUserModelResp.ProviderID)
	assert.Equal(t, workerUserModel.Email, workerUserModelResp.Email)
	assert.Equal(t, workerUserModel.BlockchainAccAddress, workerUserModelResp.BlockchainAccAddress)
	assert.Equal(t, workerUserModel.Login, workerUserModelResp.Login)
	assert.NotEqual(t, password, workerUserModelResp.Password)
}

func TestWorkerServiceDeleteWithUsers(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"

	workerUserModel := &wu.WorkerUser{
		ProviderID:           1,
		Email:                "jan.kowalski@email.com",
		BlockchainAccAddress: "jan.kowalski",
		Login:                &login,
		Password:             &password,
	}

	workerUserService := sf.New(service.WORKER_USER)

	model, err := workerUserService.Create(workerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	workerUser := model.(*wu.WorkerUser)

	workerService := sf.New(service.WORKER)
	workerServiceImpl := workerService.(impl.WorkerService)

	err = workerServiceImpl.DeleteWithUser(workerUser.WorkerID)
	assert.Nil(t, err)
}

func TestWorkerServiceDeleteWithUsersFailHasCustomer(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"

	workerUserModel := &wu.WorkerUser{
		ProviderID:           1,
		Email:                "jan.kowalski@email.com",
		BlockchainAccAddress: "jan.kowalski",
		Login:                &login,
		Password:             &password,
	}

	workerUserService := sf.New(service.WORKER_USER)

	model, err := workerUserService.Create(workerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	workerUser := model.(*wu.WorkerUser)

	loginCustomer := "janek.kowalski"
	passwordCustomer := "1234fss12"

	customerUserModel := &customer_user.CustomerUser{
		FirstName:        "Janek",
		LastName:         "Kowalski",
		Email:            "janek.kowalski@email.com",
		CustomerTypeName: "PROSUMER",
		ProviderID:       1,
		Login:            &loginCustomer,
		Password:         &passwordCustomer,
		RoleID:           int(enum.PROSUMER),
		WorkerID:         workerUser.WorkerID,
	}

	customerUserService := sf.New(service.CUSTOMER_USER)
	customerUser, err := customerUserService.Create(customerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, customerUser)

	workerService := sf.New(service.WORKER)
	workerServiceImpl := workerService.(impl.WorkerService)

	err = workerServiceImpl.DeleteWithUser(workerUser.WorkerID)
	assert.Error(t, err)

	en, err := workerService.GetByID(workerUser.WorkerID)
	assert.NotNil(t, en)
	assert.Nil(t, err)

	userService := sf.New(service.USER)
	en, err = userService.GetByID(workerUser.UserID)
	assert.NotNil(t, en)
	assert.Nil(t, err)
}
