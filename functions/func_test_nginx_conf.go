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
	stopCmd := exec.Command("systemctl", "stop", "nginx")
	stopOutput, stopErr := stopCmd.CombinedOutput()
	if stopErr != nil {
		fmt.Printf("Failed to stop Nginx: %s\n", string(stopOutput))
		return stopErr
	}

	fmt.Println("Nginx stopped successfully.")
	return nil
}
func StartNginxConfig() error {
	stopCmd := exec.Command("systemctl", "start", "nginx")
	stopOutput, stopErr := stopCmd.CombinedOutput()
	if stopErr != nil {
		fmt.Printf("Failed to stop Nginx: %s\n", string(stopOutput))
		return stopErr
	}

	fmt.Println("Nginx stopped successfully.")
	return nil
}

