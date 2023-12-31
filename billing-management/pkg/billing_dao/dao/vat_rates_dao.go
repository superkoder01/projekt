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

type vatRatesDao struct {
	bd.Session
}

func NewVatRateDao(s bd.Session) *vatRatesDao {
	return &vatRatesDao{s}
}

func (vr *vatRatesDao) NewEntity() entity.Entity {
	return entity.NewVatRates()
}

func (vr *vatRatesDao) GetByID(i int) (entity.Entity, error) {
	var dbVatRate entity.VatRates
	err := vr.Where(entity.VatRates{ID: i}).Take(&dbVatRate).Error()
	return &dbVatRate, err
}

func (vr *vatRatesDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbVatRates []entity.VatRates
	err := vr.Where(query, args).Find(&dbVatRates).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbVatRates))
	for i, v := range dbVatRates {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (vr *vatRatesDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := vr.Table(vr.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbVatRates []entity.VatRates
	err = vr.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbVatRates).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbVatRates))
	for i, v := range dbVatRates {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (vr *vatRatesDao) List() ([]entity.Entity, error) {
	var dbVatRates []entity.VatRates
	err := vr.Find(&dbVatRates).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbVatRates))
	for i, v := range dbVatRates {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (vr *vatRatesDao) Create(e entity.Entity) (entity.Entity, error) {
	err := vr.Session.Create(e).Error()
	return e, err
}

func (vr *vatRatesDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := vr.Where(&entity.VatRates{ID: i}).Updates(e).Error()
	return e, err
}

func (vr *vatRatesDao) DeleteByID(i int) error {
	return vr.Where(&entity.VatRates{ID: i}).Delete(&entity.VatRates{}).Error()
}

func (vr *vatRatesDao) BeginTransaction() bd.Session {
	return vr.Begin()
}

func (vr *vatRatesDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (vr *vatRatesDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (vr *vatRatesDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
