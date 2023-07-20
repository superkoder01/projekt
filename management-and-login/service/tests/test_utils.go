package tests

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/op/go-logging"
	i "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/init"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/configuration"
	"net"
	"net/http"
	"reflect"
	"time"
)

var (
	dbCf                  *conf.TestDb
	rdCf                  *conf.TestRedis
	bCf                   *conf.TestBlockchainAdapter
	authConfig            *auth.AuthConfig
	logger                = logging.MustGetLogger("billing_dao")
	df                    billing_dao.DaoFactory
	sf                    service.ServiceFactory
	mariaDbID             string
	mariaDbIP             string
	rabbitMqID            string
	rabbitMqIP            string
	redisID               string
	redisIP               string
	s                     mysql.Session
	rCli                  *redis.Client
	blockchainAccountStub = impl.BlockchainAccount{
		UserName:   "user123",
		PublicKey:  "key123",
		AccAddress: "address123",
	}
)

func init() {
	i.InitConfig()
}

func mariaDBConnect(ip string) (mysql.Session, error) {
	var err error

	if s == nil {
		// Read config
		// TODO initialize test configuration somewhere
		dbCf = conf.GetTestConfig().GetTestDb()

		// Create SQL URL
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/?tls=false&autocommit=true&parseTime=True&loc=Europe%%2FWarsaw", dbCf.User, dbCf.Password, ip, dbCf.Port)
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
		Create(&entity.Role{Name: entity.PROSUMER}).
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
		if reflect.TypeOf(en) == reflect.TypeOf(&entity.Worker{}) {
			ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
				Where("SUPERVISOR > ?", "0").
				Delete(en).
				Error()
		}
		err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
			Where("id > ?", "0").
			Delete(en).
			Error()
		if err != nil {
			logger.Errorf("flushing schema error: %s", err)
		}
	}

	err = ses.Exec(fmt.Sprintf("USE %s", dbCf.Database)).
		Create(&entity.Provider{ID: 1, Name: "Test", NIP: "123456789"}).
		Error()
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
		return err
	}

	return err
}

func redisConnect(ip string) (*redis.Client, error) {
	var err error

	if rCli == nil {
		// Read config
		// TODO initialize test configuration somewhere
		rdCf = conf.GetTestConfig().GetTestRedis()

		// Open connection on given URL with given driver
		for i := 1; i < 50; i++ {
			logger.Debugf("Waiting %d second(s) for redis instance readiness...", i)
			time.Sleep(time.Duration(1) * time.Second)
			_, err := net.Dial("tcp", net.JoinHostPort(ip, "6379"))
			if err != nil {
				continue
			}
			logger.Debug("Redis instance is up")
			break
		}

		// Create Redis client
		rCli = redis.NewClient(&redis.Options{
			Addr:     net.JoinHostPort(ip, "6379"),
			Password: rdCf.Password,
		})

		if rCli == nil {
			logger.Errorf("Failed to connect to redis: %s with password: %s", ip, rdCf.Password)
			return nil, err
		}
	}

	return rCli, nil
}

func blockchainAdapterStub(host string, port string, endpoint string) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		b, _ := json.Marshal(blockchainAccountStub)
		w.Write(b)
	})

	http.ListenAndServe(net.JoinHostPort(host, port), nil)
}
