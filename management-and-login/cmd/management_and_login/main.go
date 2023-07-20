package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/handler"
	api "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/server"
	i "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/init"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
)

var (
	logger = logging.MustGetLogger("main")
)

func main() {
	// Initialize configuration
	i.InitConfig()
	// Initialize MySQL database connection
	dbCf := conf.GetDatabaseConfig()
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&parseTime=True&loc=Europe%%2FWarsaw",
		dbCf.User,
		dbCf.Password,
		dbCf.Host,
		dbCf.Port,
		dbCf.Database)
	db, err := mysql.NewSession(source)
	if err != nil {
		logger.Errorf("mysql database connection error: %s", err)
		panic(err)
	}

	// Initialize Redis connection
	rCf := conf.GetRedisConfig()
	rCli := redis.NewClient(&redis.Options{
		Addr:     rCf.Address,
		Password: rCf.Password,
	})
	_, err = rCli.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("redis connection error")
		panic(err)
	}

	// Initialize DAO factory
	df := bd.NewDaoFactory(db)

	// Initialize Service factory
	cf := service.NewServiceFactory(df, rCli)

	// Initialize Handler factory
	hf := handler.NewHandlerFactory(cf)

	logger.Debug("Application initialization completed")

	// Start HTTP server
	api.NewHttpServer(hf).Run()
}
