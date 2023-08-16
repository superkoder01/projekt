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

	// Auth
	KeyPath Parameter = "service.auth.keyPath"

	SkipJwtCheckUrls Parameter = "service.skipJwtCheck.urls"
)

// Default values
const (
	LogLevelDefault     = "INFO"
	LogFormatterDefault = `%{time:2006-01-02 15:04:05.000} [%{shortpkg}] %{shortfunc} [%{level:.5s}] %{color:reset}%{message}`
	ListenPortDefault   = 3011
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

}
