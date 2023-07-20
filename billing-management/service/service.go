package service

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
)

type Service interface {
	Create(model.Model) (model.Model, error)

	DeleteByID(int) error
	Delete(interface{}) (model.Model, error)

	GetByID(int) (model.Model, error)
	Get(interface{}) (model.Model, error)
	GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error)
	List() ([]model.Model, error)

	UpdateByID(int, model.Model) (model.Model, error)
	Update(interface{}) (model.Model, error)

	ModelToEntity(model.Model) (entity.Entity, error)
	EntityToModel(entity.Entity) (model.Model, error)
}
