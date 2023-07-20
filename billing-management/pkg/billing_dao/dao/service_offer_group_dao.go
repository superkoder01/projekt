package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type serviceOfferGroupDao struct {
	bd.Session
}

func NewServiceOfferGroupDao(s bd.Session) *serviceOfferGroupDao {
	return &serviceOfferGroupDao{s}
}

func (sog *serviceOfferGroupDao) NewEntity() entity.Entity {
	return entity.NewServiceOfferGroup()
}

func (sog *serviceOfferGroupDao) GetByID(i int) (entity.Entity, error) {
	var dbServiceOfferGroup entity.ServiceOfferGroup
	err := sog.Where(entity.ServiceOfferGroup{ID: i}).Take(&dbServiceOfferGroup).Error()
	return &dbServiceOfferGroup, err
}

func (sog *serviceOfferGroupDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbServiceOfferGroups []entity.ServiceOfferGroup
	err := sog.Where(query, args).Find(&dbServiceOfferGroups).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOfferGroups))
	for i, v := range dbServiceOfferGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sog *serviceOfferGroupDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := sog.Table(sog.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbServiceOfferGroups []entity.ServiceOfferGroup
	err = sog.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbServiceOfferGroups).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOfferGroups))
	for i, v := range dbServiceOfferGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (sog *serviceOfferGroupDao) List() ([]entity.Entity, error) {
	var dbServiceOfferGroups []entity.ServiceOfferGroup
	err := sog.Find(&dbServiceOfferGroups).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbServiceOfferGroups))
	for i, v := range dbServiceOfferGroups {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (sog *serviceOfferGroupDao) Create(e entity.Entity) (entity.Entity, error) {
	err := sog.Session.Create(e).Error()
	return e, err
}

func (sog *serviceOfferGroupDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := sog.Where(&entity.ServiceOfferGroup{ID: i}).Updates(e).Error()
	return e, err
}

func (sog *serviceOfferGroupDao) DeleteByID(i int) error {
	return sog.Where(&entity.ServiceOfferGroup{ID: i}).Delete(&entity.ServiceOfferGroup{}).Error()
}

func (sog *serviceOfferGroupDao) BeginTransaction() bd.Session {
	return sog.Begin()
}

func (sog *serviceOfferGroupDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sog *serviceOfferGroupDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (sog *serviceOfferGroupDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
