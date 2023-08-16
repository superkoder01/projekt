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
package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type energyExcessDao struct {
	bd.Session
}

func NewEnergyExcessDao(s bd.Session) *energyExcessDao {
	return &energyExcessDao{s}
}

func (ee *energyExcessDao) NewEntity() entity.Entity {
	return entity.NewEnergyExcess()
}

func (ee *energyExcessDao) GetByID(i int) (entity.Entity, error) {
	var dbVatRate entity.EnergyExcess
	err := ee.Where(entity.EnergyExcess{ID: i}).Take(&dbVatRate).Error()
	return &dbVatRate, err
}

func (ee *energyExcessDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbEnergyExcess []entity.EnergyExcess
	err := ee.Where(query, args).Find(&dbEnergyExcess).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyExcessDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ee.Table(ee.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbEnergyExcess []entity.EnergyExcess
	err = ee.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbEnergyExcess).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ee *energyExcessDao) List() ([]entity.Entity, error) {
	var dbEnergyExcess []entity.EnergyExcess
	err := ee.Find(&dbEnergyExcess).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyExcessDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ee.Session.Create(e).Error()
	return e, err
}

func (ee *energyExcessDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := ee.Where(&entity.EnergyExcess{ID: i}).Updates(e).Error()
	return e, err
}

func (ee *energyExcessDao) DeleteByID(i int) error {
	return ee.Where(&entity.EnergyExcess{ID: i}).Delete(&entity.EnergyExcess{}).Error()
}

func (ee *energyExcessDao) BeginTransaction() bd.Session {
	return ee.Begin()
}

func (ee *energyExcessDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyExcessDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyExcessDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
