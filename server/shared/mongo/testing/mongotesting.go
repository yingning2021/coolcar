package mongotesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"testing"
)

const (
	image         = "mongo"
	containerPort = "27017/tcp"
)

// RunWithMongoInDocker runs the tests
func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	containerId := resp.ID
	defer func() {
		err = c.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(err)
		}
	}()

	err = c.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)

	return m.Run()

}
