package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/contract"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	utilities "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

var (
	logger = logging.MustGetLogger("controller")
)

type contractController struct {
	service mongo.Service
}

func NewContractController(s mongo.Service) *contractController {
	return &contractController{service: s}
}

func (c *contractController) List(ctx *gin.Context) {
	query := utilities.ParseQuery(ctx.Request.RequestURI)
	results, err := c.service.List(ctx, query)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, results)
}

func (c *contractController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.service.FindOne(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *contractController) Create(ctx *gin.Context) {
	var newContract contract.Contract
	err := ctx.BindJSON(&newContract)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createContract string
	createContract, err = c.service.Create(ctx, &newContract)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createContract)
}

func (c *contractController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var contract contract.Contract
	err := ctx.BindJSON(&contract)
	if err != nil {
		e.HandleError(err, ctx)
	}

	var updateContract model.Model
	updateContract, err = c.service.Update(ctx, id, &contract)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, updateContract)
	return
}
