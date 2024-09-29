package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
)

func GetDockerContainerLogs() error {
	options := container.LogsOptions{ShowStdout: true}
	out, err := Client.ContainerLogs(context.Background(), "96241f07f965", options)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, out)
	return nil
}
