package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/distribution_network_operator"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
)

type distributionNetworkOperatorService struct {
	Dao dao.Dao
}

func NewDistributionNetworkOperatorService(dao dao.Dao) *distributionNetworkOperatorService {
	return &distributionNetworkOperatorService{Dao: dao}
}

func (c *distributionNetworkOperatorService) Create(model model.Model) (model.Model, error) {
	soModel, ok := model.(*distribution_network_operator.DistributionNetworkOperator)
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

func (c *distributionNetworkOperatorService) DeleteByID(id int) error {
	err := c.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (c *distributionNetworkOperatorService) UpdateByID(id int, model model.Model) (model.Model, error) {
	soModel, ok := model.(*distribution_network_operator.DistributionNetworkOperator)
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

func (c *distributionNetworkOperatorService) GetByID(id int) (model.Model, error) {
	en, err := c.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *distributionNetworkOperatorService) List() ([]model.Model, error) {
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

func (c *distributionNetworkOperatorService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
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

func (c *distributionNetworkOperatorService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdDistributionNetworkOperator, ok := model.(*distribution_network_operator.DistributionNetworkOperator)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.DistributionNetworkOperator{
		ID:      mdDistributionNetworkOperator.ID,
		Name:    mdDistributionNetworkOperator.Name,
		Nip:     mdDistributionNetworkOperator.Nip,
		Address: mdDistributionNetworkOperator.Address,
		City:    mdDistributionNetworkOperator.City,
	}, nil
}

func (c *distributionNetworkOperatorService) EntityToModel(en entity.Entity) (model.Model, error) {
	enDistributionNetworkOperator, ok := en.(*entity.DistributionNetworkOperator)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &distribution_network_operator.DistributionNetworkOperator{
		ID:      enDistributionNetworkOperator.ID,
		Name:    enDistributionNetworkOperator.Name,
		Nip:     enDistributionNetworkOperator.Nip,
		Address: enDistributionNetworkOperator.Address,
		City:    enDistributionNetworkOperator.City,
	}, nil
}

func (c *distributionNetworkOperatorService) Get(i interface{}) (model.Model, error) {
	return nil, nil
}

func (c *distributionNetworkOperatorService) Update(i interface{}) (model.Model, error) {
	return nil, nil
}

func (c *distributionNetworkOperatorService) Delete(i interface{}) (model.Model, error) {
	return nil, nil
}
