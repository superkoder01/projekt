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

type MSzafirConfig struct {
	Host                            string `yaml:"host"`
	Port                            string `yaml:"port"`
	Prefix                          string `yaml:"prefix"`
	Timestamp                       int64  `yaml:"timestamp"`
	Mode                            string `yaml:"mode"`
	UrlSigningCompleted             string `yaml:"urlSigningCompleted"`
	UrlSigningCompletedNotification string `yaml:"urlSigningCompletedNotification"`
	Password                        string `yaml:"password"`
	SignedContractsPath             string `yaml:"signedContractsPath"`
}

func GetMSzafirConfig() *MSzafirConfig {
	return &MSzafirConfig{
		Host:                            CS.GetString(string(MSzafirHost)),
		Port:                            CS.GetString(string(MSzafirPort)),
		Prefix:                          CS.GetString(string(MSzafirPrefix)),
		Timestamp:                       int64(CS.GetInt(string(MSzafirTimestamp))),
		Mode:                            CS.GetString(string(MSzafirMode)),
		UrlSigningCompleted:             CS.GetString(string(MSzafirUrlSigningCompleted)),
		UrlSigningCompletedNotification: CS.GetString(string(MSzafirUrlSigningCompletedNotification)),
		Password:                        CS.GetString(string(MSzafirPassword)),
		SignedContractsPath:             CS.GetString(string(MSzafirSignedContractsPath)),
	}
}
