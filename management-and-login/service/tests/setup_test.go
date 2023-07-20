package tests

import (
	"errors"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/configuration"
	td "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/db"
	r "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/rabbitmq"
	rs "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/redis"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var err error

	// Run mock RabbitMQ server
	rabbitMqID, rabbitMqIP = r.RunRabbitMQ()

	// Run MariaDB instance on container
	mariaDbID, mariaDbIP = td.RunMariaDB()

	// Run redis instance on container
	redisID, redisIP = rs.RunRedis()

	// Connect to MariaDB
	s, err = mariaDBConnect(mariaDbIP)
	if err != nil {
		panic(err)
	}

	if !r.WaitForRabbitServerReadiness() {
		panic(errors.New("RabbitMQ instance failed to start"))
	}

	if !rs.WaitForRedisServerReadiness() {
		panic(errors.New("Redis instance failed to start"))
	}

	rCli, err = redisConnect(redisIP)
	if err != nil {
		panic(err)
	}

	// Initialize database schema
	mariaDBInitSchema(s)

	// Initialize DaoFactory
	df = billing_dao.NewDaoFactory(s)

	// Initialize ServiceFactory
	sf = service.NewServiceFactory(df, rCli)

	// AuthConfig
	authConfig = &auth.AuthConfig{
		KeyFilePath:                conf.GetTestConfig().KeyPath,
		RedisClient:                rCli,
		AccessTokenExpirationTime:  time.Minute * time.Duration(conf.GetTestConfig().AccessExpirationTime),
		RefreshTokenExpirationTime: time.Minute * time.Duration(conf.GetTestConfig().RefreshExpirationTime),
	}

	// Initialize blockchain adapter stub
	bCf = conf.GetTestConfig().GetTestBlockchain()
	go blockchainAdapterStub(bCf.AdapterHost, bCf.AdapterPort, bCf.Endpoint)

	// Run test cases
	code := m.Run()

	// Delete MariaDB container
	td.MariaDBStop(mariaDbID)

	// Delete RabbitMQ container
	r.StopRabbitMQ(rabbitMqID)

	// Delete Redis container
	rs.StopRedis(redisID)

	os.Exit(code)
}
