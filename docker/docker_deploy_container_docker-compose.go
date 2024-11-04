package docker

import (
	"fmt"
	"log"
	"os/exec"
)

func RunDockerComposeDeatched(composeFilePath string) error {
	commands := [][]string{
		{"down"}, 
		{"pull"},          
		{"up", "-d"},      
	}

	for _, args := range commands {
		cmd := exec.Command("/usr/local/bin/docker-compose", append([]string{"-f", composeFilePath}, args...)...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error executing docker-compose %s: %v\nOutput: %s", args[0], err, string(output))
			return err
		}
		fmt.Printf("Docker Compose %s Output: %s\n", args[0], string(output))
	}

	return nil
}