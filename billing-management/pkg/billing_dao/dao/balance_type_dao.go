package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type balanceTypeDao struct {
	bd.Session
}

func NewBalanceTypeDao(s bd.Session) *balanceTypeDao {
	return &balanceTypeDao{s}
}

func (bt *balanceTypeDao) NewEntity() entity.Entity {
	return entity.NewBalanceType()
}

func (bt *balanceTypeDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbBalanceTypes []entity.BalanceType
	err := bt.Where(query, args).Find(&dbBalanceTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbBalanceTypes))
	for i, v := range dbBalanceTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *balanceTypeDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := bt.Table(bt.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbBalanceTypes []entity.BalanceType
	err = bt.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbBalanceTypes).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbBalanceTypes))
	for i, v := range dbBalanceTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (bt *balanceTypeDao) List() ([]entity.Entity, error) {
	var dbBalanceTypes []entity.BalanceType
	err := bt.Find(&dbBalanceTypes).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbBalanceTypes))
	for i, v := range dbBalanceTypes {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *balanceTypeDao) Create(e entity.Entity) (entity.Entity, error) {
	err := bt.Session.Create(e).Error()
	return e, err
}

func (bt *balanceTypeDao) DeleteByName(name string) error {
	return bt.Where(&entity.BalanceType{Name: name}).Delete(&entity.BalanceType{}).Error()
}

func (bt *balanceTypeDao) BeginTransaction() bd.Session {
	return bt.Begin()
}

func (bt *balanceTypeDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *balanceTypeDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *balanceTypeDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *balanceTypeDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (bt *balanceTypeDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (bt *balanceTypeDao) DeleteByID(i int) error {
	return nil
}
