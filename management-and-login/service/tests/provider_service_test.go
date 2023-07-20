package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/provider"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"testing"
)

func TestProviderServiceCreate(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	providerModelResp, ok := model.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp)
	assert.Equal(t, name, providerModelResp.Name)
	assert.Equal(t, nip, providerModelResp.NIP)
}

func TestProviderServiceDeleteByID(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	providerModelResp, ok := model.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp)
	assert.Equal(t, name, providerModelResp.Name)
	assert.Equal(t, nip, providerModelResp.NIP)

	providerModelResp2, ok := model2.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp2)
	assert.Equal(t, name2, providerModelResp2.Name)
	assert.Equal(t, nip2, providerModelResp2.NIP)

	err = providerService.DeleteByID(providerModelResp.ID)
	assert.Nil(t, err)

	getModel2, err := providerService.GetByID(providerModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	getModel1, err := providerService.GetByID(providerModelResp.ID)
	assert.Error(t, err)
	assert.Nil(t, getModel1)
}

func TestProviderServiceGetByID(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	providerModelResp, ok := model.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp)
	assert.Equal(t, name, providerModelResp.Name)
	assert.Equal(t, nip, providerModelResp.NIP)

	providerModelResp2, ok := model2.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp2)
	assert.Equal(t, name2, providerModelResp2.Name)
	assert.Equal(t, nip2, providerModelResp2.NIP)

	getModel1, err := providerService.GetByID(providerModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel1)

	providerGetModelResp, ok := getModel1.(*provider.Provider)
	assert.True(t, ok)
	assert.Equal(t, name, providerGetModelResp.Name)
	assert.Equal(t, nip, providerGetModelResp.NIP)

	getModel2, err := providerService.GetByID(providerModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)

	providerGetModelResp2, ok := getModel2.(*provider.Provider)
	assert.True(t, ok)
	assert.Equal(t, name2, providerGetModelResp2.Name)
	assert.Equal(t, nip2, providerGetModelResp2.NIP)
}

func TestProviderServiceGetWithFilter(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := providerService.GetWithFilter("NAME = ?", name2)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))

	providerModelResp, ok := ens[0].(*provider.Provider)
	assert.True(t, ok)
	assert.Equal(t, name2, providerModelResp.Name)
	assert.Equal(t, nip2, providerModelResp.NIP)
}

func TestProviderServiceList(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := providerService.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 3, len(ens))
}

func TestProviderServiceUpdateByID(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"
	email := "ovoo@ovoo.com"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	providerModelResp, ok := model.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp)
	assert.Equal(t, name, providerModelResp.Name)
	assert.Equal(t, nip, providerModelResp.NIP)

	providerModelResp2, ok := model2.(*provider.Provider)
	assert.True(t, ok)
	assert.NotNil(t, providerModelResp2)
	assert.Equal(t, name2, providerModelResp2.Name)
	assert.Equal(t, nip2, providerModelResp2.NIP)

	updateModel, err := providerService.UpdateByID(providerModelResp.ID, &provider.Provider{
		Email: email,
	})
	assert.Nil(t, err)
	assert.NotNil(t, updateModel)

	getModel, err := providerService.GetByID(providerModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel)

	providerGetModelResp, ok := getModel.(*provider.Provider)
	assert.True(t, ok)
	assert.Equal(t, name, providerGetModelResp.Name)
	assert.Equal(t, nip, providerGetModelResp.NIP)
	assert.Equal(t, email, providerGetModelResp.Email)
}

func TestProviderServiceDeleteWithAdmins(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	provider1 := model.(*provider.Provider)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	provider2 := model2.(*provider.Provider)

	login := "jan.kowalski"
	password := "1234fss1"

	workerUserModel := &worker_user.WorkerUser{
		FirstName:  "Jan",
		LastName:   "Kowalski",
		Email:      "jan.kowalski@email.com",
		ProviderID: provider1.ID,
		Login:      &login,
		Password:   &password,
		RoleID:     int(enum.ADMINISTRATOR_FULL),
	}

	workerUserService := sf.New(service.WORKER_USER)
	workerUserService.Create(workerUserModel)

	providerServiceImpl := providerService.(impl.ProviderService)
	err = providerServiceImpl.DeleteWithAdmins(provider1.ID)
	assert.Nil(t, err)

	getModel2, err := providerService.GetByID(provider2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	getModel1, err := providerService.GetByID(provider1.ID)
	assert.Error(t, err)
	assert.Nil(t, getModel1)
}

func TestProviderServiceDeleteWithAdminsFailHasCustomer(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	provider1 := model.(*provider.Provider)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	provider2 := model2.(*provider.Provider)

	login := "jan.kowalski"
	password := "1234fss1"

	workerUserModel := &worker_user.WorkerUser{
		FirstName:  "Jan",
		LastName:   "Kowalski",
		Email:      "jan.kowalski@email.com",
		ProviderID: provider1.ID,
		Login:      &login,
		Password:   &password,
		RoleID:     int(enum.ADMINISTRATOR_FULL),
	}

	workerUserService := sf.New(service.WORKER_USER)
	workerUser, err := workerUserService.Create(workerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, workerUser)

	workerUserStruct := workerUser.(*worker_user.WorkerUser)

	loginCustomer := "janek.kowalski"
	passwordCustomer := "1234fss12"

	customerUserModel := &customer_user.CustomerUser{
		FirstName:        "Janek",
		LastName:         "Kowalski",
		Email:            "janek.kowalski@email.com",
		CustomerTypeName: "PROSUMER",
		ProviderID:       provider1.ID,
		Login:            &loginCustomer,
		Password:         &passwordCustomer,
		RoleID:           int(enum.PROSUMER),
		WorkerID:         workerUserStruct.WorkerID,
	}

	customerUserService := sf.New(service.CUSTOMER_USER)
	customer, err := customerUserService.Create(customerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	providerServiceImpl := providerService.(impl.ProviderService)
	err = providerServiceImpl.DeleteWithAdmins(provider1.ID)
	assert.Error(t, err)

	getModel2, err := providerService.GetByID(provider2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	getModel1, err := providerService.GetByID(provider1.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel1)
}

func TestProviderServiceQuery(t *testing.T) {
	mariaDBFlushData(s)

	name := "Ovoo"
	nip := "6762457439"

	providerModel := &provider.Provider{
		Name: name,
		NIP:  nip,
	}

	name2 := "Keno"
	nip2 := "6312671983"

	providerModel2 := &provider.Provider{
		Name: name2,
		NIP:  nip2,
	}

	name3 := "Tauron"
	nip3 := "9542583988"

	providerModel3 := &provider.Provider{
		Name: name3,
		NIP:  nip3,
	}

	providerService := sf.New(service.PROVIDER)

	model, err := providerService.Create(providerModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := providerService.Create(providerModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	model3, err := providerService.Create(providerModel3)
	assert.Nil(t, err)
	assert.NotNil(t, model3)

	queryUrl := "https://example.com/api/management/providers?sort=asc:nip&filterFields=name&filterValues=Tauron"
	query := api_utils.ParseQuery("PROVIDER", queryUrl)
	count, ens, err := providerService.Query(provider.Provider{}, query)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, count)

	providerModelResp, ok := ens[0].(*provider.Provider)
	assert.True(t, ok)
	assert.Equal(t, name3, providerModelResp.Name)
	assert.Equal(t, nip3, providerModelResp.NIP)
}
