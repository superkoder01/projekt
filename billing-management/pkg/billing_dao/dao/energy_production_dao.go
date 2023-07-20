package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type energyProductionDao struct {
	bd.Session
}

func NewEnergyProductionDao(s bd.Session) *energyProductionDao {
	return &energyProductionDao{s}
}

func (ee *energyProductionDao) NewEntity() entity.Entity {
	return entity.NewEnergyProduction()
}

func (ee *energyProductionDao) GetByID(i int) (entity.Entity, error) {
	var dbVatRate entity.EnergyProduction
	err := ee.Where(entity.EnergyProduction{ID: i}).Take(&dbVatRate).Error()
	return &dbVatRate, err
}

func (ee *energyProductionDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbEnergyProduction []entity.EnergyProduction
	err := ee.Where(query, args).Find(&dbEnergyProduction).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyProduction))
	for i, v := range dbEnergyProduction {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyProductionDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ee.Table(ee.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbEnergyProduction []entity.EnergyProduction
	err = ee.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbEnergyProduction).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyProduction))
	for i, v := range dbEnergyProduction {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ee *energyProductionDao) List() ([]entity.Entity, error) {
	var dbEnergyProduction []entity.EnergyProduction
	err := ee.Find(&dbEnergyProduction).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyProduction))
	for i, v := range dbEnergyProduction {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyProductionDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ee.Session.Create(e).Error()
	return e, err
}

func (ee *energyProductionDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := ee.Where(&entity.EnergyProduction{ID: i}).Updates(e).Error()
	return e, err
}

func (ee *energyProductionDao) DeleteByID(i int) error {
	return ee.Where(&entity.EnergyProduction{ID: i}).Delete(&entity.EnergyProduction{}).Error()
}

func (ee *energyProductionDao) BeginTransaction() bd.Session {
	return ee.Begin()
}

func (ee *energyProductionDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyProductionDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyProductionDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
