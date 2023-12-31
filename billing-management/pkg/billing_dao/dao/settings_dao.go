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

type settingsDao struct {
	bd.Session
}

func NewSettingsDao(s bd.Session) *settingsDao {
	return &settingsDao{s}
}

func (z *settingsDao) NewEntity() entity.Entity {
	return entity.NewSettings()
}

func (z *settingsDao) GetByID(i int) (entity.Entity, error) {
	var dbSettings entity.Settings
	err := z.Where(entity.Settings{ID: i}).Take(&dbSettings).Error()
	return &dbSettings, err
}

func (z *settingsDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbSettings []entity.Settings
	err := z.Where(query, args).Find(&dbSettings).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbSettings))
	for i, v := range dbSettings {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (z *settingsDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := z.Table(z.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbSettings []entity.Settings
	err = z.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbSettings).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbSettings))
	for i, v := range dbSettings {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (z *settingsDao) List() ([]entity.Entity, error) {
	var dbSettings []entity.Settings
	err := z.Find(&dbSettings).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbSettings))
	for i, v := range dbSettings {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (z *settingsDao) Create(e entity.Entity) (entity.Entity, error) {
	err := z.Session.Create(e).Error()
	return e, err
}

func (z *settingsDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := z.Where(&entity.Settings{ID: i}).Updates(e).Error()
	return e, err
}

func (z *settingsDao) DeleteByID(i int) error {
	return z.Where(&entity.Settings{ID: i}).Delete(&entity.Settings{}).Error()
}

func (z *settingsDao) BeginTransaction() bd.Session {
	return z.Begin()
}

func (z *settingsDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (z *settingsDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (z *settingsDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
