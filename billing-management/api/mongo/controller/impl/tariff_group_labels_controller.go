package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/tariff_group_label"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	utilities "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

type tariffGroupLabelController struct {
	service mongo.Service
}

func NewTariffGroupLabelController(s mongo.Service) *tariffGroupLabelController {
	return &tariffGroupLabelController{service: s}
}

func (c *tariffGroupLabelController) List(ctx *gin.Context) {
	query := utilities.ParseQuery(ctx.Request.RequestURI)
	results, err := c.service.List(ctx, query)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, results)
}

func (c *tariffGroupLabelController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.service.FindOne(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *tariffGroupLabelController) Create(ctx *gin.Context) {
	var newTariffGroupLabel tariff_group_label.TariffGroupLabel
	err := ctx.BindJSON(&newTariffGroupLabel)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createTariffGroupLabel string
	createTariffGroupLabel, err = c.service.Create(ctx, &newTariffGroupLabel)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createTariffGroupLabel)
}

func (c *tariffGroupLabelController) Update(ctx *gin.Context) {
	return
}
