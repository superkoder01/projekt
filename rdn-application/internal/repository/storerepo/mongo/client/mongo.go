package client

import (
	"RDN-application/pkg/config"
	"context"
	client "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Client struct {
	client   *client.Client
	dbName   string
	collName string
	timeout  time.Duration
}

func NewMongoClient(cfg config.AppConfig) *Client {
	mongoClient, err := client.NewClient(options.Client().ApplyURI(cfg.GetStoreConfig().Uri))

	repo := &Client{
		client:   mongoClient,
		dbName:   cfg.GetStoreConfig().DbName,
		collName: cfg.GetStoreConfig().CollectionName,
		timeout:  time.Duration(cfg.GetStoreConfig().Timeout) * time.Second,
	}
	ctx, timeout := context.WithTimeout(context.Background(), repo.timeout)
	defer timeout()

	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to create connection with mongoDb ")
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Can't connect to Mongo server %v", err)
	}

	return repo
}

func (mongo *Client) GetConnection() *client.Database {
	return mongo.client.Database(mongo.dbName)
	//	.Collection(mongo.collName)
}
