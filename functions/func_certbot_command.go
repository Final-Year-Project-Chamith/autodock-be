package functions

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCertbotManualDNS(domain string) error {
	cmd := exec.Command("certbot",
		"-d", domain,
		"--manual",
		"--preferred-challenges", "dns",
		"certonly",
		"--non-interactive",
		"--agree-tos",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing certbot command: %v\nOutput: %s", err, string(output))
		return fmt.Errorf("certbot failed with error: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Certbot Output: %s\n", string(output))
	return nil
}
