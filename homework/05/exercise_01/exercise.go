package exercise_01

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/olekukonko/tablewriter"
)

// Container State
const (
	ALL = 1
	RUNNING = 2
	STOPPED = 3
)

// Khai báo interface
type iDockerClient interface {
	ListContainer(int) (containers []types.Container, err error)
	StopContainerById(string) error
	StartContainerById(string) error
}

// Singleton Pattern
type DockerClient struct {
	client *client.Client
}

var docker *DockerClient

func GetDockerClientInstance() iDockerClient {
	if docker == nil {
		ins, err := connectToDockerClient()
		if err != nil {
			panic(err)
		}
		return &DockerClient{
			client: ins,
		}
	}
	return docker
}

func connectToDockerClient() (*client.Client, error) {
	ins, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// Hiển thị danh sách container tùy theo type
func (d *DockerClient) ListContainer(containerType int) (containers []types.Container, err error){
	ctx := context.Background()

	filters := filters.NewArgs()

	if containerType == RUNNING {
		filters.Add("status", "running")
	} else if containerType == STOPPED {
		filters.Add("status", "exited")
	}
	
	containers, err = d.client.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filters,
	})

	if err != nil {
		return []types.Container{}, err
	}

	return containers, nil
}

// Stop container by ID
func (d *DockerClient) StopContainerById(containerId string) error {
	ctx := context.Background()

	if err := d.client.ContainerStop(ctx, containerId, nil); err != nil {
		return err
	}

	return nil
}

// Start container by ID
func (d *DockerClient) StartContainerById(containerId string) error {
	ctx := context.Background()

	if err := d.client.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

// Print Detail
func Print(containers []types.Container) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "State"})

	var data [][]string
	for _, v := range containers {
		data = append(data, []string{v.ID[:12], v.Names[0], v.State})
	}

	table.AppendBulk(data)
	table.Render()
}


func DemoExercise_01() {
	dockerIns := GetDockerClientInstance()

	// Hiển thị danh sách tất cả container
	containerAll, _ := dockerIns.ListContainer(ALL)
	Print(containerAll)

	// Hiển thị danh sách các container đang running
	containerRunning, _ := dockerIns.ListContainer(RUNNING)
	Print(containerRunning)

	// Hiển thị danh sách các container đang stop
	containerStop, _ := dockerIns.ListContainer(STOPPED)
	Print(containerStop)

	// Running container theo ID
	_ = dockerIns.StartContainerById("2926490de30f")

	// Stop container theo ID
	_ = dockerIns.StopContainerById("7f86a07c6921")
}