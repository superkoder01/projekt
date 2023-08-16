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

type offerStatusesDao struct {
	bd.Session
}

func NewOfferStatusesDao(s bd.Session) *offerStatusesDao {
	return &offerStatusesDao{s}
}

func (bt *offerStatusesDao) NewEntity() entity.Entity {
	return entity.NewOfferStatuses()
}

func (bt *offerStatusesDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbOfferStatuses []entity.OfferStatuses
	err := bt.Where(query, args).Find(&dbOfferStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbOfferStatuses))
	for i, v := range dbOfferStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *offerStatusesDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := bt.Table(bt.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbOfferStatuses []entity.OfferStatuses
	err = bt.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbOfferStatuses).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbOfferStatuses))
	for i, v := range dbOfferStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (bt *offerStatusesDao) List() ([]entity.Entity, error) {
	var dbOfferStatuses []entity.OfferStatuses
	err := bt.Find(&dbOfferStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbOfferStatuses))
	for i, v := range dbOfferStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *offerStatusesDao) Create(e entity.Entity) (entity.Entity, error) {
	err := bt.Session.Create(e).Error()
	return e, err
}

func (bt *offerStatusesDao) DeleteByName(name string) error {
	return bt.Where(&entity.OfferStatuses{Name: name}).Delete(&entity.OfferStatuses{}).Error()
}

func (bt *offerStatusesDao) BeginTransaction() bd.Session {
	return bt.Begin()
}

func (bt *offerStatusesDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *offerStatusesDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *offerStatusesDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *offerStatusesDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (bt *offerStatusesDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (bt *offerStatusesDao) DeleteByID(i int) error {
	return nil
}
