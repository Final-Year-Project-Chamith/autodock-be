package docker

import (
	"github.com/docker/docker/client"
)

func ConnectDocker() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	Client = cli
	defer Client.Close()
}
