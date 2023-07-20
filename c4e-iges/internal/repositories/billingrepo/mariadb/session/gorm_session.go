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
