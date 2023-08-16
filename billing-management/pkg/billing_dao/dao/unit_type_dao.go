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

type unitTypeDao struct {
	bd.Session
}

func NewUnitTypeDao(s bd.Session) *unitTypeDao {
	return &unitTypeDao{s}
}

func (ut *unitTypeDao) NewEntity() entity.Entity {
	return entity.NewUnitType()
}

func (ut *unitTypeDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbUnitTypes []entity.UnitType
	err := ut.Where(query, args).Find(&dbUnitTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbUnitTypes))
	for i, v := range dbUnitTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ut *unitTypeDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ut.Table(ut.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbUnitTypes []entity.UnitType
	err = ut.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbUnitTypes).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbUnitTypes))
	for i, v := range dbUnitTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ut *unitTypeDao) List() ([]entity.Entity, error) {
	var dbUnitTypes []entity.UnitType
	err := ut.Find(&dbUnitTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbUnitTypes))
	for i, v := range dbUnitTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ut *unitTypeDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ut.Session.Create(e).Error()
	return e, err
}

func (ut *unitTypeDao) DeleteByName(name entity.UnitTypeName) error {
	return ut.Where(&entity.UnitType{Name: name}).Delete(&entity.UnitType{}).Error()
}

func (ut *unitTypeDao) BeginTransaction() bd.Session {
	return ut.Begin()
}

func (ut *unitTypeDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ut *unitTypeDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ut *unitTypeDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ut *unitTypeDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (ut *unitTypeDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (ut *unitTypeDao) DeleteByID(i int) error {
	return nil
}
