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
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type (
	invoiceRepo struct {
		log     logger.Logger
		cfg     *config.AppConfig
		client  *mongodriver.Client
		db      string
		timeout time.Duration
	}
)

func NewInvoiceRepo(log logger.Logger, cfg *config.AppConfig) ports.InvoiceRepo {
	client, err := mongodriver.NewClient(options.Client().ApplyURI(cfg.Mongo.Uri))

	repo := &invoiceRepo{
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

func (repo *invoiceRepo) GetInvoiceSummaryByIssueDate(ctx context.Context, issueStartDate time.Time, issueEndDate time.Time) (*billing.InvoiceSummary, error) {
	repo.log.Debugf("fetching invoice summary from repository: %s -> %s", issueStartDate, issueEndDate)

	//ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	//defer cancel()
	invoice := &billing.InvoiceSummary{}

	err := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName).FindOne(
		ctx,
		bson.M{
			"payload.invoiceDetails.issueDt": bson.M{"$gte": issueStartDate, "$lte": issueEndDate},
		},
	).Decode(&invoice)
	if err != nil {
		repo.log.Errorf("invoice for billing period from %s to %s not found in repository, %v", issueStartDate, issueEndDate, err)
		return nil, err
	}

	return invoice, nil
}

func (repo *invoiceRepo) StoreOne(ctx context.Context, document interface{}) error {
	repo.log.Debug("store invoice")

	//ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	//defer cancel()

	collection := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName)
	_, err := collection.InsertOne(ctx, document)

	if err != nil {
		repo.log.Errorf("failed to store repurchase invoice details: %v", err)
		return err
	}

	return nil
}

func (repo *invoiceRepo) StoreMany(ctx context.Context, documents ...interface{}) error {
	repo.log.Debugf("store %d invoice documents", len(documents))

	//ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	//defer cancel()

	collection := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName)
	_, err := collection.InsertMany(ctx, documents)

	if err != nil {
		repo.log.Errorf("failed to store repurchase invoice details: %v", err)
		return err
	}

	return nil
}

func (repo *invoiceRepo) StoreManyWithinTransaction(ctx context.Context, documents ...interface{}) error {
	repo.log.Infof("store %d invoice documents", len(documents))

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := repo.client.StartSession()
	if err != nil {
		repo.log.Errorf("failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionContext mongodriver.SessionContext) (interface{}, error) {
		collection := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName)

		var result *mongodriver.InsertOneResult
		for _, document := range documents {
			switch doc := document.(type) {
			case interface{}:
				repo.log.Infof("storing %T", doc)
				result, err = collection.InsertOne(sessionContext, doc)
				if err != nil {
					return nil, err
				}
			default:
				repo.log.Warnf("trying to store unknown document type %T - ignoring", doc)
			}
		}

		return result, nil
	}

	_, err = session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		repo.log.Errorf("failed to store documents: %v", err)
		return err
	}

	return nil
}

func (repo *invoiceRepo) CountInvoices(ctx context.Context, customerId string, from string, to string) (int64, error) {
	repo.log.Infof("count invoices of customerId: %s, billing period: from %s to %s", customerId, from, to)

	collection := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName)
	opts := options.Count().SetMaxTime(60 * time.Second)
	count, err := collection.CountDocuments(ctx,
		bson.D{
			{"header.content.type", "invoice"},
			{"payload.invoiceDetails.billingStartDt", from},
			{"payload.invoiceDetails.billingEndDt", to},
			{"payload.customerDetails.customerId", customerId},
		},
		opts)

	if err != nil {
		repo.log.Errorf("failed to count invoices: %v", err)
		return 0, err
	}

	repo.log.Infof("found %d invoices of customerId: %s, billing period: from %s to %s", count, customerId, from, to)
	return count, nil
}

func (repo *invoiceRepo) CountRepurchaseInvoices(ctx context.Context, customerId string, from string, to string) (int64, error) {
	repo.log.Infof("count repurchase invoices of customerId: %s, billing period: from %s to %s", customerId, from, to)

	collection := repo.client.Database(repo.db).Collection(repo.cfg.Mongo.InvoiceCollectionName)
	opts := options.Count().SetMaxTime(60 * time.Second)
	count, err := collection.CountDocuments(ctx,
		bson.D{
			{"header.content.type", "repurchase"},
			{"payload.invoiceDetails.billingStartDt", from},
			{"payload.invoiceDetails.billingEndDt", to},
			{"payload.sellerDetails.customerId", customerId},
		},
		opts)

	if err != nil {
		repo.log.Errorf("failed to count repurchase invoices: %v", err)
		return 0, err
	}

	repo.log.Infof("found %d repurchase invoices of customerId: %s, billing period: from %s to %s", count, customerId, from, to)
	return count, nil
}
