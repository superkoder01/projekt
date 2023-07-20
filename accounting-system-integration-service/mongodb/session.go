package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSession interface {
	List(query interface{}, collection string, options *options.FindOptions) (*mongo.Cursor, error)
	Find(query interface{}, collection string) (*mongo.Cursor, error)
	FindOne(query interface{}, collection string) *mongo.SingleResult
}
