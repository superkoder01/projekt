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
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/billing_management/contract"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	logger = logging.MustGetLogger("controller")
)

type signerController struct {
	service service.Service
}

func NewReportController(s service.Service) *signerController {
	return &signerController{service: s}
}

func (c *signerController) InitSign(ctx *gin.Context) {
	//// TODO: Handle new parameter - jsonType
	//// TODO: var newAuthorisation - authorisation.Authorisation

	var newContract contract.Contract
	err := ctx.BindJSON(&newContract)

	if err != nil {
		logger.Errorf("Cannot read contract")
		e.HandleError(err, ctx)
		return
	}
	loginHauth, err := c.service.InitSign(ctx, newContract)
	//// TODO: c.service.InitSign(ctx, newAuthorisation)

	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, bson.M{"hauth": loginHauth})
}

func (c *signerController) SigningCompletedNotification(ctx *gin.Context) {
	err := c.service.SigningCompletedNotification(ctx)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, nil)
}
