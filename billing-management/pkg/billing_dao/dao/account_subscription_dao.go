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

type accountSubscriptionDao struct {
	bd.Session
}

func NewAccountSubscriptionDao(s bd.Session) *accountSubscriptionDao {
	return &accountSubscriptionDao{s}
}

func (as *accountSubscriptionDao) NewEntity() entity.Entity {
	return entity.NewAccountSubscription()
}

func (as *accountSubscriptionDao) GetByNameAndAccountID(name string, id int) (entity.Entity, error) {
	var dbAccountSubscription entity.AccountSubscription
	err := as.Where(entity.AccountSubscription{Name: name, AccountID: id}).Take(&dbAccountSubscription).Error()
	return &dbAccountSubscription, err
}

func (as *accountSubscriptionDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbAccountSubscriptions []entity.AccountSubscription
	err := as.Where(query, args).Find(&dbAccountSubscriptions).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptions))
	for i, v := range dbAccountSubscriptions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (as *accountSubscriptionDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := as.Table(as.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbAccountSubscriptions []entity.AccountSubscription
	err = as.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbAccountSubscriptions).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptions))
	for i, v := range dbAccountSubscriptions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (as *accountSubscriptionDao) List() ([]entity.Entity, error) {
	var dbAccountSubscriptions []entity.AccountSubscription
	err := as.Find(&dbAccountSubscriptions).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptions))
	for i, v := range dbAccountSubscriptions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (as *accountSubscriptionDao) Create(e entity.Entity) (entity.Entity, error) {
	err := as.Session.Create(e).Error()
	return e, err
}

func (as *accountSubscriptionDao) UpdateByNameAndAccountID(name string, id int, e entity.Entity) (entity.Entity, error) {
	err := as.Where(&entity.AccountSubscription{Name: name, AccountID: id}).Updates(e).Error()
	return e, err
}

func (as *accountSubscriptionDao) DeleteByNameAndAccountID(name string, id int) error {
	return as.Where(&entity.AccountSubscription{Name: name, AccountID: id}).Delete(&entity.AccountSubscription{}).Error()
}

func (as *accountSubscriptionDao) BeginTransaction() bd.Session {
	return as.Begin()
}

func (as *accountSubscriptionDao) Get(i interface{}) (entity.Entity, error) {
	var dbAccountSubscription entity.AccountSubscription
	err := as.Where(i).Take(&dbAccountSubscription).Error()
	return &dbAccountSubscription, err
}

func (as *accountSubscriptionDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (as *accountSubscriptionDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (as *accountSubscriptionDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (as *accountSubscriptionDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (as *accountSubscriptionDao) DeleteByID(i int) error {
	return nil
}
