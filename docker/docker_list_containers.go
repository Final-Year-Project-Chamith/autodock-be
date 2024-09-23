package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func ListAllContainers() error {
	containers, err := Client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return err
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	}
	return nil
}
