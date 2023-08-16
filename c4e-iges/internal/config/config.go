/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package config

import (
	"fmt"
	"github.com/spf13/viper"
	rabbitmqconfig "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	"os"
)

type cfg struct {
	Service AppConfig `mapstructure:"service"`
}

type invoiceContent struct {
	HeaderContentType                         string            `mapstructure:"headerContentType"`
	HeaderContentCategory                     string            `mapstructure:"headerContentCategory"`
	PayloadInvoiceDetailsType                 string            `mapstructure:"payloadInvoiceDetailsType"`
	PayloadPaymentDetailsPaymentTitleTemplate string            `mapstructure:"payloadPaymentDetailsPaymentTitleTemplate"`
	InvoiceLineItems                          map[string]string `mapstructure:"invoiceLineItems"`
}

type invoiceDetailsContent struct {
	HeaderContentType         string `mapstructure:"headerContentType"`
	HeaderContentCategory     string `mapstructure:"headerContentCategory"`
	PayloadInvoiceDetailsType string `mapstructure:"payloadInvoiceDetailsType"`
}

type repurchaseContent struct {
	HeaderContentType         string `mapstructure:"headerContentType"`
	HeaderContentCategory     string `mapstructure:"headerContentCategory"`
	PayloadInvoiceDetailsType string `mapstructure:"payloadInvoiceDetailsType"`
}

type repurchaseDetailsContent struct {
	HeaderContentType         string `mapstructure:"headerContentType"`
	HeaderContentCategory     string `mapstructure:"headerContentCategory"`
	PayloadInvoiceDetailsType string `mapstructure:"payloadInvoiceDetailsType"`
}

type AppConfig struct {
	ServiceName              string                                `mapstructure:"name"`
	ServiceVersion           string                                `mapstructure:"version"`
	Ports                    portsConfig                           `mapstructure:"ports"`
	Logger                   loggerConfig                          `mapstructure:"logger"`
	StatusService            statusServiceConfig                   `mapstructure:"statusservice"`
	InvoiceService           invoiceServiceConfig                  `mapstructure:"invoiceservice"`
	InvoiceContent           invoiceContent                        `mapstructure:"invoiceContent"`
	InvoiceDetailsContent    invoiceDetailsContent                 `mapstructure:"invoiceDetailsContent"`
	RepurchaseContent        repurchaseContent                     `mapstructure:"repurchaseContent"`
	RepurchaseDetailsContent repurchaseDetailsContent              `mapstructure:"repurchaseDetailsContent"`
	MariaDB                  mariadbConfig                         `mapstructure:"mariadb"`
	Mongo                    mongoConfig                           `mapstructure:"mongo"`
	RabbitConsumer           rabbitmqconfig.RabbitMQConsumerConfig `mapstructure:"rabbitconsumer"`
	RabbitProducer           rabbitmqconfig.RabbitMQProducerConfig `mapstructure:"rabbitproducer"`
}

type portsConfig struct {
	ContractRepoPort           string `mapstructure:"contractRepoPort"`
	InvoiceRepoPort            string `mapstructure:"invoiceRepoPort"`
	InvoicePublisherPort       string `mapstructure:"invoicePublisherPort"`
	InvoiceEventSubscriberPort string `mapstructure:"invoiceEventSubscriberPort"`
	AlarmServicePort           string `mapstructure:"alarmServicePort"`
}

type loggerConfig struct {
	Level           string `mapstructure:"level"`
	Encoding        string `mapstructure:"encoding"`
	TimestampLayout string `mapstructure:"timestamplayout"`
}

type statusServiceConfig struct {
	HttpPort int `mapstructure:"httpPort"`
}

type invoiceServiceConfig struct {
	InvoiceEventConsumerPoolSize int    `mapstructure:"invoiceEventConsumerPoolSize"`
	EmailBodyTemplate            string `mapstructure:"emailBodyTemplate"`
	EmailTitleTemplate           string `mapstructure:"emailTitleTemplate"`
}

type mariadbConfig struct {
	Dsn                   string `mapstructure:"dsn"`
	MaxIdleConnections    int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections    int    `mapstructure:"maxOpenConnections"`
	ConnectionMaxIdleTime int    `mapstructure:"connectionMaxIdleTime"`
	ConnectionMaxLifetime int    `mapstructure:"connectionMaxLifetime"`
}

type mongoConfig struct {
	DbName                 string `mapstructure:"dbname"`
	ContractCollectionName string `mapstructure:"contractCollectionName"`
	InvoiceCollectionName  string `mapstructure:"invoiceCollectionName"`
	Uri                    string `mapstructure:"uri"`
	Timeout                int    `mapstructure:"timeout"`
}

// LoadConfig loads app from files
func LoadConfig() (*AppConfig, error) {
	configPath := os.Getenv("CONFIGPATH")
	if configPath == "" {
		configPath = "./config/"
	}

	configName := os.Getenv("CONFIGNAME")
	if configName == "" {
		configName = "config-local"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %s", err)
	}

	// Set default values
	viper.SetDefault("service.statushandler.httpPort", 1234)
	viper.SetDefault("service.logger.level", "debug")
	viper.SetDefault("service.logger.encoding", "console")
	viper.SetDefault("service.logger.timestamplayout", "2006-01-02 15:04:05.000")
	viper.SetDefault("service.mariadb.maxIdleConnections", 10)
	viper.SetDefault("service.mariadb.maxOpenConnections", 15)
	viper.SetDefault("service.mariadb.connectionMaxIdleTime", 60)
	viper.SetDefault("service.mariadb.connectionMaxLifetime", 300)

	var c cfg
	err := viper.Unmarshal(&c)

	return &c.Service, err
}
