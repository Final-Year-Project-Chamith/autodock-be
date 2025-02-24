package functions

import (
	"fmt"
	"os/exec"
)

func GenerateSSL(domain string, email string) error {
	fmt.Println("stop niginx 1")
	if err := StopNginxConfig(); err != nil{
		return err
	}
	fmt.Println("stop niginx 2")
	cmd := exec.Command("certbot", "certonly", "--standalone", "-d", domain, "--non-interactive", "--agree-tos", "--chamith.eos@gmail.com", email)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate SSL certificate: %v\nOutput: %s", err, string(output))
	}
	if err := StartNginxConfig(); err != nil{
		return err
	}
	fmt.Println("SSL certificate generated successfully!")
	fmt.Println(string(output))
	return nil
}
