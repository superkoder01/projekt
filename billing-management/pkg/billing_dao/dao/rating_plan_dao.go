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

type ratingPlanDao struct {
	bd.Session
}

func NewRatingPlanDao(s bd.Session) *ratingPlanDao {
	return &ratingPlanDao{s}
}

func (rp *ratingPlanDao) NewEntity() entity.Entity {
	return entity.NewRatingPlan()
}

func (rp *ratingPlanDao) GetByID(i int) (entity.Entity, error) {
	var dbRatingPlan entity.RatingPlan
	err := rp.Where(entity.RatingPlan{ID: i}).Take(&dbRatingPlan).Error()
	return &dbRatingPlan, err
}

func (rp *ratingPlanDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbRatingPlans []entity.RatingPlan
	err := rp.Where(query, args).Find(&dbRatingPlans).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlans))
	for i, v := range dbRatingPlans {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (rp *ratingPlanDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := rp.Table(rp.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbRatingPlans []entity.RatingPlan
	err = rp.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbRatingPlans).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlans))
	for i, v := range dbRatingPlans {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (rp *ratingPlanDao) List() ([]entity.Entity, error) {
	var dbRatingPlans []entity.RatingPlan
	err := rp.Find(&dbRatingPlans).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbRatingPlans))
	for i, v := range dbRatingPlans {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (rp *ratingPlanDao) Create(e entity.Entity) (entity.Entity, error) {
	err := rp.Session.Create(e).Error()
	return e, err
}

func (rp *ratingPlanDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := rp.Where(&entity.RatingPlan{ID: i}).Updates(e).Error()
	return e, err
}

func (rp *ratingPlanDao) DeleteByID(i int) error {
	return rp.Where(&entity.RatingPlan{ID: i}).Delete(&entity.RatingPlan{}).Error()
}

func (rp *ratingPlanDao) BeginTransaction() bd.Session {
	return rp.Begin()
}

func (rp *ratingPlanDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (rp *ratingPlanDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (rp *ratingPlanDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
