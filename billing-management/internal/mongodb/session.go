package mongodb

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSession interface {
	Aggregate(pipeline mongo.Pipeline, collection string) (*mongo.Cursor, error)
	List(query interface{}, collection string, options *options.FindOptions) (*mongo.Cursor, error)
	Find(query interface{}, collection string) (*mongo.Cursor, error)
	FindOne(query interface{}, collection string) *mongo.SingleResult
	Create(model model.Model, collection string) (*mongo.InsertOneResult, error)
	Update(filter interface{}, model interface{}, collection string, opts *options.FindOneAndReplaceOptions) *mongo.SingleResult
	Count(filter interface{}, collection string) (int64, error)
}
