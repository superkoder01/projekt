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

type contractDao struct {
	bd.Session
}

func NewContractDao(s bd.Session) *contractDao {
	return &contractDao{s}
}

func (c *contractDao) NewEntity() entity.Entity {
	return entity.NewContract()
}

func (c *contractDao) GetByID(i int) (entity.Entity, error) {
	var dbContract entity.Contract
	err := c.Where(entity.Contract{ID: i}).Take(&dbContract).Error()
	return &dbContract, err
}

func (c *contractDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbContracts []entity.Contract
	err := c.Where(query, args).Find(&dbContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *contractDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := c.Table(c.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbContracts []entity.Contract
	err = c.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbContracts).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (c *contractDao) List() ([]entity.Entity, error) {
	var dbContracts []entity.Contract
	err := c.Find(&dbContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *contractDao) Create(e entity.Entity) (entity.Entity, error) {
	err := c.Session.Create(e).Error()
	return e, err
}

func (c *contractDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := c.Where(&entity.Contract{ID: i}).Updates(e).Error()
	return e, err
}

func (c *contractDao) DeleteByID(i int) error {
	return c.Where(&entity.Contract{ID: i}).Delete(&entity.Contract{}).Error()
}

func (c *contractDao) BeginTransaction() bd.Session {
	return c.Begin()
}

func (c *contractDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *contractDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *contractDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
