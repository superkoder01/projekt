package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type customerAccountDao struct {
	bd.Session
}

func NewCustomerAccountDao(s bd.Session) *customerAccountDao {
	return &customerAccountDao{s}
}

func (c *customerAccountDao) NewEntity() entity.Entity {
	return entity.NewCustomerAccount()
}

func (c *customerAccountDao) GetByID(i int) (entity.Entity, error) {
	var dbCustomerAcc entity.CustomerAccount
	err := c.Where(entity.CustomerAccount{ID: i}).Take(&dbCustomerAcc).Error()
	return &dbCustomerAcc, err
}

func (c *customerAccountDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbCustomerAccs []entity.CustomerAccount
	err := c.Where(query, args).Find(&dbCustomerAccs).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbCustomerAccs))
	for i, v := range dbCustomerAccs {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *customerAccountDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := c.Table(c.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbCustomerAccs []entity.CustomerAccount
	err = c.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbCustomerAccs).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbCustomerAccs))
	for i, v := range dbCustomerAccs {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (c *customerAccountDao) List() ([]entity.Entity, error) {
	var dbCustomerAccs []entity.CustomerAccount
	err := c.Find(&dbCustomerAccs).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbCustomerAccs))
	for i, v := range dbCustomerAccs {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *customerAccountDao) Create(e entity.Entity) (entity.Entity, error) {
	err := c.Session.Create(e).Error()
	return e, err
}

func (c *customerAccountDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := c.Where(&entity.CustomerAccount{ID: i}).Updates(e).Error()
	return e, err
}

func (c *customerAccountDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *customerAccountDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *customerAccountDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *customerAccountDao) DeleteByID(i int) error {
	return c.Where(&entity.CustomerAccount{ID: i}).Delete(&entity.CustomerAccount{}).Error()
}

func (c *customerAccountDao) BeginTransaction() bd.Session {
	return c.Begin()
}
