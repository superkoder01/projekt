package config

import (
	"NotificationEmailService/config/smtp_security"
	"errors"
	"github.com/spf13/viper"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	fu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/file_utils"
	config2 "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	"log"
	"strings"
)

type EmailServiceConfig struct {
	Service    ServiceConfig
	Smtp       SmtpConfig
	Rabbitmq   config2.RabbitMQConsumerConfig
	Conversion conversion.Config
	Logger
}

type ServiceConfig struct {
	ServiceName           string
	TestMode              bool
	WorkerPoolSize        int
	ErrorEmailTo          []string
	ErrorEmailEnvironment string
}

type SmtpConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	SenderName string
	Security   smtp_security.Security
}

type GatewayConfig struct {
	Host         string
	ApiVersion   string
	ApiServiceId string
	ApiKey       string
	ApiUrl       string
	MethodType   string
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

func GetEmailConfiguration(configPath string) (*EmailServiceConfig, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	var cfg EmailServiceConfig

	err = cfgFile.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}
