package client

import (
	"RDN-application/pkg/config"
	"context"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"os"
)

type Client struct {
	influxdb.Client
}

func NewInfluxClient(cfg config.AppConfig) *Client {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		panic("INFLUXDB_TOKEN must be set")
	}

	influxClient := influxdb.NewClient(cfg.GetStoreConfig().Uri, dbToken)

	_, err := influxClient.Health(context.Background())

	if err != nil {
		panic("cannot connect to influx database !")
	}

	return &Client{influxClient}
}
