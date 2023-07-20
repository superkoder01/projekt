package mongo

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
)

type Service interface {
	List(*gin.Context, *utils.Query) (model.Model, error)
	FindOne(string) (model.Model, error)
	Create(*gin.Context, model.Model) (string, error)
	Update(*gin.Context, string, model.Model) (model.Model, error)
}
