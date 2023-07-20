package collections

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type offerCollection struct {
	session mongodb.MongoSession
}

func (c *offerCollection) List(ctx *gin.Context, query *mongoUtils.Query) (model.Model, error) {
	var result []model.Model
	pipeline := &mongo.Pipeline{}
	mongoUtils.SetCustomerIdPipeline(ctx, query, pipeline, mongoUtils.OFFER)
	countPipeline := *pipeline
	mongoUtils.SetCount(&countPipeline)
	countAggregate, err := c.session.Aggregate(countPipeline, mongoUtils.OFFER)
	if err != nil {
		return nil, err
	}
	count, err := mongoUtils.GetCountElementsResponse(ctx, countAggregate)
	if err != nil {
		return nil, err
	}
	mongoUtils.SetAggregateOptions(query, pipeline)
	res, err := c.session.Aggregate(*pipeline, mongoUtils.OFFER)
	//res, err := c.session.List(filters, mongoUtils.OFFER, mongoUtils.SetFindOptions(query))
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem offer.Offer
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(count, result), nil
}

func NewOfferCollection(s mongodb.MongoSession) *offerCollection {
	return &offerCollection{s}
}

func (c *offerCollection) FindOne(id string) (model.Model, error) {
	var mdOffer offer.Offer
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.OFFER).Decode(&mdOffer)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdOffer, nil
}

func (c *offerCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	if body, ok := model.(*offer.Offer); ok {
		filters := mongoUtils.SetProviderFilters(ctx)
		count, err := c.session.Count(filters, mongoUtils.OFFER)
		// TODO: move to body builder
		body.Payload.OfferDetails.Number = mongoUtils.GenerateDocumentNumber(body.Header.Provider, mongoUtils.OFFER, count)
		body.Payload.OfferDetails.CreationDate = mongoUtils.GenerateCreationDate()
		res, err := c.session.Create(body, mongoUtils.OFFER)
		if err != nil {
			return "", err
		}
		return res.InsertedID.(primitive.ObjectID).Hex(), err
	} else {
		return "", e.ApiErrInvalidDataModel
	}
}

func (c *offerCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	if body, ok := model.(*offer.Offer); ok {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": objectId}
		res := c.session.Update(filter, body, mongoUtils.OFFER, nil)
		resBody := &offer.Offer{}
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
