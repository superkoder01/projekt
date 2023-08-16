/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/tariff_group_parameter"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
)

type tariffGroupParameterService struct {
	Dao dao.Dao
}

func NewTariffGroupParameterService(dao dao.Dao) *tariffGroupParameterService {
	return &tariffGroupParameterService{Dao: dao}
}

func (c *tariffGroupParameterService) Create(model model.Model) (model.Model, error) {
	soModel, ok := model.(*tariff_group_parameter.TariffGroupParameter)
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

func (c *tariffGroupParameterService) DeleteByID(id int) error {
	err := c.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (c *tariffGroupParameterService) UpdateByID(id int, model model.Model) (model.Model, error) {
	soModel, ok := model.(*tariff_group_parameter.TariffGroupParameter)
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

func (c *tariffGroupParameterService) GetByID(id int) (model.Model, error) {
	en, err := c.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *tariffGroupParameterService) List() ([]model.Model, error) {
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

func (c *tariffGroupParameterService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
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

func (c *tariffGroupParameterService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdTariffGroupParameter, ok := model.(*tariff_group_parameter.TariffGroupParameter)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.TariffGroupParameter{
		ID:              mdTariffGroupParameter.ID,
		ParameterNameID: mdTariffGroupParameter.ParameterNameID,
		TariffGroupID:   mdTariffGroupParameter.TariffGroupID,
		Price:           mdTariffGroupParameter.Price,
	}, nil
}

func (c *tariffGroupParameterService) EntityToModel(en entity.Entity) (model.Model, error) {
	enTariffGroupParameter, ok := en.(*entity.TariffGroupParameter)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &tariff_group_parameter.TariffGroupParameter{
		ID:              enTariffGroupParameter.ID,
		TariffGroupID:   enTariffGroupParameter.TariffGroupID,
		ParameterNameID: enTariffGroupParameter.ParameterNameID,
		Price:           enTariffGroupParameter.Price,
	}, nil
}

func (c *tariffGroupParameterService) Get(i interface{}) (model.Model, error) {
	return nil, nil
}

func (c *tariffGroupParameterService) Update(i interface{}) (model.Model, error) {
	soModel, ok := i.(*tariff_group_parameter.TariffGroupParameter)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := c.ModelToEntity(soModel)
	if err != nil {
		return nil, err
	}

	en, err = c.Dao.Update(en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *tariffGroupParameterService) Delete(i interface{}) (model.Model, error) {
	soModel, ok := i.(*tariff_group_parameter.TariffGroupParameter)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := c.ModelToEntity(soModel)
	if err != nil {
		return nil, err
	}

	en, err = c.Dao.Delete(en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}
