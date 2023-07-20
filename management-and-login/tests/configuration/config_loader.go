package configuration

import (
	"github.com/op/go-logging"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	"os"
)

type testConfig struct {
	Db                    TestDb                `yaml:"db"`
	Rabbit                TestRabbitMQ          `yaml:"rabbit"`
	Redis                 TestRedis             `yaml:"redis"`
	KeyPath               string                `yaml:"keyPath"`
	AccessExpirationTime  int                   `yaml:"accessExpirationTime"`
	RefreshExpirationTime int                   `yaml:"refreshExpirationTime"`
	BlockchainAdapter     TestBlockchainAdapter `yaml:"blockchain"`
}

type TestDb struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Image    string `yaml:"image"`
}

type TestRabbitMQ struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Queue    string `yaml:"queue"`
	Image    string `yaml:"image"`
}

type TestRedis struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Image    string `yaml:"image"`
}

type TestBlockchainAdapter struct {
	AdapterHost string `yaml:"adapterHost"`
	AdapterPort string `yaml:"adapterPort"`
	Endpoint    string `yaml:"endpoint"`
}

type Parameter string

var (
	testConf testConfig
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

	// Rabbit
	EmailHost     Parameter = "service.rabbit.email.host"
	EmailPort     Parameter = "service.rabbit.email.port"
	EmailUser     Parameter = "service.rabbit.email.user"
	EmailPassword Parameter = "service.rabbit.email.password"
	EmailQueue    Parameter = "service.rabbit.email.queue"
	EmailImage    Parameter = "service.rabbit.email.image"

	// Redis
	RedisAddress Parameter = "service.redis.address"
	RedisUser    Parameter = "service.redis.username"
	RedisPass    Parameter = "service.redis.password"
	RedisImage   Parameter = "service.redis.image"

	// Blockchain adapter
	AdapterHost     Parameter = "service.blockchain.adapterHost"
	AdapterPort     Parameter = "service.blockchain.adapterPort"
	AdapterEndpoint Parameter = "service.blockchain.endpoint"

	KeyPath                    Parameter = "service.auth.keyPath"
	AccessTokenExpirationTime  Parameter = "service.auth.accessExpirationTime"
	RefreshTokenExpirationTime Parameter = "service.auth.refreshExpirationTime"
)

func GetTestConfig() *testConfig {
	if (testConfig{}) == testConf {
		logger.Info("Reading config")
		testConfigPath := os.Getenv("TEST_CONFIG_PATH")
		if testConfigPath == "" {
			logger.Error("Set TEST_CONFIG_PATH env variable first")
		}

		c, err := conf.GetConfig(testConfigPath)
		if err != nil {
			logger.Error("Could not read test config")
			panic(err)
		}

		tdb := TestDb{
			User:     c.GetString(string(DbUser)),
			Password: c.GetString(string(DbPass)),
			Host:     c.GetString(string(DbHost)),
			Port:     c.GetString(string(DbPort)),
			Database: c.GetString(string(DbName)),
			Image:    c.GetString(string(DbImage)),
		}

		te := TestRabbitMQ{
			Host:     c.GetString(string(EmailHost)),
			Port:     c.GetString(string(EmailPort)),
			User:     c.GetString(string(EmailUser)),
			Password: c.GetString(string(EmailPassword)),
			Queue:    c.GetString(string(EmailQueue)),
			Image:    c.GetString(string(EmailImage)),
		}

		tr := TestRedis{
			Username: c.GetString(string(RedisUser)),
			Password: c.GetString(string(RedisPass)),
			Address:  c.GetString(string(RedisAddress)),
			Image:    c.GetString(string(RedisImage)),
		}

		tb := TestBlockchainAdapter{
			AdapterHost: c.GetString(string(AdapterHost)),
			AdapterPort: c.GetString(string(AdapterPort)),
			Endpoint:    c.GetString(string(AdapterEndpoint)),
		}

		testConf.Db = tdb
		testConf.Rabbit = te
		testConf.Redis = tr
		testConf.BlockchainAdapter = tb
		testConf.KeyPath = c.GetString(string(KeyPath))
		testConf.AccessExpirationTime = c.GetInt(string(AccessTokenExpirationTime))
		testConf.RefreshExpirationTime = c.GetInt(string(RefreshTokenExpirationTime))
	}

	return &testConf
}

func (tcf *testConfig) GetTestDb() *TestDb {
	return &tcf.Db
}

func (tcf *testConfig) GetTestKeyPath() string {
	return tcf.KeyPath
}

func (tcf *testConfig) GetTestRabbitMQ() *TestRabbitMQ {
	return &tcf.Rabbit
}

func (tcf *testConfig) GetTestRedis() *TestRedis {
	return &tcf.Redis
}

func (tcf *testConfig) GetTestBlockchain() *TestBlockchainAdapter {
	return &tcf.BlockchainAdapter
}
