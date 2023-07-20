package init

import (
	"github.com/op/go-logging"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/configuration"
	"os"
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
