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
package main

import (
	"fmt"
	"github.com/op/go-logging"
	server "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/handler"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/controller"
	rabbit "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/rabbitmq/handler"
	i "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/init"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mongo"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
	mongoservice "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

var (
	logger = logging.MustGetLogger("main")
)

func main() {
	// Initialize config etc, panic on error
	i.InitApp()

	//Initialize MySQL database connection
	dbConfig := conf.GetDatabaseConfig()
	mysqlSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&parseTime=True",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)
	db, err := mysql.NewSession(mysqlSource)
	if err != nil {
		logger.Errorf("mysql database connection error: %s", err)
	}

	//Initialize MongoDB database connection
	mongoConfig := conf.GetMongoDatabaseConfig()
	mongoDb, err := mongodb.NewMongoSession(*mongoConfig)

	//Initialize RabbitMQ
	rabbitConfig := conf.GetRabbitMQConfig()

	// Initialize DAO factory
	df := bd.NewDaoFactory(db)
	cf := mongo.NewCollectionFactory(mongoDb)

	// Initialize Service factory
	sf := service.NewServiceFactory(df)
	msf := mongoservice.NewServiceFactory(cf)

	// Initialize Handler factory
	hf := handler.NewHandlerFactory(sf)
	mcf := controller.NewControllerFactory(msf)

	eh := rabbit.NewEmailHandler(rabbitConfig, msf.New("CONTRACT"), msf.New("OFFER"))

	logger.Debug("Application initialization completed")

	// Start HTTP server
	server.NewHttpServer(hf, mcf, *eh).Run()
}
