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

type serviceAccessPointDao struct {
	bd.Session
}

func NewServiceAccessPointDao(s bd.Session) *serviceAccessPointDao {
	return &serviceAccessPointDao{s}
}

func (sap *serviceAccessPointDao) NewEntity() entity.Entity {
	return entity.NewServiceAccessPoint()
}

func (sap *serviceAccessPointDao) GetByID(i int) (entity.Entity, error) {
	var dbServiceAccessPoint entity.ServiceAccessPoint
	err := sap.Where(entity.ServiceAccessPoint{ID: i}).Take(&dbServiceAccessPoint).Error()
	return &dbServiceAccessPoint, err
}

func (sap *serviceAccessPointDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbServiceAccessPoints []entity.ServiceAccessPoint
	err := sap.Where(query, args).Find(&dbServiceAccessPoints).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPoints))
	for i, v := range dbServiceAccessPoints {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sap *serviceAccessPointDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := sap.Table(sap.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbServiceAccessPoints []entity.ServiceAccessPoint
	err = sap.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbServiceAccessPoints).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPoints))
	for i, v := range dbServiceAccessPoints {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (sap *serviceAccessPointDao) List() ([]entity.Entity, error) {
	var dbServiceAccessPoints []entity.ServiceAccessPoint
	err := sap.Find(&dbServiceAccessPoints).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPoints))
	for i, v := range dbServiceAccessPoints {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sap *serviceAccessPointDao) Create(e entity.Entity) (entity.Entity, error) {
	err := sap.Session.Create(e).Error()
	return e, err
}

func (sap *serviceAccessPointDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := sap.Where(&entity.ServiceAccessPoint{ID: i}).Updates(e).Error()
	return e, err
}

func (sap *serviceAccessPointDao) DeleteByID(i int) error {
	return sap.Where(&entity.ServiceAccessPoint{ID: i}).Delete(&entity.ServiceAccessPoint{}).Error()
}

func (sap *serviceAccessPointDao) BeginTransaction() bd.Session {
	return sap.Begin()
}

func (sap *serviceAccessPointDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sap *serviceAccessPointDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sap *serviceAccessPointDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
