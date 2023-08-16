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
package db

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/tests/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net"
	"os"
	"time"
)

var (
	logger = logging.MustGetLogger("tests")
	conn   *gorm.DB
)

// RunMariaDB - returns containerId and IP of container
func RunMariaDB() (string, string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Errorf("%s", err)
	}

	// TODO initialize test configuration somewhere
	db := conf.GetTestConfig().GetTestDb()
	image := db.Image

	imgSum, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		logger.Errorf("%s", err)
	}

	var imageAlreadyPresent bool
	for _, imgInfo := range imgSum {
		if imageAlreadyPresent {
			break
		}
		for _, repoTag := range imgInfo.RepoTags {
			if repoTag == image {
				imageAlreadyPresent = true
				break
			}
		}
	}

	if !imageAlreadyPresent {
		out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
		if err != nil {
			logger.Errorf("%s", err)
		}
		defer out.Close()
		io.Copy(os.Stdout, out)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Env:   []string{"MARIADB_ROOT_PASSWORD=" + db.Password},
	}, nil, nil, nil, "mariadb")
	if err != nil {
		logger.Errorf("%s", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		logger.Errorf("%s", err)
	}

	cInfo, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		logger.Errorf("%s", err)
	}

	ip := cInfo.NetworkSettings.IPAddress

	return resp.ID, ip
}

func MariaDBStop(id string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Errorf("Docker customer_account failed: %s", err)
	}

	err = cli.ContainerStop(ctx, id, nil)
	if err != nil {
		logger.Errorf("Container %s stop failed: %s", id, err)
	}

	err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
	if err != nil {
		logger.Errorf("Container %s remove failed: %s", id, err)
	}

	_, err = cli.VolumesPrune(ctx, filters.Args{})
	if err != nil {
		logger.Errorf("Volume prune failed: %s", err)
	}
}

func MariaDBConnect(ip string) (*gorm.DB, error) {
	var err error

	if conn == nil {
		// Read config
		// TODO initialize test configuration somewhere
		dbCf := conf.GetTestConfig().GetTestDb()

		// Create SQL URL
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/?tls=false&autocommit=true", dbCf.User, dbCf.Password, ip, dbCf.Port)
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

		conn, err = gorm.Open(mysql.Open(source), &gorm.Config{
			SkipDefaultTransaction: true,
		})

		if err != nil {
			logger.Errorf("Failed to connect to database: %s:%s with user: %s", ip, dbCf.Port, dbCf.User)
			return nil, err
		}
	}

	return conn, nil
}

func MariaDBInitSchema(conn *gorm.DB) error {
	conn.Exec("CREATE DATABASE if not exists c4e character set UTF8")
	conn.Exec("USE c4e")
	err := conn.AutoMigrate(&entity.User{}, &entity.CustomerAccount{})
	if err != nil {
		logger.Errorf("creating schema error: %s", err)
	}
	return err
}

func MariaDBFlushData(conn *gorm.DB) error {
	entities := []entity.Entity{&entity.User{}, &entity.CustomerAccount{}}

	var err error
	for _, en := range entities {
		err = conn.Where("id > ?", "0").Delete(en).Error
		if err != nil {
			logger.Errorf("flushing schema error: %s", err)
		}
	}
	return err
}
