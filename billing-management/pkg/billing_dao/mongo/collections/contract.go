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
package collections

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/contract"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/validation"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

var (
	logger = logging.MustGetLogger("blockchain_service")
)

type contractCollection struct {
	session mongodb.MongoSession
}

func NewContractCollection(s mongodb.MongoSession) *contractCollection {
	return &contractCollection{s}
}

func (c *contractCollection) List(ctx *gin.Context, query *mongoUtils.Query) (model.Model, error) {
	var result []model.Model
	pipeline := &mongo.Pipeline{}
	mongoUtils.SetCustomerIdPipeline(ctx, query, pipeline, mongoUtils.CONTRACTS)
	countPipeline := *pipeline
	mongoUtils.SetCount(&countPipeline)
	countAggregate, err := c.session.Aggregate(countPipeline, mongoUtils.CONTRACTS)
	if err != nil {
		return nil, err
	}
	count, err := mongoUtils.GetCountElementsResponse(ctx, countAggregate)
	if err != nil {
		return nil, err
	}
	mongoUtils.SetAggregateOptions(query, pipeline)
	res, err := c.session.Aggregate(*pipeline, mongoUtils.CONTRACTS)
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem contract.Contract
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(count, result), nil
}

func (c *contractCollection) FindOne(id string) (model.Model, error) {
	var mdContract contract.Contract
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.CONTRACTS).Decode(&mdContract)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdContract, nil
}

func (c *contractCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	if body, ok := model.(*contract.Contract); ok {
		filters := mongoUtils.SetProviderFilters(ctx)
		count, err := c.session.Count(filters, mongoUtils.CONTRACTS)
		// TODO: move to body builder
		body.Payload.ContractDetails.Number = mongoUtils.GenerateDocumentNumber(body.Header.Provider, mongoUtils.CONTRACTS, count)
		body.Payload.ContractDetails.CreationDate = mongoUtils.GenerateCreationDate()
		res, err := c.session.Create(body, mongoUtils.CONTRACTS)
		if err != nil {
			return "", err
		}
		return res.InsertedID.(primitive.ObjectID).Hex(), err
	} else {
		return "", e.ApiErrInvalidDataModel
	}
}

func (c *contractCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	if body, ok := model.(*contract.Contract); ok {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": objectId}
		oldContractRes := c.session.FindOne(filter, mongoUtils.CONTRACTS)
		oldContract := &contract.Contract{}
		err = oldContractRes.Decode(oldContract)
		if !validation.ValidateIfCanUpdateContract(oldContract) {
			e.HandleError(e.CannotUpdateWhenStateIsAccepted, ctx)
			return nil, e.CannotUpdateWhenStateIsAccepted
		}
		checkIfIsGettingAccepted(ctx, oldContract, body)
		res := c.session.Update(filter, body, mongoUtils.CONTRACTS, nil)
		resBody := &contract.Contract{}
		err = res.Decode(resBody)
		if err != nil {
			return nil, err
		}
		return resBody, res.Err()
	} else {
		e.HandleError(e.ApiErrInvalidDataModel, ctx)
		return nil, e.ApiErrInvalidDataModel
	}
}

func checkIfIsGettingAccepted(ctx *gin.Context, oldContract, newContract *contract.Contract) {
	if strings.EqualFold(newContract.Payload.ContractDetails.State, enum.CS_ACCEPTED.Name()) &&
		!strings.EqualFold(oldContract.Payload.ContractDetails.State, enum.CS_ACCEPTED.Name()) &&
		(newContract.Header.Provider == "Keno Energia Sp. z o.o." || newContract.Header.Provider == "Ovoo") {
		newContract.Payload.SellerDetails.BankAccountNumber = mongoUtils.GenerateBankNumber(ctx, newContract.Payload.CustomerDetails.RegistrationNumber)
	}
}
