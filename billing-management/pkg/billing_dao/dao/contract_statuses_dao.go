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

type contractStatusesDao struct {
	bd.Session
}

func NewContractStatusesDao(s bd.Session) *contractStatusesDao {
	return &contractStatusesDao{s}
}

func (bt *contractStatusesDao) NewEntity() entity.Entity {
	return entity.NewContractStatuses()
}

func (bt *contractStatusesDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbContractStatuses []entity.ContractStatuses
	err := bt.Where(query, args).Find(&dbContractStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *contractStatusesDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := bt.Table(bt.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbContractStatuses []entity.ContractStatuses
	err = bt.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbContractStatuses).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (bt *contractStatusesDao) List() ([]entity.Entity, error) {
	var dbContractStatuses []entity.ContractStatuses
	err := bt.Find(&dbContractStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *contractStatusesDao) Create(e entity.Entity) (entity.Entity, error) {
	err := bt.Session.Create(e).Error()
	return e, err
}

func (bt *contractStatusesDao) DeleteByName(name string) error {
	return bt.Where(&entity.ContractStatuses{Name: name}).Delete(&entity.ContractStatuses{}).Error()
}

func (bt *contractStatusesDao) BeginTransaction() bd.Session {
	return bt.Begin()
}

func (bt *contractStatusesDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) DeleteByID(i int) error {
	return nil
}
