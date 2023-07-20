package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/distribution_network_operator"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
	"strconv"
)

type distributionNetworkOperatorHandler struct {
	svc service.Service
}

func NewDistributionNetworkOperatorHandler(svc service.Service) *distributionNetworkOperatorHandler {
	return &distributionNetworkOperatorHandler{svc: svc}
}

func (c *distributionNetworkOperatorHandler) List(ctx *gin.Context) {
	models, err := c.svc.List()
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, utils.WrapQueryResult(int64(len(models)), models))
}

func (c *distributionNetworkOperatorHandler) Create(ctx *gin.Context) {
	var newDistributionNetworkOperator distribution_network_operator.DistributionNetworkOperator
	err := ctx.BindJSON(&newDistributionNetworkOperator)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createDistributionNetworkOperator model.Model
	createDistributionNetworkOperator, err = c.svc.Create(&newDistributionNetworkOperator)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createDistributionNetworkOperator)
}

func (c *distributionNetworkOperatorHandler) GetWithFilter(ctx *gin.Context) {
	models, err := c.svc.GetWithFilter(nil) // TODO filters
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, utils.WrapQueryResult(int64(len(models)), models))
}

func (c *distributionNetworkOperatorHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	model, err := c.svc.GetByID(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *distributionNetworkOperatorHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	var updateDistributionNetworkOperator distribution_network_operator.DistributionNetworkOperator
	err = ctx.BindJSON(&updateDistributionNetworkOperator)
	if err != nil {
		e.HandleError(err, ctx)
	}

	model, err := c.svc.UpdateByID(id, &updateDistributionNetworkOperator)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *distributionNetworkOperatorHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	err = c.svc.DeleteByID(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, "OK")
}

func (c *distributionNetworkOperatorHandler) Get(ctx *gin.Context) {
}

func (c *distributionNetworkOperatorHandler) Update(ctx *gin.Context) {
}

func (c *distributionNetworkOperatorHandler) Delete(ctx *gin.Context) {
}
