package redis

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/op/go-logging"
	"io"
	"net"
	"os"
	"time"
)

import (
	"context"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/tests/configuration"
)

var (
	logger = logging.MustGetLogger("tests")
)

// RunRedis - returns containerId and IP of container
func RunRedis() (string, string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Errorf("%s", err)
	}

	rd := conf.GetTestConfig().GetTestRedis()
	image := rd.Image

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
		Env:   []string{"PASSWORD=" + rd.Username},
	}, nil, nil, nil, "redis")
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

func StopRedis(id string) {
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

func WaitForRedisServerReadiness() bool {
	rd := conf.GetTestConfig().GetTestRedis()
	for i := 1; i < 50; i++ {
		logger.Debugf("Waiting %d second(s) for redis instance readiness...", i)
		time.Sleep(time.Duration(1) * time.Second)
		_, err := net.Dial("tcp", rd.Address)
		if err != nil {
			continue
		}
		logger.Debug("Redis instance is up")
		return true
	}
	return false
}
