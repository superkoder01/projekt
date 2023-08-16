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

type ratingPlanTypeDao struct {
	bd.Session
}

func NewRatingPlanTypeDao(s bd.Session) *ratingPlanTypeDao {
	return &ratingPlanTypeDao{s}
}

func (rpt *ratingPlanTypeDao) NewEntity() entity.Entity {
	return entity.NewRatingPlanType()
}

func (rpt *ratingPlanTypeDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbRatingPlanTypes []entity.RatingPlanType
	err := rpt.Where(query, args).Find(&dbRatingPlanTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlanTypes))
	for i, v := range dbRatingPlanTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (rpt *ratingPlanTypeDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := rpt.Table(rpt.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbRatingPlanTypes []entity.RatingPlanType
	err = rpt.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbRatingPlanTypes).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlanTypes))
	for i, v := range dbRatingPlanTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (rpt *ratingPlanTypeDao) List() ([]entity.Entity, error) {
	var dbRatingPlanTypes []entity.RatingPlanType
	err := rpt.Find(&dbRatingPlanTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlanTypes))
	for i, v := range dbRatingPlanTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (rpt *ratingPlanTypeDao) Create(e entity.Entity) (entity.Entity, error) {
	err := rpt.Session.Create(e).Error()
	return e, err
}

func (rpt *ratingPlanTypeDao) DeleteByName(name entity.RatingPlanTypeName) error {
	return rpt.Where(&entity.RatingPlanType{Name: name}).Delete(&entity.RatingPlanType{}).Error()
}

func (rpt *ratingPlanTypeDao) BeginTransaction() bd.Session {
	return rpt.Begin()
}

func (rpt *ratingPlanTypeDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (rpt *ratingPlanTypeDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (rpt *ratingPlanTypeDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (rpt *ratingPlanTypeDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (rpt *ratingPlanTypeDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (rpt *ratingPlanTypeDao) DeleteByID(i int) error {
	return nil
}
