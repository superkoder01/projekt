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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSession interface {
	Aggregate(pipeline mongo.Pipeline, collection string) (*mongo.Cursor, error)
	List(query interface{}, collection string, options *options.FindOptions) (*mongo.Cursor, error)
	Find(query interface{}, collection string) (*mongo.Cursor, error)
	FindOne(query interface{}, collection string) *mongo.SingleResult
	Create(model model.Model, collection string) (*mongo.InsertOneResult, error)
	Update(filter interface{}, model interface{}, collection string, opts *options.FindOneAndReplaceOptions) *mongo.SingleResult
	Count(filter interface{}, collection string) (int64, error)
}
