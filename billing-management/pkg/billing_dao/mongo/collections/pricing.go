package collections

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/pricing"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type pricingCollection struct {
	session mongodb.MongoSession
}

func NewPricingCollection(s mongodb.MongoSession) *pricingCollection {
	return &pricingCollection{s}
}

func (c *pricingCollection) List(ctx *gin.Context, query *mongoUtils.Query) (model.Model, error) {
	var result []model.Model
	pipeline := &mongo.Pipeline{}
	mongoUtils.SetProviderPipeline(ctx, query, pipeline)
	countPipeline := *pipeline
	mongoUtils.SetCount(&countPipeline)
	countAggregate, err := c.session.Aggregate(countPipeline, mongoUtils.PRICINGS)
	if err != nil {
		return nil, err
	}
	count, err := mongoUtils.GetCountElementsResponse(ctx, countAggregate)
	if err != nil {
		return nil, err
	}
	mongoUtils.SetAggregateOptions(query, pipeline)
	res, err := c.session.Aggregate(*pipeline, mongoUtils.PRICINGS)
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem pricing.Pricing
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(count, result), nil
}

func (c *pricingCollection) FindOne(id string) (model.Model, error) {
	var mdPricing pricing.Pricing
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.PRICINGS).Decode(&mdPricing)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdPricing, nil
}

func (c *pricingCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	res, err := c.session.Create(model, mongoUtils.PRICINGS)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (c *pricingCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	if body, ok := model.(*pricing.Pricing); ok {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": objectId}
		res := c.session.Update(filter, body, mongoUtils.PRICINGS, nil)
		resBody := &pricing.Pricing{}
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
