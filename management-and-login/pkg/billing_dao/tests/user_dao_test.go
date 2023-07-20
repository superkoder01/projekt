package tests

import (
	"github.com/stretchr/testify/assert"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)
	user := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
	}

	userEntity, err := userDao.Create(user)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user, userEntityE)
}

func TestGetUsers(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)

	user1 := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		Email:      "testemail",
	}

	user2 := &entity.User{
		ProviderID: 1,
		Login:      "login102",
		Password:   "pass2!",
		Email:      "testemail2",
	}

	userEntity1, err := userDao.Create(user1)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity1)

	userEntity2, err := userDao.Create(user2)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity2)

	ens, err := userDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	assert.Equal(t, user1, userEntity1)
	assert.Equal(t, user2, userEntity2)
}

func TestUpdateUserByID(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)

	user1 := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		Email:      "testemail",
	}

	user2 := &entity.User{
		ProviderID: 1,
		Login:      "login102",
		Password:   "pass2!",
		Email:      "testemail2",
	}

	userEntity, err := userDao.Create(user1)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user1, userEntityE)

	userEntity2, err := userDao.Create(user2)
	userEntityE2 := userEntity2.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity2)
	assert.Equal(t, user2, userEntityE2)

	user1.SetMustChangePassword(true)
	user1.SetActive(true)
	user1.SetPassword("newPassword321!")
	userEntity, err = userDao.UpdateByID(userEntityE.ID, user1)
	userEntityE = userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user1, userEntityE)

	userEntity2, err = userDao.GetByID(user2.ID)
	assert.Nil(t, err)
	assert.Equal(t, false, userEntityE2.MustChangePassword)
}

func TestDeleteUserByID(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)

	user1 := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		Email:      "testemail",
	}

	user2 := &entity.User{
		ProviderID: 1,
		Login:      "login102",
		Password:   "pass2!",
		Email:      "testemail2",
	}

	user3 := &entity.User{
		ProviderID: 1,
		Login:      "login103",
		Password:   "pass3!",
		Email:      "testemail3",
	}

	userEntity, err := userDao.Create(user1)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user1, userEntityE)

	userEntity2, err := userDao.Create(user2)
	userEntityE2 := userEntity2.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity2)
	assert.Equal(t, user2, userEntityE2)

	userEntity3, err := userDao.Create(user3)
	userEntityE3 := userEntity3.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity3)
	assert.Equal(t, user3, userEntityE3)

	err = userDao.DeleteByID(userEntityE2.ID)
	assert.Nil(t, err)

	ens, err := userDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestGetUsersByFilter(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)

	user1 := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		Active:     false,
		Email:      "testemail",
	}

	user2 := &entity.User{
		ProviderID: 1,
		Login:      "login102",
		Password:   "pass2!",
		Active:     true,
		Email:      "testemail2",
	}

	user3 := &entity.User{
		ProviderID: 1,
		Login:      "login103",
		Password:   "pass3!",
		Active:     true,
		Email:      "testemail3",
	}

	userEntity1, err := userDao.Create(user1)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity1)

	userEntity2, err := userDao.Create(user2)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity2)

	userEntity3, err := userDao.Create(user3)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity3)

	ens, err := userDao.GetByFilter(entity.User{Active: true})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	for _, en := range ens {
		u := en.(*entity.User)
		assert.Equal(t, true, u.Active)
	}
}

func TestUserCustomerAssociation(t *testing.T) {
	mariaDBFlushData(s)
	customerAccountDao := df.New(bd.CUSTOMER_ACCOUNT)

	customerAccount := &entity.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		Email:            "jan.kowalski@email.com",
	}
	customerAccountDao.Create(customerAccount)

	userDao := df.New(bd.USER)
	user := &entity.User{
		ProviderID:        1,
		Login:             "login101",
		Password:          "pass1!",
		CustomerAccountID: customerAccount.ID,
		Email:             "testemail",
	}

	userEntity, err := userDao.Create(user)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user, userEntityE)

	getUser, err := userDao.GetByID(userEntityE.ID)
	userEntityE = getUser.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.NotNil(t, userEntityE.CustomerAccount)
	assert.Equal(t, "jan.kowalski@email.com", userEntityE.CustomerAccount.Email)
}

func TestUserWorkerAssociation(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker := &entity.Worker{
		ProviderID: 1,
		Email:      "jan.kowalski@email.com",
		FirstName:  "Jan",
		LastName:   "Kowalski",
	}
	workerDao.Create(worker)

	userDao := df.New(bd.USER)
	user := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		WorkerID:   worker.ID,
		Email:      "testemail",
	}

	userEntity, err := userDao.Create(user)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user, userEntityE)

	userList, err := userDao.List()
	userEntityE = userList[0].(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.NotNil(t, userEntityE.Worker)
	assert.Equal(t, "Jan", userEntityE.Worker.FirstName)
	assert.Equal(t, "Kowalski", userEntityE.Worker.LastName)
}

func TestQueryUsers(t *testing.T) {
	mariaDBFlushData(s)
	userDao := df.New(bd.USER)

	user1 := &entity.User{
		ProviderID: 1,
		Login:      "login101",
		Password:   "pass1!",
		Email:      "testemail",
	}

	user2 := &entity.User{
		ProviderID: 1,
		Login:      "login102",
		Password:   "pass2!",
		Email:      "testemail2",
	}

	user3 := &entity.User{
		ProviderID: 1,
		Login:      "login103",
		Password:   "pass3!",
		Email:      "testemail3",
	}

	userEntity1, err := userDao.Create(user1)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity1)

	userEntity2, err := userDao.Create(user2)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity2)

	userEntity3, err := userDao.Create(user3)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity3)

	// "SELECT * FROM `USER` LIMIT 1 OFFSET 1"
	count, ens, err := userDao.Query(
		entity.User{
			ProviderID: 1,
		},
		&mysql.Query{
			Limit:  1,
			Offset: 1,
			Order:  "",
			Filter: entity.User{},
		})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 3, count)
	assert.Equal(t, 1, len(ens))

	assert.Equal(t, userEntity2, ens[0])
}
