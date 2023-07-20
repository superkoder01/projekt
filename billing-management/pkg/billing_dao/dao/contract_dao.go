package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type contractDao struct {
	bd.Session
}

func NewContractDao(s bd.Session) *contractDao {
	return &contractDao{s}
}

func (c *contractDao) NewEntity() entity.Entity {
	return entity.NewContract()
}

func (c *contractDao) GetByID(i int) (entity.Entity, error) {
	var dbContract entity.Contract
	err := c.Where(entity.Contract{ID: i}).Take(&dbContract).Error()
	return &dbContract, err
}

func (c *contractDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbContracts []entity.Contract
	err := c.Where(query, args).Find(&dbContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *contractDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := c.Table(c.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbContracts []entity.Contract
	err = c.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbContracts).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (c *contractDao) List() ([]entity.Entity, error) {
	var dbContracts []entity.Contract
	err := c.Find(&dbContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContracts))
	for i, v := range dbContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (c *contractDao) Create(e entity.Entity) (entity.Entity, error) {
	err := c.Session.Create(e).Error()
	return e, err
}

func (c *contractDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := c.Where(&entity.Contract{ID: i}).Updates(e).Error()
	return e, err
}

func (c *contractDao) DeleteByID(i int) error {
	return c.Where(&entity.Contract{ID: i}).Delete(&entity.Contract{}).Error()
}

func (c *contractDao) BeginTransaction() bd.Session {
	return c.Begin()
}

func (c *contractDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *contractDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (c *contractDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
