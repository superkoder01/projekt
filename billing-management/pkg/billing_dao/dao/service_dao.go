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

type serviceDao struct {
	bd.Session
}

func NewServiceDao(s bd.Session) *serviceDao {
	return &serviceDao{s}
}

func (s *serviceDao) NewEntity() entity.Entity {
	return entity.NewService()
}

func (s *serviceDao) GetByID(i int) (entity.Entity, error) {
	var dbService entity.Service
	err := s.Where(entity.Service{ID: i}).Take(&dbService).Error()
	return &dbService, err
}

func (s *serviceDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbServices []entity.Service
	err := s.Where(query, args).Find(&dbServices).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServices))
	for i, v := range dbServices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (s *serviceDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := s.Table(s.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbServices []entity.Service
	err = s.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbServices).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbServices))
	for i, v := range dbServices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (s *serviceDao) List() ([]entity.Entity, error) {
	var dbServices []entity.Service
	err := s.Find(&dbServices).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServices))
	for i, v := range dbServices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (s *serviceDao) Create(e entity.Entity) (entity.Entity, error) {
	err := s.Session.Create(e).Error()
	return e, err
}

func (s *serviceDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := s.Where(&entity.Service{ID: i}).Updates(e).Error()
	return e, err
}

func (s *serviceDao) DeleteByID(i int) error {
	return s.Where(&entity.Service{ID: i}).Delete(&entity.Service{}).Error()
}

func (s *serviceDao) BeginTransaction() bd.Session {
	return s.Begin()
}

func (s *serviceDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (s *serviceDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (s *serviceDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
