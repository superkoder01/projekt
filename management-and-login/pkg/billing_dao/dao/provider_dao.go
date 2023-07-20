package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
)

type providerDao struct {
	bd.Session
}

func NewProviderDao(s bd.Session) *providerDao {
	return &providerDao{s}
}

func (p *providerDao) NewEntity() entity.Entity {
	return entity.NewProvider()
}

func (p *providerDao) GetByID(i int) (entity.Entity, error) {
	var dbProvider entity.Provider
	err := p.Where(entity.Provider{ID: i}).Take(&dbProvider).Error()
	return &dbProvider, err
}

func (p *providerDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbProviders []entity.Provider
	err := p.Where(query, args).Find(&dbProviders).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbProviders))
	for i, v := range dbProviders {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (p *providerDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := p.Debug().Table(p.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbProviders []entity.Provider
	err = p.Debug().Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbProviders).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbProviders))
	for i, v := range dbProviders {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (p *providerDao) Check(q *bd.Query) (int, error) {
	var count int64
	err := p.Table(p.NewEntity().TableName()).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (p *providerDao) List() ([]entity.Entity, error) {
	var dbProviders []entity.Provider
	err := p.Find(&dbProviders).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbProviders))
	for i, v := range dbProviders {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (p *providerDao) Create(e entity.Entity) (entity.Entity, error) {
	err := p.Session.Create(e).Error()
	return e, err
}

func (p *providerDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := p.Where(&entity.Provider{ID: i}).Updates(e).Error()
	return e, err
}

func (p *providerDao) DeleteByID(i int) error {
	return p.Where(&entity.Provider{ID: i}).Delete(&entity.Provider{}).Error()
}

func (p *providerDao) BeginTransaction() bd.Session {
	return p.Begin()
}

func (p *providerDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (p *providerDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (p *providerDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
