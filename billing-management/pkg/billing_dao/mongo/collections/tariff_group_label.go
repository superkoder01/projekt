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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/tariff_group_label"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type tariffGroupLabelCollection struct {
	session mongodb.MongoSession
}

func NewTariffGroupLabelCollection(s mongodb.MongoSession) *tariffGroupLabelCollection {
	return &tariffGroupLabelCollection{s}
}

func (c *tariffGroupLabelCollection) List(ctx *gin.Context, query *mongoUtils.Query) (model.Model, error) {
	var result []model.Model
	//filters := mongoUtils.SetCustomerIdPipeline(ctx, query, mongoUtils.TARIFF_GROUP_LABEL)
	res, err := c.session.List(bson.D{}, mongoUtils.TARIFF_GROUP_LABEL, mongoUtils.SetFindOptions(query))
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem tariff_group_label.TariffGroupLabel
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(int64(len(result)), result), nil
}

func (c *tariffGroupLabelCollection) FindOne(id string) (model.Model, error) {
	var mdTariffGroupLabel tariff_group_label.TariffGroupLabel
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.TARIFF_GROUP_LABEL).Decode(&mdTariffGroupLabel)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdTariffGroupLabel, nil
}

func (c *tariffGroupLabelCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	res, err := c.session.Create(model, mongoUtils.TARIFF_GROUP_LABEL)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (c *tariffGroupLabelCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	return nil, nil
}
