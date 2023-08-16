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
package mariadb_session

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DBSession interface {
	GetConnection() *gorm.DB
}

type session struct {
	db *gorm.DB
}

func New(cfg *config.AppConfig) (DBSession, error) {
	db, err := gorm.Open(mysql.Open(cfg.MariaDB.Dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to open connection to MariaDb: %s", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MariaDB.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MariaDB.MaxOpenConnections)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.MariaDB.ConnectionMaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MariaDB.ConnectionMaxLifetime) * time.Second)

	return &session{db: db}, nil
}

func (s *session) GetConnection() *gorm.DB {
	return s.db
}
