package config

import (
	"RDN-application/pkg/logger"
	"errors"
	"github.com/spf13/viper"
	common "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	"log"
)

type DataType string

const (
	Hour   DataType = "hour"
	Fixing DataType = "fixing"
	Value  DataType = "value"
)

type AppConfig interface {
	GetServiceConfig() *serviceConfig
	GetLoggerConfig() *logger.LoggerConfig
	GetStoreConfig() *storeConfig
	GetCollectorConfig() *collectorConfig
	GetPortsConfig() *portsConfig
	GetRabbitMqProducerConfig() *common.RabbitMQProducerConfig
	GetNotificationReceiversConfig() *notifierConfig
	GetServerConfig() *serverConfig
}

type config struct {
	AppConfig applicationConfig
}

func (cfg *config) GetServiceConfig() *serviceConfig {
	return &cfg.AppConfig.Service
}

func (cfg *config) GetCollectorConfig() *collectorConfig {
	return &cfg.AppConfig.Collector
}

func (cfg *config) GetLoggerConfig() *logger.LoggerConfig {
	return &cfg.AppConfig.Logger
}

func (cfg *config) GetStoreConfig() *storeConfig {
	return &cfg.AppConfig.Store
}

func (cfg *config) GetRabbitMqProducerConfig() *common.RabbitMQProducerConfig {
	return &cfg.AppConfig.RabbitProducer
}

func (cfg *config) GetPortsConfig() *portsConfig {
	return &cfg.AppConfig.Ports
}

func (cfg *config) GetNotificationReceiversConfig() *notifierConfig {
	return &cfg.AppConfig.Notifications
}

func (cfg *config) GetServerConfig() *serverConfig {
	return &cfg.AppConfig.Server
}

type applicationConfig struct {
	Service        serviceConfig
	RabbitProducer common.RabbitMQProducerConfig
	Store          storeConfig
	Collector      collectorConfig
	Ports          portsConfig
	Server         serverConfig
	Notifications  notifierConfig
	Logger         logger.LoggerConfig
}

type serviceConfig struct {
	ServiceName string
}

type serverConfig struct {
	Port           string
	TrustedProxies []string
}

type collectorConfig struct {
	RetryTime      int
	DestinationUrl string
	RequestMethod  string
	Headers        map[string][]string
	HeaderNames    []string
	Scrapper       Scrapper
}

type Scrapper struct {
	MainSelector      string
	ConditionSelector ConditionSelector
	DataSelector      DataSelector
}

type ConditionSelector struct {
	SelectorName   string
	ConditionValue string
}

type DataSelector struct {
	SelectorName string
	Data         [2]Data
}

type Data struct {
	SelectorName string
	DataType     DataType
	DataName     string
}

type notifierConfig struct {
	Notification []Notification
}

type Notification struct {
	NotificationType string
	Receivers        []string
}

type portsConfig struct {
	NotificationServiceType string
	StoreRepoType           string
	CollectorType           string
}

type storeConfig struct {
	DbName         string
	CollectionName string
	BucketName     string
	Uri            string
	Timeout        int
}

/*todo commons*/
// Load config file from given path

func GetConfigurationLocal(configPath string) (*config, error) {
	cfgFile, err := LoadConfigLocal(configPath)
	if err != nil {
		return nil, err
	}
	var cfg applicationConfig

	err = cfgFile.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &config{AppConfig: cfg}, nil
}

// Load config file from given path
func LoadConfigLocal(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func GetConfigPath() string {
	return "./config/config-local"
}
