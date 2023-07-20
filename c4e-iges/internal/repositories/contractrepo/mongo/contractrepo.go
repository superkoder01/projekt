package mongo

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	contractRepo struct {
		log     logger.Logger
		cfg     *config.AppConfig
		client  *mongodriver.Client
		db      string
		timeout time.Duration
	}
)

func NewContractRepo(log logger.Logger, cfg *config.AppConfig) ports.ContractRepo {
	client, err := mongodriver.NewClient(options.Client().ApplyURI(cfg.Mongo.Uri))

	repo := &contractRepo{
		client:  client,
		db:      cfg.Mongo.DbName,
		timeout: time.Duration(cfg.Mongo.Timeout) * time.Second,
		log:     log,
		cfg:     cfg,
	}

	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		panic(fmt.Errorf("failed to create connection with contract repo: %v", err))
	}
	//defer func(client *mongodriver.Client, ctx context.Context) {
	//	err := client.Disconnect(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//}(client, ctx)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("can not connect to contract repo: %v", err)
	}

	log.Info("successfully connected and pinged")
	return repo
}

func (repo *contractRepo) GetContractByContractNumber(ctx context.Context, contractNumber string) (*billing.Contract, error) {
	repo.log.Infof("fetching contract from repository: %s", contractNumber)

	//ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	//defer cancel()
	contract := &billing.Contract{}

	err := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.ContractCollectionName).FindOne(
		ctx,
		bson.M{
			"payload.contractDetails.number": contractNumber,
		},
	).Decode(&contract)
	if err != nil {
		repo.log.Errorf("contract %s not found in repository, %v", contractNumber, err)
		return nil, err
	}

	return contract, nil
}
