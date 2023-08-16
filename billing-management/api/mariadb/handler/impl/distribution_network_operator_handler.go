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
