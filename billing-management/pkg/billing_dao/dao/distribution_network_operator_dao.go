package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type distributionNetworkOperatorDao struct {
	bd.Session
}

func NewDistributionNetworkOperatorDao(s bd.Session) *distributionNetworkOperatorDao {
	return &distributionNetworkOperatorDao{s}
}

func (dso *distributionNetworkOperatorDao) NewEntity() entity.Entity {
	return entity.NewDistributionNetworkOperator()
}

func (dso *distributionNetworkOperatorDao) GetByID(i int) (entity.Entity, error) {
	var dbDistributionNetworkOperator entity.DistributionNetworkOperator
	err := dso.Where(entity.DistributionNetworkOperator{ID: i}).Take(&dbDistributionNetworkOperator).Error()
	return &dbDistributionNetworkOperator, err
}

func (dso *distributionNetworkOperatorDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbDistributionNetworkOperators []entity.DistributionNetworkOperator
	err := dso.Where(query, args).Find(&dbDistributionNetworkOperators).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbDistributionNetworkOperators))
	for i, v := range dbDistributionNetworkOperators {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (dso *distributionNetworkOperatorDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := dso.Table(dso.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbDistributionNetworkOperators []entity.DistributionNetworkOperator
	err = dso.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbDistributionNetworkOperators).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbDistributionNetworkOperators))
	for i, v := range dbDistributionNetworkOperators {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (dso *distributionNetworkOperatorDao) List() ([]entity.Entity, error) {
	var dbDistributionNetworkOperators []entity.DistributionNetworkOperator
	err := dso.Find(&dbDistributionNetworkOperators).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbDistributionNetworkOperators))
	for i, v := range dbDistributionNetworkOperators {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (dso *distributionNetworkOperatorDao) Create(e entity.Entity) (entity.Entity, error) {
	err := dso.Session.Create(e).Error()
	return e, err
}

func (dso *distributionNetworkOperatorDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := dso.Where(&entity.DistributionNetworkOperator{ID: i}).Updates(e).Error()
	return e, err
}

func (dso *distributionNetworkOperatorDao) DeleteByID(i int) error {
	return dso.Where(&entity.DistributionNetworkOperator{ID: i}).Delete(&entity.DistributionNetworkOperator{}).Error()
}

func (dso *distributionNetworkOperatorDao) BeginTransaction() bd.Session {
	return dso.Begin()
}

func (dso *distributionNetworkOperatorDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (dso *distributionNetworkOperatorDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (dso *distributionNetworkOperatorDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
