package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
	"gorm.io/gorm/clause"
)

type userDao struct {
	bd.Session
}

func NewUserDao(s bd.Session) *userDao {
	return &userDao{s}
}

func (ud *userDao) NewEntity() entity.Entity {
	return entity.NewUser()
}

func (ud *userDao) GetByID(i int) (entity.Entity, error) {
	var dbUser entity.User
	err := ud.Preload(clause.Associations).Where(entity.User{ID: i}).Take(&dbUser).Error()
	return &dbUser, err
}

func (ud *userDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbUsers []entity.User
	err := ud.Preload(clause.Associations).Where(query, args).Find(&dbUsers).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbUsers))
	for i, v := range dbUsers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ud *userDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := ud.Table(ud.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbUsers []entity.User
	err = ud.Preload(clause.Associations).Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbUsers).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbUsers))
	for i, v := range dbUsers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (ud *userDao) List() ([]entity.Entity, error) {
	var dbUsers []entity.User
	err := ud.Preload(clause.Associations).Find(&dbUsers).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbUsers))
	for i, v := range dbUsers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (ud *userDao) Create(e entity.Entity) (entity.Entity, error) {
	err := ud.Session.Create(e).Error()
	return e, err
}

func (ud *userDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := ud.Where(&entity.User{ID: i}).Updates(e).Error()
	return e, err
}

func (ud *userDao) DeleteByID(i int) error {
	return ud.Where(&entity.User{ID: i}).Delete(&entity.User{}).Error()
}

func (ud *userDao) BeginTransaction() bd.Session {
	return ud.Begin()
}

func (ud *userDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ud *userDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (ud *userDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
