package tests

import (
	"fmt"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/configuration"
	"net"
	"time"
)

var (
	dbCf      *conf.TestDb
	logger    = logging.MustGetLogger("billing_dao")
	df        billing_dao.DaoFactory
	mariaDbID string
	mariaDbIP string
	s         mysql.Session
)

func mariaDBConnect(ip string) (mysql.Session, error) {
	var err error

	if s == nil {
		// Read config
		// TODO initialize test configuration somewhere
		dbCf = conf.GetTestConfig().GetTestDb()

		// Create SQL URL
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/?tls=false&autocommit=true&parseTime=True", dbCf.User, dbCf.Password, ip, dbCf.Port)
		// Open connection on given URL with given driver
		for i := 1; i < 50; i++ {
			logger.Debugf("Waiting %d second(s) for mariadb instance readiness...", i)
			time.Sleep(time.Duration(1) * time.Second)
			_, err := net.Dial("tcp", net.JoinHostPort(ip, dbCf.Port))
			if err != nil {
				continue
			}
			logger.Debug("Mariadb instance is up")
			break
		}

		s, err = mysql.NewSession(source)

		if err != nil {
			logger.Errorf("Failed to connect to database: %s:%s with user: %s", ip, dbCf.Port, dbCf.User)
			return nil, err
		}
	}

	return s, nil
}

func mariaDBInitSchema(ses mysql.Session) error {
	err := ses.Exec(fmt.Sprintf("CREATE DATABASE if not exists %s character set UTF8", dbCf.Database)).Error()
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
		AutoMigrate(&entity.Provider{},
			&entity.Worker{},
			&entity.Role{},
			&entity.CustomerType{},
			&entity.CustomerAccount{},
			&entity.User{})
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
		Create(&entity.Role{Name: entity.SUPER_ADMIN}).
		Create(&entity.Role{Name: entity.ADMINISTRATOR_FULL}).
		Create(&entity.Role{Name: entity.ADMINISTRATOR_BASIC}).
		Create(&entity.Role{Name: entity.TRADER}).
		Create(&entity.Role{Name: entity.SUPER_AGENT}).
		Create(&entity.Role{Name: entity.AGENT}).
		Create(&entity.Role{Name: entity.ACCOUNTER}).
		Error()
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
		Create(&entity.CustomerType{Name: entity.CT_CONSUMER}).
		Create(&entity.CustomerType{Name: entity.CT_PROSUMER}).
		Create(&entity.CustomerType{Name: entity.CT_PRODUCER}).
		Error()
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
		Create(&entity.Provider{Name: "OVOO", NIP: "6762457439"}).
		Error()
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	return err
}

func mariaDBFlushData(ses mysql.Session) error {
	entities := []entity.Entity{&entity.User{},
		&entity.CustomerAccount{},
		&entity.Worker{},
		&entity.Provider{}}

	var err error
	for _, en := range entities {
		err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
			Where("id > ?", "0").
			Delete(en).
			Error()
		if err != nil {
			logger.Errorf("flushing schema error: %s", err)
		}
	}
	return err
}

func newBillingSession() (mysql.Session, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&parseTime=True", dbCf.User, dbCf.Password, mariaDbIP, dbCf.Port, dbCf.Database)

	return mysql.NewSession(url)
}
