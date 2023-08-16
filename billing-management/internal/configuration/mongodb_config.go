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
package configuration

type MongoConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func GetMongoDatabaseConfig() *MongoConfig {
	cf := MongoConfig{
		User:     CS.GetString(string(MongoUser)),
		Password: CS.GetString(string(MongoPass)),
		Host:     CS.GetString(string(MongoHost)),
		Port:     CS.GetString(string(MongoPort)),
		Database: CS.GetString(string(MongoName)),
	}

	return &cf
}
