package functions

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCertbotWebroot(domain, email, webroot string) error {
	cmd := exec.Command("sudo", "certbot", "certonly",
		"--webroot",
		"-w", webroot,
		"-d", domain,
		"--non-interactive",
		"--agree-tos",
		"--email", email,
	)

	// Capture output and error
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing certbot command: %v\nOutput:\n%s", err, string(output))
		return fmt.Errorf("certbot failed: %v\nOutput:\n%s", err, string(output))
	}

	fmt.Printf("Certbot Output:\n%s\n", string(output))
	return nil
}
