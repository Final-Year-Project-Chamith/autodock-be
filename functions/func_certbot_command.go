package functions

import (
	"fmt"
	"os/exec"
)

func RunCertbot(domain string) error {
	// Define the certbot command and arguments
	cmd := exec.Command("certbot", "--nginx", "-d", domain)

	// Capture the output and error
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing certbot command: %v\nOutput: %s", err, string(output))
	}

	// Print the output
	fmt.Println("Command output:")
	fmt.Println(string(output))

	return nil
}
