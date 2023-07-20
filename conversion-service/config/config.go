package config

import (
	"errors"
	"github.com/spf13/viper"
	fu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/file_utils"
	"log"
	"strings"
	"time"
)

type Config struct {
	Service            ServiceConfig
	HtmlToPdfConverter HtmlToPdfConverterConfig
	Template           []TemplateEngineConfig
	Logger
	Language string
}

/*todo add provider config*/

type HtmlToPdfConverterConfig struct {
	Url             string
	Dpi             int
	MarginLRTB      []int
	Timeout         time.Duration
	DebugJavascript bool
}

type TemplateEngineConfig struct {
	Type               string
	Versions           []string
	Path               string
	FooterTemplatePath string
	FooterMargin       int
	Conditions         map[string]interface{}
}

type ServiceConfig struct {
	ServiceName    string
	Port           string
	ReleaseMode    bool
	TrustedProxies []string
	DebugMode      bool
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

func GetConfiguration(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	var cfg Config

	err = cfgFile.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func GetConfigurationLocal(configPath string) (*Config, error) {
	cfgFile, err := LoadConfigLocal(configPath)
	if err != nil {
		return nil, err
	}
	var cfg Config

	err = cfgFile.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
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

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}
