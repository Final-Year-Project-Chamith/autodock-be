package functions

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCertbotNginxCommand(domain string, email string) error {

	cmd := exec.Command("/usr/bin/certbot", "--nginx", "-d", domain, "--non-interactive", "--agree-tos", "--email", email)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing certbot command: %v\nOutput: %s", err, string(output))
		return fmt.Errorf("certbot failed with error: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Certbot Output: %s\n", string(output))
	return nil
}
