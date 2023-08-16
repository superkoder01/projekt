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

type regionDao struct {
	bd.Session
}

func NewRegionDao(s bd.Session) *regionDao {
	return &regionDao{s}
}

func (r *regionDao) NewEntity() entity.Entity {
	return entity.NewRegion()
}

func (r *regionDao) GetByID(i int) (entity.Entity, error) {
	var dbRegion entity.Region
	err := r.Where(entity.Region{ID: i}).Take(&dbRegion).Error()
	return &dbRegion, err
}

func (r *regionDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbRegions []entity.Region
	err := r.Where(query, args).Find(&dbRegions).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRegions))
	for i, v := range dbRegions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (r *regionDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := r.Table(r.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbRegions []entity.Region
	err = r.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbRegions).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbRegions))
	for i, v := range dbRegions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (r *regionDao) List() ([]entity.Entity, error) {
	var dbRegions []entity.Region
	err := r.Find(&dbRegions).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRegions))
	for i, v := range dbRegions {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (r *regionDao) Create(e entity.Entity) (entity.Entity, error) {
	err := r.Session.Create(e).Error()
	return e, err
}

func (r *regionDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := r.Where(&entity.Region{ID: i}).Updates(e).Error()
	return e, err
}

func (r *regionDao) DeleteByID(i int) error {
	return r.Where(&entity.Region{ID: i}).Delete(&entity.Region{}).Error()
}

func (r *regionDao) BeginTransaction() bd.Session {
	return r.Begin()
}

func (r *regionDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (r *regionDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (r *regionDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
