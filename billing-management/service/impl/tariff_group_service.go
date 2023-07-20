package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/tariff_group"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
)

type tariffGroupService struct {
	Dao dao.Dao
}

func NewTariffGroupService(dao dao.Dao) *tariffGroupService {
	return &tariffGroupService{Dao: dao}
}

func (c *tariffGroupService) Create(model model.Model) (model.Model, error) {
	soModel, ok := model.(*tariff_group.TariffGroup)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	customerEntity, err := c.ModelToEntity(soModel)
	if err != nil {
		return nil, err
	}

	en, err := c.Dao.Create(customerEntity)
	if err != nil {
		return nil, e.ErrInternalServerError
	}

	return c.EntityToModel(en)
}

func (c *tariffGroupService) DeleteByID(id int) error {
	err := c.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (c *tariffGroupService) UpdateByID(id int, model model.Model) (model.Model, error) {
	soModel, ok := model.(*tariff_group.TariffGroup)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := c.ModelToEntity(soModel)
	if err != nil {
		return nil, err
	}

	en, err = c.Dao.UpdateByID(id, en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *tariffGroupService) GetByID(id int) (model.Model, error) {
	en, err := c.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *tariffGroupService) List() ([]model.Model, error) {
	ens, err := c.Dao.List()
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var soModel model.Model
		var err error

		soModel, err = c.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, soModel)
	}

	return mdls, nil
}

func (c *tariffGroupService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	ens, err := c.Dao.GetByFilter(query, args)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var soModel model.Model
		var err error

		soModel, err = c.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, soModel)
	}

	return mdls, nil
}

func (c *tariffGroupService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdTariffGroup, ok := model.(*tariff_group.TariffGroup)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.TariffGroup{
		ID:                            mdTariffGroup.ID,
		DistributionNetworkOperatorID: mdTariffGroup.DistributionNetworkOperatorID,
		TariffGroupLabelName:          mdTariffGroup.TariffGroupLabelName,
		Name:                          mdTariffGroup.Name,
		StartDate:                     mdTariffGroup.StartDate,
		EndDate:                       mdTariffGroup.EndDate,
	}, nil
}

func (c *tariffGroupService) EntityToModel(en entity.Entity) (model.Model, error) {
	enTariffGroup, ok := en.(*entity.TariffGroup)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &tariff_group.TariffGroup{
		ID:                            enTariffGroup.ID,
		DistributionNetworkOperatorID: enTariffGroup.DistributionNetworkOperatorID,
		TariffGroupLabelName:          enTariffGroup.TariffGroupLabelName,
		Name:                          enTariffGroup.Name,
		StartDate:                     enTariffGroup.StartDate,
		EndDate:                       enTariffGroup.EndDate,
	}, nil
}

func (c *tariffGroupService) Get(i interface{}) (model.Model, error) {
	return nil, nil
}

func (c *tariffGroupService) Update(i interface{}) (model.Model, error) {
	return nil, nil
}

func (c *tariffGroupService) Delete(i interface{}) (model.Model, error) {
	return nil, nil
}
