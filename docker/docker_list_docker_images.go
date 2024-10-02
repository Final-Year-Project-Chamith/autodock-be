package docker

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types/image"
)

func ListAllDockerImages() (string, error) {

	images, err := Client.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(images)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
