package tests

import (
	"github.com/stretchr/testify/assert"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"testing"
)

func TestCreateProvider(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)
	provider := &entity.Provider{
		Name: "Keno",
		NIP:  "6312671983",
	}

	providerEntity, err := providerDao.Create(provider)
	providerEntityE := providerEntity.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity)
	assert.Equal(t, provider, providerEntityE)
}

func TestGetProviders(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)

	provider1 := &entity.Provider{
		Name: "Keno",
		NIP:  "6312671983",
	}

	provider2 := &entity.Provider{
		Name: "Tauron",
		NIP:  "9542583988",
	}

	providerEntity1, err := providerDao.Create(provider1)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity1)

	providerEntity2, err := providerDao.Create(provider2)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity2)

	ens, err := providerDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))

	assert.Equal(t, provider1, providerEntity1)
	assert.Equal(t, provider2, providerEntity2)
}

func TestUpdateProviderByID(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)

	provider1 := &entity.Provider{
		Name: "Keno",
		NIP:  "6312671983",
	}

	provider2 := &entity.Provider{
		Name: "Tauron",
		NIP:  "9542583988",
	}

	providerEntity, err := providerDao.Create(provider1)
	providerEntityE := providerEntity.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity)
	assert.Equal(t, provider1, providerEntityE)

	providerEntity2, err := providerDao.Create(provider2)
	providerEntityE2 := providerEntity2.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity2)
	assert.Equal(t, provider2, providerEntityE2)

	provider1.SetREGON("367870863")
	providerEntity, err = providerDao.UpdateByID(providerEntityE.ID, provider1)
	providerEntityE = providerEntity.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity)
	assert.Equal(t, provider1, providerEntityE)

	providerEntity2, err = providerDao.GetByID(provider2.ID)
	assert.Nil(t, err)
	assert.Equal(t, "", providerEntityE2.REGON)
}

func TestDeleteProviderByID(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)

	provider1 := &entity.Provider{
		Name: "Keno",
		NIP:  "6312671983",
	}

	provider2 := &entity.Provider{
		Name: "Tauron",
		NIP:  "9542583988",
	}

	provider3 := &entity.Provider{
		Name: "Hymon",
		NIP:  "9930650091",
	}

	providerEntity, err := providerDao.Create(provider1)
	providerEntityE := providerEntity.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity)
	assert.Equal(t, provider1, providerEntityE)

	providerEntity2, err := providerDao.Create(provider2)
	providerEntityE2 := providerEntity2.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity2)
	assert.Equal(t, provider2, providerEntityE2)

	providerEntity3, err := providerDao.Create(provider3)
	providerEntityE3 := providerEntity3.(*entity.Provider)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity3)
	assert.Equal(t, provider3, providerEntityE3)

	err = providerDao.DeleteByID(providerEntityE2.ID)
	assert.Nil(t, err)

	ens, err := providerDao.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestGetProvidersByFilter(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)

	provider1 := &entity.Provider{
		Name:   "Keno",
		NIP:    "6312671983",
		Status: false,
	}

	provider2 := &entity.Provider{
		Name:   "Tauron",
		NIP:    "9542583988",
		Status: false,
	}

	provider3 := &entity.Provider{
		Name:   "Hymon",
		NIP:    "9930650091",
		Status: true,
	}

	providerEntity1, err := providerDao.Create(provider1)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity1)

	providerEntity2, err := providerDao.Create(provider2)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity2)

	providerEntity3, err := providerDao.Create(provider3)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity3)

	ens, err := providerDao.GetByFilter("STATUS = ?", 1)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))
}

func TestQueryProviders(t *testing.T) {
	mariaDBFlushData(s)
	providerDao := df.New(bd.PROVIDER)

	provider1 := &entity.Provider{
		Name: "Keno",
		NIP:  "6312671983",
	}

	provider2 := &entity.Provider{
		Name: "Tauron",
		NIP:  "9542583988",
	}

	provider3 := &entity.Provider{
		Name: "PGE",
		NIP:  "8130268082",
	}

	providerEntity1, err := providerDao.Create(provider1)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity1)

	providerEntity2, err := providerDao.Create(provider2)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity2)

	providerEntity3, err := providerDao.Create(provider3)
	assert.Nil(t, err)
	assert.NotNil(t, providerEntity3)

	// "SELECT * FROM `PROVIDER` ORDER BY `NIP` ASC LIMIT 1"
	count, ens, err := providerDao.Query(
		entity.Provider{},
		&mysql.Query{
			Limit:  1,
			Offset: 0,
			Order:  "`NIP` ASC",
			Filter: entity.Provider{},
		})
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 3, count)
	assert.Equal(t, 1, len(ens))

	assert.Equal(t, providerEntity1, ens[0])
}
