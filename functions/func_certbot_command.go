package functions

import (
	"fmt"
	"log"
	"os/exec"
)


func RunCertbotCommand(domain string, email string) error {

	cmd := exec.Command("/usr/bin/certbot", "certonly", "--standalone",
		"--non-interactive", "--agree-tos", "--email", email,
		"-d", domain)


	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing certbot command: %v\nOutput: %s", err, string(output))
		return fmt.Errorf("certbot failed with error: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Certbot Output: %s\n", string(output))
	return nil
}
