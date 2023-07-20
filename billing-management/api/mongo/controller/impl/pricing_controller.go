package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/pricing"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	utilities "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

type pricingController struct {
	service mongo.Service
}

func NewPricingController(s mongo.Service) *pricingController {
	return &pricingController{service: s}
}

func (c *pricingController) List(ctx *gin.Context) {
	query := utilities.ParseQuery(ctx.Request.RequestURI)
	results, err := c.service.List(ctx, query)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, results)
}

func (c *pricingController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.service.FindOne(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *pricingController) Create(ctx *gin.Context) {
	var newPricing pricing.Pricing
	err := ctx.BindJSON(&newPricing)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createPricing string
	createPricing, err = c.service.Create(ctx, &newPricing)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createPricing)
}

func (c *pricingController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pricing pricing.Pricing
	err := ctx.BindJSON(&pricing)
	if err != nil {
		e.HandleError(err, ctx)
	}
	var updatePricing model.Model
	updatePricing, err = c.service.Update(ctx, id, &pricing)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, updatePricing)
}
