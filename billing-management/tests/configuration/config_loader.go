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
	"os"
	"strings"
	"time"
)

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

type TestConfig struct {
	Db TestDb `yaml:"db"`
}

type TestDb struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Image    string `yaml:"image"`
}

type Parameter string

var (
	testConf TestConfig
	logger   = logging.MustGetLogger("tests")
)

const (
	// HTTP server
	ListenPort Parameter = "service.http.port"

	// Database
	DbHost  Parameter = "service.persistence.host"
	DbPort  Parameter = "service.persistence.port"
	DbUser  Parameter = "service.persistence.user"
	DbPass  Parameter = "service.persistence.password"
	DbName  Parameter = "service.persistence.database"
	DbImage Parameter = "service.persistence.image"
)

func GetTestConfig() *TestConfig {
	if (TestConfig{}) == testConf {
		logger.Info("Reading config")
		testConfigPath := os.Getenv("TEST_CONFIG_PATH")
		if testConfigPath == "" {
			logger.Error("Set TEST_CONFIG_PATH env variable first")
		}

		c := readConfigFromFile(testConfigPath)
		if c == nil {
			logger.Error("Could not read test config")
		}

		tdb := TestDb{
			User:     c.GetString(string(DbUser)),
			Password: c.GetString(string(DbPass)),
			Host:     c.GetString(string(DbHost)),
			Port:     c.GetString(string(DbPort)),
			Database: c.GetString(string(DbName)),
			Image:    c.GetString(string(DbImage)),
		}

		testConf.Db = tdb
	}

	return &testConf
}

func (tcf *TestConfig) GetTestDb() *TestDb {
	return &tcf.Db
}

func readConfigFromFile(configFilePath string) ConfigService {
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
