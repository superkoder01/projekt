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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type invoiceCollection struct {
	session mongodb.MongoSession
}

func (c *invoiceCollection) List(ctx *gin.Context, query *utils.Query) (model.Model, error) {
	var result []model.Model
	pipeline := &mongo.Pipeline{}
	mongoUtils.SetCustomerIdPipeline(ctx, query, pipeline, mongoUtils.INVOICES)
	countPipeline := *pipeline
	mongoUtils.SetCount(&countPipeline)
	countAggregate, err := c.session.Aggregate(countPipeline, mongoUtils.INVOICES)
	if err != nil {
		return nil, err
	}
	count, err := mongoUtils.GetCountElementsResponse(ctx, countAggregate)
	if err != nil {
		return nil, err
	}
	mongoUtils.SetAggregateOptions(query, pipeline)
	res, err := c.session.Aggregate(*pipeline, mongoUtils.INVOICES)
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem invoice.Invoice
		err := res.Decode(&elem)
		//TODO: strategy
		err = mongoUtils.SetInvoicePayloadModel(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(count, result), nil
}

func NewInvoiceCollection(s mongodb.MongoSession) *invoiceCollection {
	return &invoiceCollection{s}
}

func (c *invoiceCollection) FindOne(id string) (model.Model, error) {
	var mdInvoice invoice.Invoice
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.INVOICES).Decode(&mdInvoice)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdInvoice, nil
}

func (c *invoiceCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	return "", nil
}

func (c *invoiceCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	return nil, nil
}
