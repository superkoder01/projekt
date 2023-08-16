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
