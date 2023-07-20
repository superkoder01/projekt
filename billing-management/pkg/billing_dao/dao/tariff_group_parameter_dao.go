package dao

import (
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
	"gorm.io/gorm/clause"
)

type tariffGroupParameterDao struct {
	bd.Session
}

func NewTariffGroupParameterDao(s bd.Session) *tariffGroupParameterDao {
	return &tariffGroupParameterDao{s}
}

func (tgp *tariffGroupParameterDao) NewEntity() entity.Entity {
	return entity.NewTariffGroupParameter()
}

func (tgp *tariffGroupParameterDao) GetByID(i int) (entity.Entity, error) {
	var dbTariffGroupParameter entity.TariffGroupParameter
	err := tgp.Preload(clause.Associations).Where(entity.TariffGroupParameter{ID: i}).Take(&dbTariffGroupParameter).Error()
	return &dbTariffGroupParameter, err
}

func (tgp *tariffGroupParameterDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbTariffGroupParameters []entity.TariffGroupParameter
	err := tgp.Preload(clause.Associations).Where(query, args).Find(&dbTariffGroupParameters).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupParameters))
	for i, v := range dbTariffGroupParameters {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (tgp *tariffGroupParameterDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := tgp.Table(tgp.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbTariffGroupParameters []entity.TariffGroupParameter
	err = tgp.Preload(clause.Associations).Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbTariffGroupParameters).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupParameters))
	for i, v := range dbTariffGroupParameters {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (tgp *tariffGroupParameterDao) List() ([]entity.Entity, error) {
	var dbTariffGroupParameters []entity.TariffGroupParameter
	err := tgp.Preload(clause.Associations).Find(&dbTariffGroupParameters).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbTariffGroupParameters))
	for i, v := range dbTariffGroupParameters {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (tgp *tariffGroupParameterDao) Create(e entity.Entity) (entity.Entity, error) {
	err := tgp.Session.Create(e).Error()
	return e, err
}

func (tgp *tariffGroupParameterDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := tgp.Where(&entity.TariffGroupParameter{ID: i}).Updates(e).Error()
	return e, err
}

func (tgp *tariffGroupParameterDao) DeleteByID(i int) error {
	return tgp.Where(&entity.TariffGroupParameter{ID: i}).Delete(&entity.TariffGroupParameter{}).Error()
}

func (tgp *tariffGroupParameterDao) BeginTransaction() bd.Session {
	return tgp.Begin()
}

func (tgp *tariffGroupParameterDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (tgp *tariffGroupParameterDao) Update(i interface{}) (entity.Entity, error) {
	tgpModel, ok := i.(*entity.TariffGroupParameter)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}
	err := tgp.Where(&entity.TariffGroupParameter{TariffGroupID: tgpModel.TariffGroupID, ParameterNameID: tgpModel.ParameterNameID}).Updates(tgpModel).Error()
	return tgpModel, err
}

func (tgp *tariffGroupParameterDao) Delete(i interface{}) (entity.Entity, error) {
	tgpModel, ok := i.(*entity.TariffGroupParameter)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}
	err := tgp.Where(&entity.TariffGroupParameter{TariffGroupID: tgpModel.TariffGroupID, ParameterNameID: tgpModel.ParameterNameID}).Delete(&entity.TariffGroupParameter{}).Error()
	return tgpModel, err
}
