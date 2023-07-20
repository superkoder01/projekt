package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	utilities "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

type offerController struct {
	service mongo.Service
}

func (c *offerController) List(ctx *gin.Context) {
	query := utilities.ParseQuery(ctx.Request.RequestURI)
	results, err := c.service.List(ctx, query)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, results)
}

func NewOfferController(s mongo.Service) *offerController {
	return &offerController{service: s}
}

func (c *offerController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.service.FindOne(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *offerController) Create(ctx *gin.Context) {
	var newOffer offer.Offer
	err := ctx.BindJSON(&newOffer)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createOffer string
	createOffer, err = c.service.Create(ctx, &newOffer)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createOffer)
}

func (c *offerController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var offer offer.Offer
	err := ctx.BindJSON(&offer)
	if err != nil {
		e.HandleError(err, ctx)
	}
	var updateOffer model.Model
	updateOffer, err = c.service.Update(ctx, id, &offer)
	if err != nil {
		e.HandleError(err, ctx)
	}
	ctx.JSON(200, updateOffer)
	return
}
