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
