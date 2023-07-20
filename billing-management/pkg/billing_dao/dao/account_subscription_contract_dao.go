package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type accountSubscriptionContractDao struct {
	bd.Session
}

func NewAccountSubscriptionContractDao(s bd.Session) *accountSubscriptionContractDao {
	return &accountSubscriptionContractDao{s}
}

func (asc *accountSubscriptionContractDao) NewEntity() entity.Entity {
	return entity.NewAccountSubscriptionContract()
}

func (asc *accountSubscriptionContractDao) GetByContractID(contractID int) (entity.Entity, error) {
	var dbAccountSubscriptionContract entity.AccountSubscriptionContract
	err := asc.Where(entity.AccountSubscriptionContract{ContractID: contractID}).Take(&dbAccountSubscriptionContract).Error()
	return &dbAccountSubscriptionContract, err
}

func (asc *accountSubscriptionContractDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := asc.Table(asc.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbAccountSubscriptionContracts []entity.AccountSubscriptionContract
	err = asc.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbAccountSubscriptionContracts).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptionContracts))
	for i, v := range dbAccountSubscriptionContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (asc *accountSubscriptionContractDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbAccountSubscriptionContracts []entity.AccountSubscriptionContract
	err := asc.Where(query, args).Find(&dbAccountSubscriptionContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptionContracts))
	for i, v := range dbAccountSubscriptionContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (asc *accountSubscriptionContractDao) List() ([]entity.Entity, error) {
	var dbAccountSubscriptionContracts []entity.AccountSubscriptionContract
	err := asc.Find(&dbAccountSubscriptionContracts).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbAccountSubscriptionContracts))
	for i, v := range dbAccountSubscriptionContracts {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (asc *accountSubscriptionContractDao) Create(e entity.Entity) (entity.Entity, error) {
	err := asc.Session.Create(e).Error()
	return e, err
}

func (asc *accountSubscriptionContractDao) UpdateByContractID(contractID int, e entity.Entity) (entity.Entity, error) {
	err := asc.Where(&entity.AccountSubscriptionContract{ContractID: contractID}).Updates(e).Error()
	return e, err
}

func (asc *accountSubscriptionContractDao) DeleteByContractID(contractID int) error {
	return asc.Where(&entity.AccountSubscriptionContract{ContractID: contractID}).Delete(&entity.AccountSubscriptionContract{}).Error()
}

func (asc *accountSubscriptionContractDao) BeginTransaction() bd.Session {
	return asc.Begin()
}

func (asc *accountSubscriptionContractDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (asc *accountSubscriptionContractDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (asc *accountSubscriptionContractDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (asc *accountSubscriptionContractDao) GetByID(i int) (entity.Entity, error) {
	return nil, nil
}

func (asc *accountSubscriptionContractDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	return nil, nil
}

func (asc *accountSubscriptionContractDao) DeleteByID(i int) error {
	return nil
}
