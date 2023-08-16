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

type accountBalanceDao struct {
	bd.Session
}

func NewAccountBalanceDao(s bd.Session) *accountBalanceDao {
	return &accountBalanceDao{s}
}

func (ab *accountBalanceDao) NewEntity() entity.Entity {
	return entity.NewAccountBalance()
}

func (ab *accountBalanceDao) GetByID(i int) (entity.Entity, error) {
	var dbAccountBalance entity.AccountBalance
	err := ab.Where(entity.AccountBalance{ID: i}).Take(&dbAccountBalance).Error()
	return &dbAccountBalance, err
}

func (ab *accountBalanceDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbAccountBalances []entity.AccountBalance
	err := ab.Where(query, args).Find(&dbAccountBalances).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountBalances))
	for i, v := range dbAccountBalances {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ab *accountBalanceDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ab.Table(ab.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbAccountBalances []entity.AccountBalance
	err = ab.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbAccountBalances).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbAccountBalances))
	for i, v := range dbAccountBalances {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ab *accountBalanceDao) List() ([]entity.Entity, error) {
	var dbAccountBalances []entity.AccountBalance
	err := ab.Find(&dbAccountBalances).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountBalances))
	for i, v := range dbAccountBalances {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ab *accountBalanceDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ab.Session.Create(e).Error()
	return e, err
}

func (ab *accountBalanceDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := ab.Where(&entity.AccountBalance{ID: i}).Updates(e).Error()
	return e, err
}

func (ab *accountBalanceDao) DeleteByID(i int) error {
	return ab.Where(&entity.AccountBalance{ID: i}).Delete(&entity.AccountBalance{}).Error()
}

func (ab *accountBalanceDao) BeginTransaction() bd.Session {
	return ab.Begin()
}

func (ab *accountBalanceDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ab *accountBalanceDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ab *accountBalanceDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
