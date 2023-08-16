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
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
)

type parameterNameHandler struct {
	svc service.Service
}

func NewParameterNameHandler(svc service.Service) *parameterNameHandler {
	return &parameterNameHandler{svc: svc}
}

func (c *parameterNameHandler) List(ctx *gin.Context) {
	models, err := c.svc.List()
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, models)
}

func (c *parameterNameHandler) Create(ctx *gin.Context) {
}

func (c *parameterNameHandler) GetWithFilter(ctx *gin.Context) {
}

func (c *parameterNameHandler) GetByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) UpdateByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) DeleteByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) Get(ctx *gin.Context) {
}

func (c *parameterNameHandler) Update(ctx *gin.Context) {
}

func (c *parameterNameHandler) Delete(ctx *gin.Context) {
}
