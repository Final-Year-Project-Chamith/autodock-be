package functions

import (
	"fmt"
	"os/exec"
)

func GenerateSSL(domain string, email string) error {
	fmt.Println("starting.............")
	fmt.Println("stop niginx 2")
	cmd := exec.Command("certbot", "--nginx", "-d", domain, "--non-interactive", "--agree-tos", "--email", "chamith.eos@gmail.com")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate SSL certificate: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("done.............")
	
	fmt.Println("SSL certificate generated successfully!")
	fmt.Println(string(output))
	return nil
}
