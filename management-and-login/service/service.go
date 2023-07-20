package service

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type Service interface {
	Create(model.Model) (model.Model, error)

	//Delete(model.Model) error
	DeleteByID(int) error

	GetByID(int) (model.Model, error)
	GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error)
	List() ([]model.Model, error)

	Query(interface{}, *api_utils.Query) (int, []model.Model, error)

	Check(*api_utils.Query) (int, error)

	UpdateByID(int, model.Model) (model.Model, error)
	//Update(model.Model) (model.Model, error)

	ModelToEntity(model.Model) (entity.Entity, error)
	EntityToModel(entity.Entity) (model.Model, error)
}
