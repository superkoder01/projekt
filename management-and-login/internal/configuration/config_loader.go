package configuration

import (
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
	ApiPrefix  Parameter = "service.http.apiPrefix"

	// Blockchain adapter
	AdapterHost     Parameter = "service.blockchain.adapterHost"
	AdapterPort     Parameter = "service.blockchain.adapterPort"
	AdapterEndpoint Parameter = "service.blockchain.endpoint"

	// Database
	DbHost Parameter = "service.persistence.host"
	DbPort Parameter = "service.persistence.port"
	DbUser Parameter = "service.persistence.user"
	DbPass Parameter = "service.persistence.password"
	DbName Parameter = "service.persistence.database"

	// Redis
	RedisUser    Parameter = "service.redis.username"
	RedisPass    Parameter = "service.redis.password"
	RedisAddress Parameter = "service.redis.address"

	// Auth
	KeyPath               Parameter = "service.auth.keyPath"
	AccessExpirationTime  Parameter = "service.auth.accessExpirationTime"
	RefreshExpirationTime Parameter = "service.auth.refreshExpirationTime"

	// Rabbit
	EmailHost     Parameter = "service.rabbit.email.host"
	EmailPort     Parameter = "service.rabbit.email.port"
	EmailUser     Parameter = "service.rabbit.email.user"
	EmailPassword Parameter = "service.rabbit.email.password"
	EmailQueue    Parameter = "service.rabbit.email.queue"

	// ActivationMessage
	EmailActivationTitle Parameter = "service.emailMessage.activation.title"
	EmailActivationBody  Parameter = "service.emailMessage.activation.body"
	// ResetPasswordMessage
	EmailResetPasswordTitle Parameter = "service.emailMessage.resetPassword.title"
	EmailResetPasswordBody  Parameter = "service.emailMessage.resetPassword.body"
)

// Default values
const (
	LogLevelDefault     = "INFO"
	LogFormatterDefault = `%{time:2006-01-02 15:04:05.000} [%{shortpkg}] %{shortfunc} [%{level:.5s}] %{color:reset}%{message}`
	ListenPortDefault   = 8181

	DbHostDefault = "mysql"
	DbPortDefault = "3306"
)

var (
	logger = logging.MustGetLogger("config_loader")
	CS     ConfigService
)

func LoadConfig(configFilePath string) error {
	var err error
	CS, err = readConfigFromFile(configFilePath)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig(configFilePath string) (ConfigService, error) {
	return readConfigFromFile(configFilePath)
}

// readConfigFromFile - reads config parameters from given file.
// If ENV variable exist parameter is overwritten by its value.
func readConfigFromFile(configFilePath string) (ConfigService, error) {
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
		return nil, err
	}

	return viper.GetViper(), nil
}

func setDefaults() {
	viper.SetDefault(string(LogLevel), LogLevelDefault)
	viper.SetDefault(string(LogFormatter), LogFormatterDefault)
	viper.SetDefault(string(ListenPort), ListenPortDefault)

	viper.SetDefault(string(DbHost), DbHostDefault)
	viper.SetDefault(string(DbPort), DbPortDefault)
}
