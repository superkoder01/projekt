package mongodb

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type mongoSession struct {
	Database *mongo.Database
	Ctx      context.Context
}

func NewMongoSession(config configuration.MongoConfig) (*mongoSession, error) {
	host := fmt.Sprintf("%s:%s", config.Host, config.Port)
	opts := options.Client().SetHosts([]string{host})
	opts.SetAuth(options.Credential{Username: config.User, Password: config.Password, AuthMechanism: "SCRAM-SHA-1", AuthSource: "admin"})
	//opts.Auth
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.Database)
	return &mongoSession{Database: database}, err
}

func (s *mongoSession) Aggregate(pipeline mongo.Pipeline, collection string) (*mongo.Cursor, error) {
	return s.Database.Collection(collection).Aggregate(s.Ctx, pipeline)
}

func (s *mongoSession) FindOne(query interface{}, collection string) *mongo.SingleResult {
	return s.Database.Collection(collection).FindOne(s.Ctx, query)
}

func (s *mongoSession) Find(query interface{}, collection string) (*mongo.Cursor, error) {
	return s.Database.Collection(collection).Find(s.Ctx, query)
}

func (s *mongoSession) List(query interface{}, collection string, options *options.FindOptions) (*mongo.Cursor, error) {
	return s.Database.Collection(collection).Find(s.Ctx, query, options)
}

func (s *mongoSession) Create(model model.Model, collection string) (*mongo.InsertOneResult, error) {
	return s.Database.Collection(collection).InsertOne(s.Ctx, model)
}

func (s *mongoSession) Update(filter interface{}, model interface{}, collection string, opts *options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return s.Database.Collection(collection).FindOneAndReplace(s.Ctx, filter, model, opts)
}

func (s *mongoSession) Count(filter interface{}, collection string) (int64, error) {
	return s.Database.Collection(collection).CountDocuments(s.Ctx, filter)
}
