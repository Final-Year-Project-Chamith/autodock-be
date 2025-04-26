package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func CertbotCmd() *cobra.Command {
	var domain string

	cmd := &cobra.Command{
		Use:   "certbot",
		Short: "Run certbot for SSL certificate issuance",
		Args:  cobra.ExactArgs(1), 
		Run: func(cmd *cobra.Command, args []string) {
			domain = args[0]
			fmt.Println("Running certbot for domain:", domain)
			certbot := exec.Command("certbot", "--nginx", "-d", domain)
			certbot.Stdout = os.Stdout
			certbot.Stderr = os.Stderr

			err := certbot.Run()
			if err != nil {
				fmt.Println("Error running certbot:", err)
				os.Exit(1)
			}
		},
	}
	cmd.Flags().StringVar(&domain, "domain", "", "The domain name for certbot")

	return cmd
}
