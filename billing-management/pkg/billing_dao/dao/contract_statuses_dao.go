package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type contractStatusesDao struct {
	bd.Session
}

func NewContractStatusesDao(s bd.Session) *contractStatusesDao {
	return &contractStatusesDao{s}
}

func (bt *contractStatusesDao) NewEntity() entity.Entity {
	return entity.NewContractStatuses()
}

func (bt *contractStatusesDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbContractStatuses []entity.ContractStatuses
	err := bt.Where(query, args).Find(&dbContractStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *contractStatusesDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := bt.Table(bt.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbContractStatuses []entity.ContractStatuses
	err = bt.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbContractStatuses).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (bt *contractStatusesDao) List() ([]entity.Entity, error) {
	var dbContractStatuses []entity.ContractStatuses
	err := bt.Find(&dbContractStatuses).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbContractStatuses))
	for i, v := range dbContractStatuses {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (bt *contractStatusesDao) Create(e entity.Entity) (entity.Entity, error) {
	err := bt.Session.Create(e).Error()
	return e, err
}

func (bt *contractStatusesDao) DeleteByName(name string) error {
	return bt.Where(&entity.ContractStatuses{Name: name}).Delete(&entity.ContractStatuses{}).Error()
}

func (bt *contractStatusesDao) BeginTransaction() bd.Session {
	return bt.Begin()
}

func (bt *contractStatusesDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (bt *contractStatusesDao) DeleteByID(i int) error {
	return nil
}
