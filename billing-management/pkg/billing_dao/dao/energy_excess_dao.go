package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type energyExcessDao struct {
	bd.Session
}

func NewEnergyExcessDao(s bd.Session) *energyExcessDao {
	return &energyExcessDao{s}
}

func (ee *energyExcessDao) NewEntity() entity.Entity {
	return entity.NewEnergyExcess()
}

func (ee *energyExcessDao) GetByID(i int) (entity.Entity, error) {
	var dbVatRate entity.EnergyExcess
	err := ee.Where(entity.EnergyExcess{ID: i}).Take(&dbVatRate).Error()
	return &dbVatRate, err
}

func (ee *energyExcessDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbEnergyExcess []entity.EnergyExcess
	err := ee.Where(query, args).Find(&dbEnergyExcess).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyExcessDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ee.Table(ee.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbEnergyExcess []entity.EnergyExcess
	err = ee.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbEnergyExcess).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ee *energyExcessDao) List() ([]entity.Entity, error) {
	var dbEnergyExcess []entity.EnergyExcess
	err := ee.Find(&dbEnergyExcess).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbEnergyExcess))
	for i, v := range dbEnergyExcess {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ee *energyExcessDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ee.Session.Create(e).Error()
	return e, err
}

func (ee *energyExcessDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := ee.Where(&entity.EnergyExcess{ID: i}).Updates(e).Error()
	return e, err
}

func (ee *energyExcessDao) DeleteByID(i int) error {
	return ee.Where(&entity.EnergyExcess{ID: i}).Delete(&entity.EnergyExcess{}).Error()
}

func (ee *energyExcessDao) BeginTransaction() bd.Session {
	return ee.Begin()
}

func (ee *energyExcessDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyExcessDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ee *energyExcessDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
