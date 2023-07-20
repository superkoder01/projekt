package init

import (
	"github.com/op/go-logging"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	"os"
)

const (
	CONFIG_FILE_PATH = "CONFIG_FILE_PATH"
	RBAC_FILE_PATH   = "RBAC_FILE_PATH"
)

var (
	logger = logging.MustGetLogger("init")
	// Global logger config
	LogBackend logging.LeveledBackend
	Level      logging.Level
)

// Initialize config etc, panic on error
func init() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
	err = initRbacConf()
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
	// setting logger
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(conf.CS.GetString(string(conf.LogFormatter)))
	backendFormatter := logging.NewBackendFormatter(logBackend, format)
	logBackendLeveled := logging.AddModuleLevel(backendFormatter)
	loglevel, err := logging.LogLevel(conf.CS.GetString(string(conf.LogLevel)))
	if err != nil {
		return err
	}

	LogBackend = logBackendLeveled
	Level = loglevel

	logBackendLeveled.SetLevel(loglevel, "")
	logging.SetBackend(logBackendLeveled)

	logger.Debugf("Config file: %s read", configPath)

	return nil
}

func initRbacConf() error {
	rbacPath := os.Getenv(RBAC_FILE_PATH)
	err := conf.LoadRBACConfig(rbacPath)
	if err != nil {
		logger.Errorf("Error while reading RBAC config file: %s", err)
		return err
	}
	logger.Debugf("Config file: %s read", rbacPath)
	return nil
}

func InitConfig() {

}
