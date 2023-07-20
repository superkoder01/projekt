package collections

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer_draft"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	mongoUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type offerDraftCollection struct {
	session mongodb.MongoSession
}

func (c *offerDraftCollection) List(ctx *gin.Context, query *mongoUtils.Query) (model.Model, error) {
	var result []model.Model
	pipeline := &mongo.Pipeline{}
	mongoUtils.SetProviderPipeline(ctx, query, pipeline)
	countPipeline := *pipeline
	mongoUtils.SetCount(&countPipeline)
	countAggregate, err := c.session.Aggregate(countPipeline, mongoUtils.OFFER_DRAFTS)
	if err != nil {
		return nil, err
	}
	count, err := mongoUtils.GetCountElementsResponse(ctx, countAggregate)
	if err != nil {
		return nil, err
	}
	mongoUtils.SetAggregateOptions(query, pipeline)
	res, err := c.session.Aggregate(*pipeline, mongoUtils.OFFER_DRAFTS)
	//res, err := c.session.List(filters, mongoUtils.OFFER_DRAFTS, mongoUtils.SetFindOptions(query))
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var elem offer_draft.OfferDraft
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return mongoUtils.WrapQueryResult(count, result), nil
}

func NewOfferDraftCollection(s mongodb.MongoSession) *offerDraftCollection {
	return &offerDraftCollection{s}
}

func (c *offerDraftCollection) FindOne(id string) (model.Model, error) {
	var mdOfferDraft offer_draft.OfferDraft
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.session.FindOne(bson.M{"_id": objectId}, mongoUtils.OFFER_DRAFTS).Decode(&mdOfferDraft)
	if err != nil {
		return nil, err
	}
	//return nil, nil
	return &mdOfferDraft, nil
}

func (c *offerDraftCollection) Create(ctx *gin.Context, model model.Model) (string, error) {
	if body, ok := model.(*offer_draft.OfferDraft); ok {
		// TODO: move to body builder
		body.Payload.OfferDetails.CreationDate = mongoUtils.GenerateCreationDate()
		res, err := c.session.Create(model, mongoUtils.OFFER_DRAFTS)
		if err != nil {
			return "", err
		}
		return res.InsertedID.(primitive.ObjectID).Hex(), err
	} else {
		return "", e.ApiErrInvalidDataModel
	}
}

func (c *offerDraftCollection) Update(ctx *gin.Context, id string, model model.Model) (model.Model, error) {
	if body, ok := model.(*offer_draft.OfferDraft); ok {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": objectId}
		res := c.session.Update(filter, body, mongoUtils.OFFER_DRAFTS, nil)
		resBody := &offer_draft.OfferDraft{}
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
