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
package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
)

type mongoSession struct {
	Database *mongo.Database
	Ctx      context.Context
}

func NewMongoSession(config configuration.MongoConfig) (*mongoSession, error) {
	host := fmt.Sprintf("%s:%s", config.Host, config.Port)
	opts := options.Client().SetHosts([]string{host})
	opts.SetAuth(options.Credential{Username: config.User, Password: config.Password, AuthMechanism: "SCRAM-SHA-1", AuthSource: "admin"})
	//opts.Auth
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.Database)
	return &mongoSession{Database: database}, err
}

func (s *mongoSession) FindOne(query interface{}, collection string) *mongo.SingleResult {
	return s.Database.Collection(collection).FindOne(s.Ctx, query)
}

func (s *mongoSession) Find(query interface{}, collection string) (*mongo.Cursor, error) {
	return s.Database.Collection(collection).Find(s.Ctx, query)
}

func (s *mongoSession) List(query interface{}, collection string, options *options.FindOptions) (*mongo.Cursor, error) {
	return s.Database.Collection(collection).Find(s.Ctx, query, options)
}



