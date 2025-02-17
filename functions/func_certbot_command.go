package functions

import (
	"fmt"
	"os/exec"
)

func RunCertbot(domain string) error {
	// Check Certbot version
	versionCmd := exec.Command("certbot", "--version")
	versionOutput, versionErr := versionCmd.CombinedOutput()
	if versionErr != nil {
		return fmt.Errorf("error checking certbot version: %v\nOutput: %s", versionErr, string(versionOutput))
	}

	fmt.Println("Certbot Version:", string(versionOutput))

	// Define the certbot command with non-interactive options
	cmd := exec.Command("certbot", "--nginx", "-d", domain, "--non-interactive", "--agree-tos", "-m", "chamith.eos@gmail.com")

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
