package functions

import (
	"fmt"
	"os/exec"
)

func TestNginxConfig() error {
	cmd := exec.Command("/usr/sbin/nginx", "-t")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx test failed: %s\n", string(output))
		return err
	}

	fmt.Printf("Nginx test succeeded: %s\n", string(output))
	return nil
}

func StopNginxConfig() error {
	cmd := exec.Command("/usr/sbin/nginx", "-s", "stop")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop Nginx: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Nginx stopped successfully.")
	return nil
}

// StartNginxConfig starts Nginx by directly calling the Nginx executable.
func StartNginxConfig() error {
	cmd := exec.Command("/usr/sbin/nginx")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start Nginx: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Nginx started successfully.")
	return nil
}
