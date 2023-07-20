package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type Dao interface {
	NewEntity() entity.Entity

	GetByID(int) (entity.Entity, error)
	Get(interface{}) (entity.Entity, error)
	GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error)
	List() ([]entity.Entity, error)

	Create(entity.Entity) (entity.Entity, error)

	UpdateByID(int, entity.Entity) (entity.Entity, error)
	Update(interface{}) (entity.Entity, error)

	DeleteByID(int) error
	Delete(interface{}) (entity.Entity, error)

	Query(interface{}, *mysql.Query) (int, []entity.Entity, error)

	BeginTransaction() mysql.Session
}
