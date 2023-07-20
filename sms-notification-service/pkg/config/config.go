package config

import (
	"errors"
	"github.com/spf13/viper"
	fu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/file_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	"log"
	"strings"
)

type SmsServiceConfig struct {
	Service  ServiceConfig
	Rabbitmq config.RabbitMQConsumerConfig
	Gateway  GatewayConfig
	Logger
}

type ServiceConfig struct {
	ServiceName    string
	WorkerPoolSize int
}

type SmtpConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	SenderName string
}

type GatewayConfig struct {
	Host         string
	ApiVersion   string
	ApiServiceId string
	ApiKey       string
	ApiUrl       string
	MethodType   string
	SenderId     string
}

type Logger struct {
	Development       bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Load config file from given path
func LoadConfig(filePath string) (*viper.Viper, error) {
	v := viper.New()

	v.SetEnvPrefix("")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigName(fu.ExtractFileNameFromAbsolutePath(filePath)) // name of config file
	v.AddConfigPath(fu.ExtractPathFromAbsolutePath(filePath))     // path to look for the config file in
	v.SetConfigType(fu.ExtractFileExtensionFromAbsolutePath(filePath))

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func GetSmsConfiguration(configPath string) (*SmsServiceConfig, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	var cfg SmsServiceConfig

	err = cfgFile.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}
