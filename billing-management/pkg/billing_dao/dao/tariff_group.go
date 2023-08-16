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

type tariffGroupDao struct {
	bd.Session
}

func NewTariffGroupDao(s bd.Session) *tariffGroupDao {
	return &tariffGroupDao{s}
}

func (r *tariffGroupDao) NewEntity() entity.Entity {
	return entity.NewTariffGroup()
}

func (r *tariffGroupDao) GetByID(i int) (entity.Entity, error) {
	var dbTariffGroup entity.TariffGroup
	err := r.Where(entity.TariffGroup{ID: i}).Take(&dbTariffGroup).Error()
	return &dbTariffGroup, err
}

func (r *tariffGroupDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbTariffGroups []entity.TariffGroup
	err := r.Where(query, args).Find(&dbTariffGroups).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroups))
	for i, v := range dbTariffGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (r *tariffGroupDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := r.Table(r.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbTariffGroups []entity.TariffGroup
	err = r.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbTariffGroups).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroups))
	for i, v := range dbTariffGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (r *tariffGroupDao) List() ([]entity.Entity, error) {
	var dbTariffGroups []entity.TariffGroup
	err := r.Find(&dbTariffGroups).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroups))
	for i, v := range dbTariffGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (r *tariffGroupDao) Create(e entity.Entity) (entity.Entity, error) {
	err := r.Session.Create(e).Error()
	return e, err
}

func (r *tariffGroupDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := r.Where(&entity.TariffGroup{ID: i}).Updates(e).Error()
	return e, err
}

func (r *tariffGroupDao) DeleteByID(i int) error {
	return r.Where(&entity.TariffGroup{ID: i}).Delete(&entity.TariffGroup{}).Error()
}

func (r *tariffGroupDao) BeginTransaction() bd.Session {
	return r.Begin()
}

func (r *tariffGroupDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (r *tariffGroupDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (r *tariffGroupDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
