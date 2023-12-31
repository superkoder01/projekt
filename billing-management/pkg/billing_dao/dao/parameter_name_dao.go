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

type parameterNameDao struct {
	bd.Session
}

func NewParameterNameDao(s bd.Session) *parameterNameDao {
	return &parameterNameDao{s}
}

func (pn *parameterNameDao) NewEntity() entity.Entity {
	return entity.NewParameterName()
}

func (pn *parameterNameDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbParameterNames []entity.ParameterName
	err := pn.Where(query, args).Find(&dbParameterNames).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbParameterNames))
	for i, v := range dbParameterNames {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (pn *parameterNameDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := pn.Table(pn.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbParameterNames []entity.ParameterName
	err = pn.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbParameterNames).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbParameterNames))
	for i, v := range dbParameterNames {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (pn *parameterNameDao) List() ([]entity.Entity, error) {
	var dbParameterNames []entity.ParameterName
	err := pn.Find(&dbParameterNames).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbParameterNames))
	for i, v := range dbParameterNames {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (pn *parameterNameDao) Create(e entity.Entity) (entity.Entity, error) {
	err := pn.Session.Create(e).Error()
	return e, err
}

func (pn *parameterNameDao) DeleteByName(name string) error {
	return pn.Where(&entity.ParameterName{Name: name}).Delete(&entity.ParameterName{}).Error()
}

func (pn *parameterNameDao) BeginTransaction() bd.Session {
	return pn.Begin()
}

func (pn *parameterNameDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (pn *parameterNameDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (pn *parameterNameDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (pn *parameterNameDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (pn *parameterNameDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (pn *parameterNameDao) DeleteByID(i int) error {
	return nil
}
