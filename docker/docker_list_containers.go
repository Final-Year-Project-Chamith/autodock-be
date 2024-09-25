package docker

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types/container"
)

func ListAllContainers() (string, error) {
	containers, err := Client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(containers)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}
