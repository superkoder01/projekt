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

type serviceAccessPointContractDao struct {
	bd.Session
}

func NewServiceAccessPointContractDao(s bd.Session) *serviceAccessPointContractDao {
	return &serviceAccessPointContractDao{s}
}

func (sapc *serviceAccessPointContractDao) NewEntity() entity.Entity {
	return entity.NewServiceAccessPointContract()
}

func (sapc *serviceAccessPointContractDao) GetByID(id int) (entity.Entity, error) {
	var dbServiceAccessPointContract entity.ServiceAccessPointContract
	err := sapc.Where(entity.ServiceAccessPointContract{ID: id}).Take(&dbServiceAccessPointContract).Error()
	return &dbServiceAccessPointContract, err
}

func (sapc *serviceAccessPointContractDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbServiceAccessPointContracts []entity.ServiceAccessPointContract
	err := sapc.Where(query, args).Find(&dbServiceAccessPointContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPointContracts))
	for i, v := range dbServiceAccessPointContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sapc *serviceAccessPointContractDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := sapc.Table(sapc.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbServiceAccessPointContracts []entity.ServiceAccessPointContract
	err = sapc.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbServiceAccessPointContracts).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPointContracts))
	for i, v := range dbServiceAccessPointContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (sapc *serviceAccessPointContractDao) List() ([]entity.Entity, error) {
	var dbServiceAccessPointContracts []entity.ServiceAccessPointContract
	err := sapc.Find(&dbServiceAccessPointContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceAccessPointContracts))
	for i, v := range dbServiceAccessPointContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sapc *serviceAccessPointContractDao) Create(e entity.Entity) (entity.Entity, error) {
	err := sapc.Session.Create(e).Error()
	return e, err
}

func (sapc *serviceAccessPointContractDao) UpdateByID(id int, e entity.Entity) (entity.Entity, error) {
	err := sapc.Where(&entity.ServiceAccessPointContract{ID: id}).Updates(e).Error()
	return e, err
}

func (sapc *serviceAccessPointContractDao) DeleteByID(id int) error {
	return sapc.Where(&entity.ServiceAccessPointContract{ID: id}).Delete(&entity.ServiceAccessPointContract{}).Error()
}

func (sapc *serviceAccessPointContractDao) BeginTransaction() bd.Session {
	return sapc.Begin()
}

func (sapc *serviceAccessPointContractDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sapc *serviceAccessPointContractDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sapc *serviceAccessPointContractDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
