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
