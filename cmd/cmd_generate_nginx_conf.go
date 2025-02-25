package cmd

import (
	"autodock-be/dto"
	"autodock-be/functions"
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateNginxConf() *cobra.Command {
	var domain string
	var port string
	cmd := &cobra.Command{
		Use:   "Nginx",
		Short: "Run Nginx for generate configuration file",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			domain = args[0]
			port = args[1]
			nginxConf := dto.NginxConf{
				ServerName: domain,
				Port:       port,
			}
			if err := functions.GenerateNginxFileCMD(nginxConf); err != nil {
				return err
			}
			fmt.Println("nginx configuration file generated successfully!")
			return nil
		},
	}
	cmd.Flags().StringVar(&domain, "domain", "", "The domain name")
	cmd.Flags().StringVar(&port, "port", "", "The port")
	return cmd
}
