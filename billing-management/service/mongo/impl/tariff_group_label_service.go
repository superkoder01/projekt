package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mongo/collections"
)

type tariffGroupLabelService struct {
	collection collections.Collection
}

func NewTariffGroupLabelService(c collections.Collection) *tariffGroupLabelService {
	return &tariffGroupLabelService{collection: c}
}

func (c *tariffGroupLabelService) List(ctx *gin.Context, query *utils.Query) (model.Model, error) {
	en, err := c.collection.List(ctx, query)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	return en, err
}

func (c *tariffGroupLabelService) FindOne(id string) (model.Model, error) {
	en, err := c.collection.FindOne(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	return en, err
}

func (c *tariffGroupLabelService) Create(ctx *gin.Context, model model.Model) (string, error) {
	en, err := c.collection.Create(ctx, model)
	if err != nil {
		return "", err
	}
	return en, err
}

func (c *tariffGroupLabelService) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	en, err := c.collection.Update(ctx, id, model)
	if err != nil {
		return nil, err
	}
	return en, err
}