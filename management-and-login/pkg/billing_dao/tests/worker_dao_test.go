package tests

import (
	"github.com/stretchr/testify/assert"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"testing"
)

func TestCreateWorker(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)
	worker := &entity.Worker{
		ProviderID: 1,
		FirstName:  "Jan",
		LastName:   "Kowalski",
	}

	workerEntity, err := workerDao.Create(worker)
	workerEntityE := workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker, workerEntityE)
}

func TestGetWorkers(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker1 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
	}

	worker2 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Masztalski",
		Email:                "jan.masztalski@ovoo.pl",
		BlockchainAccAddress: "jan.masztalski@ovoo.pl",
	}

	workerEntity1, err := workerDao.Create(worker1)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity1)

	workerEntity2, err := workerDao.Create(worker2)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity2)

	ens, err := workerDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	assert.Equal(t, worker1, workerEntity1)
	assert.Equal(t, worker2, workerEntity2)
}

func TestUpdateWorkerByID(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker1 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
	}

	worker2 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Masztalski",
		Email:                "jan.masztalski@ovoo.pl",
		BlockchainAccAddress: "jan.masztalski@ovoo.pl",
	}

	workerEntity, err := workerDao.Create(worker1)
	workerEntityE := workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker1, workerEntityE)

	workerEntity2, err := workerDao.Create(worker2)
	workerEntityE2 := workerEntity2.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity2)
	assert.Equal(t, worker2, workerEntityE2)

	worker1.SetApartmentNumber("33")
	workerEntity, err = workerDao.UpdateByID(workerEntityE.ID, worker1)
	workerEntityE = workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker1, workerEntityE)

	workerEntity2, err = workerDao.GetByID(worker2.ID)
	assert.Nil(t, err)
	assert.Equal(t, "", workerEntityE2.ApartmentNumber)
}

func TestDeleteWorkerByID(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker1 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
	}

	worker2 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Masztalski",
		Email:                "jan.masztalski@ovoo.pl",
		BlockchainAccAddress: "jan.masztalski@ovoo.pl",
	}

	worker3 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Iksinski",
		Email:                "jan.iksinski@ovoo.pl",
		BlockchainAccAddress: "jan.iksinski@ovoo.pl",
	}

	workerEntity, err := workerDao.Create(worker1)
	workerEntityE := workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker1, workerEntityE)

	workerEntity2, err := workerDao.Create(worker2)
	workerEntityE2 := workerEntity2.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity2)
	assert.Equal(t, worker2, workerEntityE2)

	workerEntity3, err := workerDao.Create(worker3)
	workerEntityE3 := workerEntity3.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity3)
	assert.Equal(t, worker3, workerEntityE3)

	err = workerDao.DeleteByID(workerEntityE2.ID)
	assert.Nil(t, err)

	ens, err := workerDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestGetWorkerByFilter(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker1 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
		Country:              "Poland",
	}

	worker2 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Masztalski",
		Email:                "jan.masztalski@ovoo.pl",
		BlockchainAccAddress: "jan.masztalski@ovoo.pl",
		Country:              "Ukraine",
	}

	worker3 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Iksinski",
		Email:                "jan.iksinski@ovoo.pl",
		BlockchainAccAddress: "jan.iksinski@ovoo.pl",
		Country:              "Poland",
	}

	workerEntity, err := workerDao.Create(worker1)
	workerEntityE := workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker1, workerEntityE)

	workerEntity2, err := workerDao.Create(worker2)
	workerEntityE2 := workerEntity2.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity2)
	assert.Equal(t, worker2, workerEntityE2)

	workerEntity3, err := workerDao.Create(worker3)
	workerEntityE3 := workerEntity3.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity3)
	assert.Equal(t, worker3, workerEntityE3)

	ens, err := workerDao.GetByFilter(entity.Worker{Country: "Ukraine"})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))
}

func TestGetWorkerByID(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)
	userDao := df.New(bd.USER)

	worker := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
		Country:              "Poland",
	}

	workerEntity, err := workerDao.Create(worker)
	workerEntityE := workerEntity.(*entity.Worker)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity)
	assert.Equal(t, worker, workerEntityE)

	user := &entity.User{
		Login:    "login",
		Password: "password",
		WorkerID: workerEntityE.ID,
		RoleID:   1,
	}

	userEntity, err := userDao.Create(user)
	userEntityE := userEntity.(*entity.User)
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user, userEntityE)

	workerById, err := workerDao.GetByID(workerEntityE.ID)
	workerByIdE := workerById.(*entity.WorkerJoinUserRole)
	assert.Nil(t, err)
	assert.NotNil(t, workerById)
	assert.Equal(t, 1, workerByIdE.RoleID)
}

func TestQueryWorkers(t *testing.T) {
	mariaDBFlushData(s)
	workerDao := df.New(bd.WORKER)

	worker1 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Kowalski",
		Email:                "jan.kowalski@ovoo.pl",
		BlockchainAccAddress: "jan.kowalski@ovoo.pl",
	}

	worker2 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Masztalski",
		Email:                "jan.masztalski@ovoo.pl",
		BlockchainAccAddress: "jan.masztalski@ovoo.pl",
	}

	worker3 := &entity.Worker{
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Bimber",
		Email:                "jan.bimber@ovoo.pl",
		BlockchainAccAddress: "jan.bimber@ovoo.pl",
	}

	workerEntity1, err := workerDao.Create(worker1)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity1)

	workerEntity2, err := workerDao.Create(worker2)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity2)

	workerEntity3, err := workerDao.Create(worker3)
	assert.Nil(t, err)
	assert.NotNil(t, workerEntity3)

	userWorker3 := &entity.User{
		Login:      "login",
		Password:   "pass",
		Email:      "email@ovoo.pl",
		ProviderID: 1,
		RoleID:     3,
		Active:     true,
		WorkerID:   workerEntity3.(*entity.Worker).ID,
	}

	workerDao.Create(userWorker3)

	// "SELECT * FROM `WORKER` ORDER BY `LAST_NAME` DESC LIMIT 2 OFFSET 2"
	count, ens, err := workerDao.Query(
		entity.Worker{
			ProviderID: 1,
		},
		&mysql.Query{
			Limit:  2,
			Offset: 2,
			Order:  "`LAST_NAME` DESC",
			Filter: entity.Worker{},
		})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 3, count)
	assert.Equal(t, 1, len(ens))

	assert.Equal(t, &entity.WorkerJoinUserRole{
		ID:                   workerEntity3.(*entity.Worker).ID,
		ProviderID:           1,
		FirstName:            "Jan",
		LastName:             "Bimber",
		Email:                "jan.bimber@ovoo.pl",
		BlockchainAccAddress: "jan.bimber@ovoo.pl",
		RoleID:               3,
	}, ens[0])
}
