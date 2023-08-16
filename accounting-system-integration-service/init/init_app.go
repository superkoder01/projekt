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
	"os"

	"github.com/op/go-logging"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
)

const (
	CONFIG_FILE_PATH = "CONFIG_FILE_PATH"
	RBAC_FILE_PATH   = "RBAC_FILE_PATH"
)

var (
	logger = logging.MustGetLogger("init_app")
	// Global logger config
	LogBackend logging.LeveledBackend
	Level      logging.Level
)

func init() {
	initConfig()
	initRbacConf()
	logger.Debug("Application initialization completed")
}

func initConfig() {
	configPath := os.Getenv(CONFIG_FILE_PATH)
	err := conf.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	// setting logger
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(conf.CS.GetString(string(conf.LogFormatter)))
	backendFormatter := logging.NewBackendFormatter(logBackend, format)
	logBackendLeveled := logging.AddModuleLevel(backendFormatter)
	loglevel, err := logging.LogLevel(conf.CS.GetString(string(conf.LogLevel)))
	if err != nil {
		panic(err)
	}

	LogBackend = logBackendLeveled
	Level = loglevel

	logBackendLeveled.SetLevel(loglevel, "")
	logging.SetBackend(logBackendLeveled)

	logger.Debugf("Configuration loaded succesfully: %s", configPath)
}

func initRbacConf() {
	rbacPath := os.Getenv(RBAC_FILE_PATH)
	err := conf.LoadRBACConfig(rbacPath)
	if err != nil {
		panic(err)
	}
}

func InitApp() {
}
