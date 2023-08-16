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
package init

import (
	"github.com/op/go-logging"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/configuration"
	"os"
)

const (
	CONFIG_FILE_PATH = "CONFIG_FILE_PATH"
)

var (
	logger = logging.MustGetLogger("init_config")
)

func init() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
	logger.Debug("Configuration loaded succesfully")
}

func initConfig() error {
	configPath := os.Getenv(CONFIG_FILE_PATH)
	err := conf.LoadConfig(configPath)
	if err != nil {
		return err
	}
	logger.Debug(conf.GetAuthConfig().KeyPath)
	// setting logger
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(conf.CS.GetString(string(conf.LogFormatter)))
	backendFormatter := logging.NewBackendFormatter(logBackend, format)
	logBackendLeveled := logging.AddModuleLevel(backendFormatter)
	loglevel, err := logging.LogLevel(conf.CS.GetString(string(conf.LogLevel)))
	if err != nil {
		return err
	}

	logBackendLeveled.SetLevel(loglevel, "")
	logging.SetBackend(logBackendLeveled)

	logger.Debugf("Config file: %s read", configPath)

	return nil
}

func InitConfig() {

}
