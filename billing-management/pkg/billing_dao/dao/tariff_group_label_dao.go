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

type tariffGroupLabelDao struct {
	bd.Session
}

func NewTariffGroupLabelDao(s bd.Session) *tariffGroupLabelDao {
	return &tariffGroupLabelDao{s}
}

func (tgl *tariffGroupLabelDao) NewEntity() entity.Entity {
	return entity.NewTariffGroupLabel()
}

func (tgl *tariffGroupLabelDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbTariffGroupLabels []entity.TariffGroupLabel
	err := tgl.Where(query, args).Find(&dbTariffGroupLabels).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupLabels))
	for i, v := range dbTariffGroupLabels {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (tgl *tariffGroupLabelDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := tgl.Table(tgl.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbTariffGroupLabels []entity.TariffGroupLabel
	err = tgl.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbTariffGroupLabels).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupLabels))
	for i, v := range dbTariffGroupLabels {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (tgl *tariffGroupLabelDao) List() ([]entity.Entity, error) {
	var dbTariffGroupLabels []entity.TariffGroupLabel
	err := tgl.Find(&dbTariffGroupLabels).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupLabels))
	for i, v := range dbTariffGroupLabels {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (tgl *tariffGroupLabelDao) Create(e entity.Entity) (entity.Entity, error) {
	err := tgl.Session.Create(e).Error()
	return e, err
}

func (tgl *tariffGroupLabelDao) DeleteByName(name string) error {
	return tgl.Where(&entity.TariffGroupLabel{Name: name}).Delete(&entity.TariffGroupLabel{}).Error()
}

func (tgl *tariffGroupLabelDao) BeginTransaction() bd.Session {
	return tgl.Begin()
}

func (tgl *tariffGroupLabelDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (tgl *tariffGroupLabelDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (tgl *tariffGroupLabelDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (tgl *tariffGroupLabelDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (tgl *tariffGroupLabelDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (tgl *tariffGroupLabelDao) DeleteByID(i int) error {
	return nil
}
