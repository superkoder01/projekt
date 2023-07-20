package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"testing"
)

func TestWorkerServiceCreate(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	workerModelResp, ok := model.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp)
	assert.Equal(t, email, workerModelResp.Email)
	assert.Equal(t, blockchainAddress, workerModelResp.BlockchainAccAddress)
}

// TODO return error on getById empty result
func TestWorkerServiceDeleteByID(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	workerModelResp, ok := model.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp)
	assert.Equal(t, email, workerModelResp.Email)
	assert.Equal(t, blockchainAddress, workerModelResp.BlockchainAccAddress)

	workerModelResp2, ok := model2.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp2)
	assert.Equal(t, email2, workerModelResp2.Email)
	assert.Equal(t, blockchainAddress2, workerModelResp2.BlockchainAccAddress)

	err = workerService.DeleteByID(workerModelResp.ID)
	assert.Nil(t, err)

	getModel2, err := workerService.GetByID(workerModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	//TODO
	//getModel1, err := workerService.GetByID(workerModelResp.ID)
	//assert.Error(t, err)
	//assert.Nil(t, getModel1)
}

func TestWorkerServiceGetByID(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	workerModelResp, ok := model.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp)
	assert.Equal(t, email, workerModelResp.Email)
	assert.Equal(t, blockchainAddress, workerModelResp.BlockchainAccAddress)

	workerModelResp2, ok := model2.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp2)
	assert.Equal(t, email2, workerModelResp2.Email)
	assert.Equal(t, blockchainAddress2, workerModelResp2.BlockchainAccAddress)

	getModel1, err := workerService.GetByID(workerModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel1)

	workerGetModelResp, ok := getModel1.(*worker.Worker)
	assert.True(t, ok)
	assert.Equal(t, email, workerGetModelResp.Email)
	assert.Equal(t, blockchainAddress, workerGetModelResp.BlockchainAccAddress)

	getModel2, err := workerService.GetByID(workerModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)

	workerGetModelResp2, ok := getModel2.(*worker.Worker)
	assert.True(t, ok)
	assert.Equal(t, email2, workerGetModelResp2.Email)
	assert.Equal(t, blockchainAddress2, workerGetModelResp2.BlockchainAccAddress)
}

func TestWorkerServiceGetWithFilter(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := workerService.GetWithFilter(worker.Worker{Email: email2})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))

	workerModelResp, ok := ens[0].(*worker.Worker)
	assert.True(t, ok)
	assert.Equal(t, email2, workerModelResp.Email)
	assert.Equal(t, blockchainAddress2, workerModelResp.BlockchainAccAddress)
}

func TestWorkerServiceList(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}
	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := workerService.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestWorkerServiceUpdateByID(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"
	buildingNumber := "22"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	workerModelResp, ok := model.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp)
	assert.Equal(t, email, workerModelResp.Email)
	assert.Equal(t, blockchainAddress, workerModelResp.BlockchainAccAddress)

	workerModelResp2, ok := model2.(*worker.Worker)
	assert.True(t, ok)
	assert.NotNil(t, workerModelResp2)
	assert.Equal(t, email2, workerModelResp2.Email)
	assert.Equal(t, blockchainAddress2, workerModelResp2.BlockchainAccAddress)

	updateModel, err := workerService.UpdateByID(workerModelResp.ID, &worker.Worker{
		BuildingNumber: buildingNumber,
	})
	assert.Nil(t, err)
	assert.NotNil(t, updateModel)

	getModel, err := workerService.GetByID(workerModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel)

	workerGetModelResp, ok := getModel.(*worker.Worker)
	assert.True(t, ok)
	assert.Equal(t, email, workerGetModelResp.Email)
	assert.Equal(t, blockchainAddress, workerGetModelResp.BlockchainAccAddress)
	assert.Equal(t, buildingNumber, workerGetModelResp.BuildingNumber)
}

func TestWorkerServiceQuery(t *testing.T) {
	mariaDBFlushData(s)

	email := "jan.kowalski@ovoo.pl"
	blockchainAddress := "block123"

	workerModel := &worker.Worker{
		ProviderID:           1,
		Email:                email,
		BlockchainAccAddress: blockchainAddress,
	}

	email2 := "marian.nowak@keno.pl"
	blockchainAddress2 := "block124"

	workerModel2 := &worker.Worker{
		ProviderID:           1,
		Email:                email2,
		BlockchainAccAddress: blockchainAddress2,
	}

	email3 := "czeslaw.klopsik@keno.pl"
	blockchainAddress3 := "block125"

	workerModel3 := &worker.Worker{
		ProviderID:           1,
		Email:                email3,
		BlockchainAccAddress: blockchainAddress3,
	}

	workerService := sf.New(service.WORKER)

	model, err := workerService.Create(workerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := workerService.Create(workerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	model3, err := workerService.Create(workerModel3)
	assert.Nil(t, err)
	assert.NotNil(t, model3)

	queryUrl := "https://example.com/api/management/workers?filterFields=email&filterValues=marian.nowak@keno.pl"
	query := api_utils.ParseQuery("WORKER", queryUrl)
	count, ens, err := workerService.Query(worker.Worker{ProviderID: 1}, query)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, count)

	workerModelResp, ok := ens[0].(*worker.Worker)
	assert.True(t, ok)
	assert.Equal(t, email2, workerModelResp.Email)
	assert.Equal(t, blockchainAddress2, workerModelResp.BlockchainAccAddress)
}
