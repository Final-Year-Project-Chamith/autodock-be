package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
)

func ListAllDockerImages() error {
	images, err := Client.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		return err
	}
	for _, image := range images {
		fmt.Println(image.Size)
	}
	return nil
}
