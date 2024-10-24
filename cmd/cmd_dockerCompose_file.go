package cmd

import (
	"autodock-be/dto"
	"autodock-be/functions"
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateDockerComposeCMD() *cobra.Command {
	var serviceName, imageName, containerName string
	var ports, volumes []string
	var envVars map[string]string

	cmd := &cobra.Command{
		Use:   "gdc",
		Short: "Generate the docker-compose file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if serviceName == "" || imageName == "" || containerName == "" {
				return fmt.Errorf("service name, image, and container name are required")
			}

			// Create service from input flags
			service := dto.Service{
				Name:      serviceName,
				Image:     imageName,
				Container: containerName,
				Ports:     ports,
				Volumes:   volumes,
				EnvVars:   envVars,
			}

			dockerCompose := dto.DockerCompose{
				Services: []dto.Service{service},
			}


			if err := functions.GenerateDockerComposeFile(dockerCompose,"chamith"); err != nil {
				return err
			}

			fmt.Println("docker-compose.yml file generated successfully!")
			return nil
		},
	}


	cmd.Flags().StringVarP(&serviceName, "service", "s", "", "Service name (required)")
	cmd.Flags().StringVarP(&imageName, "image", "i", "", "Docker image name (required)")
	cmd.Flags().StringVarP(&containerName, "container", "c", "", "Container name (required)")
	cmd.Flags().StringSliceVar(&ports, "ports", []string{}, "Ports to expose (format: hostPort:containerPort)")
	cmd.Flags().StringSliceVar(&volumes, "volumes", []string{}, "Volumes to mount (format: hostPath:containerPath)")
	cmd.Flags().StringToStringVar(&envVars, "env", map[string]string{}, "Environment variables (key=value)")

	cmd.MarkFlagRequired("service")
	cmd.MarkFlagRequired("image")
	cmd.MarkFlagRequired("container")

	return cmd
}
