package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const (
	ALL = 1
	RUNNING = 2
	STOPPED = 3
)

var cli *client.Client

// Singleton Pattern
func GetDockerClientInstance() *client.Client {
	if cli == nil {
		instance, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		if err != nil {
			panic(err)
		}

		return instance
	}
	return cli
}

// Hiển thị danh sách container tùy theo type
func ListContainer(containerType int) (containers []types.Container, err error){
	ctx := context.Background()
	cli := GetDockerClientInstance()

	filters := filters.NewArgs()

	if containerType == RUNNING {
		filters.Add("status", "running")
	} else if containerType == STOPPED {
		filters.Add("status", "exited")
	}
	
	containers, err = cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filters,
	})

	if err != nil {
		return []types.Container{}, err
	}

	return containers, nil
}

// Stop container by ID
func StopContainerById(containerId string) error {
	ctx := context.Background()
	cli := GetDockerClientInstance()

	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		return err
	}

	return nil
}

// Start container by ID
func StartContainerById(containerId string) error {
	ctx := context.Background()
	cli := GetDockerClientInstance()

	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}