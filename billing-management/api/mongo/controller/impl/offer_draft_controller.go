package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer_draft"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	utilities "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

type offerDraftController struct {
	service mongo.Service
}

func NewOfferDraftController(s mongo.Service) *offerDraftController {
	return &offerDraftController{service: s}
}

func (c *offerDraftController) List(ctx *gin.Context) {
	query := utilities.ParseQuery(ctx.Request.RequestURI)
	results, err := c.service.List(ctx, query)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, results)
}

func (c *offerDraftController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.service.FindOne(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, model)
}

func (c *offerDraftController) Create(ctx *gin.Context) {
	var newOfferDraft offer_draft.OfferDraft
	err := ctx.BindJSON(&newOfferDraft)

	if err != nil {
		e.HandleError(err, ctx)
	}

	var createOfferDraft string
	createOfferDraft, err = c.service.Create(ctx, &newOfferDraft)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, createOfferDraft)
}

func (c *offerDraftController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var offerDraft offer_draft.OfferDraft
	err := ctx.BindJSON(&offerDraft)
	if err != nil {
		e.HandleError(err, ctx)
	}
	var updateDraft model.Model
	updateDraft, err = c.service.Update(ctx, id, &offerDraft)
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, updateDraft)
}
