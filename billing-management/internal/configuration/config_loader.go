package configuration

import (
	"errors"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	fu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/file_utils"
	"strings"
	"time"
)

type Parameter string

type ConfigService interface {
	Get(key string) interface{}
	GetFloat64(key string) float64
	GetBool(key string) bool
	GetIntSlice(key string) []int
	GetInt(key string) int
	GetStringMap(key string) map[string]interface{}
	GetString(key string) string
	GetStringSlice(key string) []string
	GetStringMapString(key string) map[string]string
	GetDuration(key string) time.Duration
	GetTime(key string) time.Time
	AllSettings() map[string]interface{}
	IsSet(key string) bool
}

// Parameter names
const (
	// Logger
	LogLevel     Parameter = "service.logger.level"
	LogFormatter Parameter = "service.logger.formatter"

	// HTTP server
	ListenPort Parameter = "service.http.port"

	// Database
	DbHost Parameter = "service.persistence.host"
	DbPort Parameter = "service.persistence.port"
	DbUser Parameter = "service.persistence.user"
	DbPass Parameter = "service.persistence.password"
	DbName Parameter = "service.persistence.database"

	// Mongo
	MongoHost Parameter = "service.mongo.host"
	MongoPort Parameter = "service.mongo.port"
	MongoUser Parameter = "service.mongo.user"
	MongoPass Parameter = "service.mongo.password"
	MongoName Parameter = "service.mongo.database"

	// Mongo
	ManagementAndLoginHost   Parameter = "service.management-and-login.host"
	ManagementAndLoginPort   Parameter = "service.management-and-login.port"
	ManagementAndLoginPrefix Parameter = "service.management-and-login.prefix"

	// RabbitMQ
	ExchangeName  Parameter = "rabbitproducer.ExchangeName"
	ExchangeType  Parameter = "rabbitproducer.ExchangeType"
	QueueBindings Parameter = "rabbitproducer.QueueBindings"
	RabbitUrl     Parameter = "rabbitproducer.RabbitMQConfig.RabbitUrl"
	IsDurable     Parameter = "rabbitproducer.RabbitMQConfig.IsDurable"
	IsAutoDelete  Parameter = "rabbitproducer.RabbitMQConfig.IsAutoDelete"
)

// Default values
const (
	LogLevelDefault     = "INFO"
	LogFormatterDefault = `%{time:2006-01-02 15:04:05.000} [%{shortpkg}] %{shortfunc} [%{level:.5s}] %{color:reset}%{message}`
	ListenPortDefault   = 8080

	DbHostDefault = "mysql"
	DbPortDefault = "3306"

	MongoHostDefault = "mongo"
	MongoPortDefault = "27017"
)

var (
	logger = logging.MustGetLogger("config_loader")
	CS     ConfigService
)

func LoadConfig(configFilePath string) error {
	CS = readConfigFromFile(configFilePath)
	if CS == nil {
		panic(errors.New("error config read"))
	}
	return nil
}

func GetConfig(configFilePath string) ConfigService {
	return readConfigFromFile(configFilePath)
}

// readConfigFromFile - reads config parameters from given file.
// If ENV variable exist parameter is overwritten by its value.
func readConfigFromFile(configFilePath string) ConfigService {
	setDefaults()

	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(fu.ExtractFileNameFromAbsolutePath(configFilePath)) // name of config file
	viper.AddConfigPath(fu.ExtractPathFromAbsolutePath(configFilePath))     // path to look for the config file in
	viper.SetConfigType(fu.ExtractFileExtensionFromAbsolutePath(configFilePath))

	viper.AutomaticEnv() // override config parameters from ENV variables (if exists)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logger.Errorf("Error while reading config file: %s", err)
	}

	return viper.GetViper()
}

func setDefaults() {
	viper.SetDefault(string(LogLevel), LogLevelDefault)
	viper.SetDefault(string(LogFormatter), LogFormatterDefault)
	viper.SetDefault(string(ListenPort), ListenPortDefault)

	viper.SetDefault(string(DbHost), DbHostDefault)
	viper.SetDefault(string(DbPort), DbPortDefault)

	viper.SetDefault(string(MongoHost), MongoHostDefault)
	viper.SetDefault(string(MongoPort), MongoPortDefault)
}
