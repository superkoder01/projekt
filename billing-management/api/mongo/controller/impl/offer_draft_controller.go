/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
