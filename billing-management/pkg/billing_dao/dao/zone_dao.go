package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type zoneDao struct {
	bd.Session
}

func NewZoneDao(s bd.Session) *zoneDao {
	return &zoneDao{s}
}

func (z *zoneDao) NewEntity() entity.Entity {
	return entity.NewZone()
}

func (z *zoneDao) GetByID(i int) (entity.Entity, error) {
	var dbZone entity.Zone
	err := z.Where(entity.Zone{ID: i}).Take(&dbZone).Error()
	return &dbZone, err
}

func (z *zoneDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbZones []entity.Zone
	err := z.Where(query, args).Find(&dbZones).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbZones))
	for i, v := range dbZones {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (z *zoneDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := z.Table(z.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbZones []entity.Zone
	err = z.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbZones).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbZones))
	for i, v := range dbZones {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (z *zoneDao) List() ([]entity.Entity, error) {
	var dbZones []entity.Zone
	err := z.Find(&dbZones).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbZones))
	for i, v := range dbZones {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (z *zoneDao) Create(e entity.Entity) (entity.Entity, error) {
	err := z.Session.Create(e).Error()
	return e, err
}

func (z *zoneDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := z.Where(&entity.Zone{ID: i}).Updates(e).Error()
	return e, err
}

func (z *zoneDao) DeleteByID(i int) error {
	return z.Where(&entity.Zone{ID: i}).Delete(&entity.Zone{}).Error()
}

func (z *zoneDao) BeginTransaction() bd.Session {
	return z.Begin()
}

func (z *zoneDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (z *zoneDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (z *zoneDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
