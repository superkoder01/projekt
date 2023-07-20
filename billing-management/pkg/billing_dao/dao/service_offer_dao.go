package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type serviceOfferDao struct {
	bd.Session
}

func NewServiceOfferDao(s bd.Session) *serviceOfferDao {
	return &serviceOfferDao{s}
}

func (so *serviceOfferDao) NewEntity() entity.Entity {
	return entity.NewServiceOffer()
}

func (so *serviceOfferDao) GetByID(i int) (entity.Entity, error) {
	var dbServiceOffer entity.ServiceOffer
	err := so.Where(entity.ServiceOffer{ID: i}).Take(&dbServiceOffer).Error()
	return &dbServiceOffer, err
}

func (so *serviceOfferDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbServiceOffers []entity.ServiceOffer
	err := so.Where(query, args).Find(&dbServiceOffers).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOffers))
	for i, v := range dbServiceOffers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (so *serviceOfferDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := so.Table(so.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbServiceOffers []entity.ServiceOffer
	err = so.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbServiceOffers).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOffers))
	for i, v := range dbServiceOffers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (so *serviceOfferDao) List() ([]entity.Entity, error) {
	var dbServiceOffers []entity.ServiceOffer
	err := so.Find(&dbServiceOffers).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOffers))
	for i, v := range dbServiceOffers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (so *serviceOfferDao) Create(e entity.Entity) (entity.Entity, error) {
	err := so.Session.Create(e).Error()
	return e, err
}

func (so *serviceOfferDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := so.Where(&entity.ServiceOffer{ID: i}).Updates(e).Error()
	return e, err
}

func (so *serviceOfferDao) DeleteByID(i int) error {
	return so.Where(&entity.ServiceOffer{ID: i}).Delete(&entity.ServiceOffer{}).Error()
}

func (so *serviceOfferDao) BeginTransaction() bd.Session {
	return so.Begin()
}

func (so *serviceOfferDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (so *serviceOfferDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (so *serviceOfferDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
