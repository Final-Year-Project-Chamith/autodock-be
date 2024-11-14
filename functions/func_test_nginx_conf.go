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
