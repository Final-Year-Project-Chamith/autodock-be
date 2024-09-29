package docker

import (
	"fmt"
	"log"
	"os/exec"
)

func RunDockerComposeDeatched(composeFilePath string) error {
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing docker-compose up -d: %v", err)
		return err
	}

	fmt.Printf("Docker Compose Output: %s\n", string(output))
	return nil
}
